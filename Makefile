# OnceHuman工具集 Makefile

# Go环境设置
GOPATH := $(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)

# 变量定义
RELEASE_DIR := release
BACKEND_DIR := backend
FRONTEND_DIR := frontend
BOT_DIR := bot
LAUNCHER_DIR := cmd/launcher

# 构建目标
BACKEND_BIN := $(RELEASE_DIR)/backend/server
FRONTEND_DIST := $(RELEASE_DIR)/frontend
BOT_BIN := $(RELEASE_DIR)/bot/bot
LAUNCHER_BIN := $(RELEASE_DIR)/launcher

# Go编译参数
GO_BUILD := go build -ldflags="-s -w"
GO_BUILD_STATIC := CGO_ENABLED=0 $(GO_BUILD) -a -installsuffix cgo

# 颜色输出
GREEN := \033[0;32m
YELLOW := \033[0;33m
RED := \033[0;31m
NC := \033[0m

.PHONY: all clean build-integrated build-release run-integrated help

# 默认目标
all: build-integrated

# 帮助信息
help:
	@echo "$(GREEN)OnceHuman工具集 构建系统$(NC)"
	@echo ""
	@echo "可用命令："
	@echo "  $(YELLOW)make build-integrated$(NC)  - 构建集成版本（开发用）"
	@echo "  $(YELLOW)make build-release$(NC)     - 构建发布版本（生产用）"
	@echo "  $(YELLOW)make run-integrated$(NC)    - 运行集成版本"
	@echo "  $(YELLOW)make run-bot$(NC)           - 运行集成版本（包含机器人）"
	@echo "  $(YELLOW)make clean$(NC)             - 清理构建文件"
	@echo "  $(YELLOW)make test$(NC)              - 运行测试"
	@echo ""
	@echo "单独构建："
	@echo "  $(YELLOW)make build-backend$(NC)     - 只构建后端"
	@echo "  $(YELLOW)make build-frontend$(NC)    - 只构建前端"
	@echo "  $(YELLOW)make build-bot$(NC)         - 只构建机器人"
	@echo "  $(YELLOW)make build-launcher$(NC)    - 只构建启动器"

# 清理
clean:
	@echo "$(YELLOW)清理构建文件...$(NC)"
	rm -rf $(RELEASE_DIR)
	cd $(BACKEND_DIR) && make clean
	cd $(FRONTEND_DIR) && rm -rf dist node_modules/.cache
	cd $(BOT_DIR) && go clean
	@echo "$(GREEN)清理完成！$(NC)"

# 创建目录
$(RELEASE_DIR):
	@mkdir -p $(RELEASE_DIR)
	@mkdir -p $(RELEASE_DIR)/backend
	@mkdir -p $(RELEASE_DIR)/frontend
	@mkdir -p $(RELEASE_DIR)/bot
	@mkdir -p $(RELEASE_DIR)/config

# 构建后端
build-backend: $(RELEASE_DIR)
	@echo "$(YELLOW)构建后端服务...$(NC)"
	cd $(BACKEND_DIR) && $(GO_BUILD_STATIC) -o ../$(BACKEND_BIN) cmd/server/main.go
	@cp -r $(BACKEND_DIR)/api ../$(RELEASE_DIR)/backend/
	@echo "$(GREEN)后端构建完成！$(NC)"

# 构建前端
build-frontend: $(RELEASE_DIR)
	@echo "$(YELLOW)构建前端应用...$(NC)"
	cd $(FRONTEND_DIR) && npm install && npm run build
	@cp -r $(FRONTEND_DIR)/dist/* $(FRONTEND_DIST)/
	@echo "$(GREEN)前端构建完成！$(NC)"

# 构建机器人
build-bot: $(RELEASE_DIR)
	@echo "$(YELLOW)构建Discord机器人...$(NC)"
	cd $(BOT_DIR) && $(GO_BUILD_STATIC) -o ../$(BOT_BIN) main.go
	@cp $(BOT_DIR)/.env.example $(RELEASE_DIR)/bot/
	@echo "$(GREEN)机器人构建完成！$(NC)"

# 构建启动器
build-launcher: $(RELEASE_DIR)
	@echo "$(YELLOW)构建启动器...$(NC)"
	$(GO_BUILD_STATIC) -o $(LAUNCHER_BIN) $(LAUNCHER_DIR)/main.go
	@echo "$(GREEN)启动器构建完成！$(NC)"

# 集成构建（开发模式）
build-integrated: clean $(RELEASE_DIR)
	@echo "$(GREEN)========== 集成构建模式 ==========$(NC)"
	@$(MAKE) build-backend
	@$(MAKE) build-frontend
	@$(MAKE) build-bot
	@$(MAKE) build-launcher
	@echo ""
	@echo "$(YELLOW)创建配置文件...$(NC)"
	@cp -f scripts/integrated-config.sh $(RELEASE_DIR)/
	@chmod +x $(RELEASE_DIR)/integrated-config.sh
	@echo ""
	@echo "$(GREEN)========== 构建完成！ ==========$(NC)"
	@echo "运行: $(YELLOW)make run-integrated$(NC) 启动服务"
	@echo ""

# 发布构建（生产模式）
build-release: clean $(RELEASE_DIR)
	@echo "$(GREEN)========== 发布构建模式 ==========$(NC)"
	@$(MAKE) build-backend
	@$(MAKE) build-frontend
	@$(MAKE) build-bot
	@$(MAKE) build-launcher
	@echo ""
	@echo "$(YELLOW)准备发布文件...$(NC)"
	# 创建独立的启动脚本
	@cp -f scripts/start-backend.sh $(RELEASE_DIR)/backend/
	@cp -f scripts/start-frontend.sh $(RELEASE_DIR)/frontend/
	@cp -f scripts/start-bot.sh $(RELEASE_DIR)/bot/
	@chmod +x $(RELEASE_DIR)/backend/start-backend.sh
	@chmod +x $(RELEASE_DIR)/frontend/start-frontend.sh
	@chmod +x $(RELEASE_DIR)/bot/start-bot.sh
	# 创建配置文件模板
	@cp -f configs/backend.env.example $(RELEASE_DIR)/backend/.env.example
	@cp -f configs/frontend.env.example $(RELEASE_DIR)/frontend/.env.example
	# 创建README
	@cp -f scripts/RELEASE_README.md $(RELEASE_DIR)/README.md
	@echo ""
	@echo "$(GREEN)========== 发布构建完成！ ==========$(NC)"
	@echo "发布文件位于: $(YELLOW)$(RELEASE_DIR)/$(NC)"
	@echo ""

# 运行集成版本
run-integrated: build-integrated
	@echo "$(GREEN)启动集成服务...$(NC)"
	cd $(RELEASE_DIR) && ./launcher -mode integrated

# 运行集成版本（包含机器人）
run-bot: build-integrated
	@echo "$(GREEN)启动集成服务（含机器人）...$(NC)"
	cd $(RELEASE_DIR) && ./launcher -mode integrated -bot

# 开发模式
dev:
	@echo "$(GREEN)启动开发模式...$(NC)"
	@echo "请在不同终端运行："
	@echo "  1. cd backend && make run"
	@echo "  2. cd frontend && npm run dev"
	@echo "  3. cd bot && go run main.go"

# 测试
test:
	@echo "$(YELLOW)运行测试...$(NC)"
	cd $(BACKEND_DIR) && go test ./...
	cd $(BOT_DIR) && go test ./...
	@echo "$(GREEN)测试完成！$(NC)"

# Docker构建（可选）
docker-build:
	@echo "$(YELLOW)构建Docker镜像...$(NC)"
	docker build -t oncehuman-backend -f docker/Dockerfile.backend .
	docker build -t oncehuman-frontend -f docker/Dockerfile.frontend .
	docker build -t oncehuman-bot -f docker/Dockerfile.bot .
	@echo "$(GREEN)Docker镜像构建完成！$(NC)"

# 版本信息
version:
	@echo "OnceHuman工具集 v1.0.0"
	@echo "构建时间: $$(date '+%Y-%m-%d %H:%M:%S')"