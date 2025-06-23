.PHONY: build run test clean swagger generate-swagger

# Go环境设置
GOPATH := $(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)

# 构建可执行文件
build:
	go build -o bin/server cmd/server/main.go

# 运行服务器
run:
	go run cmd/server/main.go --port=8080

# 运行测试
test:
	go test ./...

# 清理生成的文件
clean:
	@# 备份重要的配置文件
	@if [ -f restapi/configure_oncehuman_tools.go ]; then \
		cp restapi/configure_oncehuman_tools.go configure_oncehuman_tools.go.backup; \
	fi
	rm -rf bin/
	rm -rf models/
	rm -rf restapi/
	@# 恢复配置文件
	@if [ -f configure_oncehuman_tools.go.backup ]; then \
		mkdir -p restapi; \
		mv configure_oncehuman_tools.go.backup restapi/configure_oncehuman_tools.go; \
	fi

# 生成swagger代码
generate-swagger:
	swagger generate server -A oncehuman-tools -f api/swagger.yaml -t . --exclude-main

# 更新依赖
deps:
	go mod tidy

# 格式化代码
fmt:
	go fmt ./...

# 检查代码
lint:
	golangci-lint run

# 开发模式运行
dev:
	go run cmd/server/main.go --port=8080 --host=0.0.0.0