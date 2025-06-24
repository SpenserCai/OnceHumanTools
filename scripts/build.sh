#!/bin/bash

# OnceHuman工具集 - 跨平台构建脚本

set -e

# 设置 Go 环境变量
export GOPATH=$(go env GOPATH)
export PATH="$GOPATH/bin:$PATH"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# 构建模式
BUILD_MODE=${1:-integrated}
TARGET_OS=${2:-$(go env GOOS)}
TARGET_ARCH=${3:-$(go env GOARCH)}

# 设置可执行文件扩展名
EXE_EXT=""
if [ "$TARGET_OS" == "windows" ]; then
    EXE_EXT=".exe"
fi

echo -e "${GREEN}OnceHuman工具集 - 构建脚本${NC}"
echo -e "构建模式: ${YELLOW}$BUILD_MODE${NC}"
echo -e "目标系统: ${YELLOW}$TARGET_OS/$TARGET_ARCH${NC}"
echo ""

# 创建输出目录
RELEASE_DIR="release"
if [ "$BUILD_MODE" == "cross" ]; then
    RELEASE_DIR="release/${TARGET_OS}_${TARGET_ARCH}"
fi

rm -rf $RELEASE_DIR
mkdir -p $RELEASE_DIR/{backend,frontend,bot,config}

# 构建后端
echo -e "${YELLOW}构建后端服务...${NC}"
cd backend
echo -e "${YELLOW}生成Swagger代码...${NC}"
make generate-swagger
echo -e "${YELLOW}编译后端可执行文件...${NC}"
GOOS=$TARGET_OS GOARCH=$TARGET_ARCH CGO_ENABLED=0 \
    go build -ldflags="-s -w" -o ../$RELEASE_DIR/backend/server${EXE_EXT} cmd/server/main.go
cp -r api ../$RELEASE_DIR/backend/
cd ..
echo -e "${GREEN}✓ 后端构建完成${NC}"

# 构建前端
echo -e "${YELLOW}构建前端应用...${NC}"
cd frontend
npm install --silent
npm run build
cp -r dist/* ../$RELEASE_DIR/frontend/
cd ..
echo -e "${GREEN}✓ 前端构建完成${NC}"

# 构建机器人
# echo -e "${YELLOW}构建Discord机器人...${NC}"
# cd bot
# GOOS=$TARGET_OS GOARCH=$TARGET_ARCH CGO_ENABLED=0 \
#     go build -ldflags="-s -w" -o ../$RELEASE_DIR/bot/bot main.go
# cp .env.example ../$RELEASE_DIR/bot/
# cd ..
# echo -e "${GREEN}✓ 机器人构建完成${NC}"

echo -e "${YELLOW}跳过机器人构建（需要修复接口问题）${NC}"

# 构建启动器
echo -e "${YELLOW}构建启动器...${NC}"
GOOS=$TARGET_OS GOARCH=$TARGET_ARCH CGO_ENABLED=0 \
    go build -ldflags="-s -w" -o $RELEASE_DIR/launcher${EXE_EXT} cmd/launcher/main.go
echo -e "${GREEN}✓ 启动器构建完成${NC}"

# 复制配置和脚本
echo -e "${YELLOW}准备配置文件...${NC}"
if [ "$TARGET_OS" == "windows" ]; then
    # Windows批处理脚本
    cp scripts/windows/*.bat $RELEASE_DIR/
else
    # Unix脚本
    cp scripts/*.sh $RELEASE_DIR/
    chmod +x $RELEASE_DIR/*.sh
    cp scripts/start-*.sh $RELEASE_DIR/backend/
    cp scripts/start-*.sh $RELEASE_DIR/frontend/
    cp scripts/start-*.sh $RELEASE_DIR/bot/
fi

# 复制文档
cp scripts/RELEASE_README.md $RELEASE_DIR/README.md

# 创建版本信息
cat > $RELEASE_DIR/VERSION << EOF
OnceHuman工具集
版本: 1.0.0
构建时间: $(date '+%Y-%m-%d %H:%M:%S')
目标平台: $TARGET_OS/$TARGET_ARCH
构建模式: $BUILD_MODE
EOF

# 打包（如果是发布模式）
if [ "$BUILD_MODE" == "release" ] || [ "$BUILD_MODE" == "cross" ]; then
    ARCHIVE_NAME="oncehuman-tools-${TARGET_OS}-${TARGET_ARCH}.tar.gz"
    echo -e "${YELLOW}创建发布包: $ARCHIVE_NAME${NC}"
    cd release
    tar -czf ../$ARCHIVE_NAME *
    cd ..
    echo -e "${GREEN}✓ 发布包创建完成: $ARCHIVE_NAME${NC}"
fi

echo ""
echo -e "${GREEN}========== 构建完成！ ==========${NC}"
echo -e "输出目录: ${YELLOW}$RELEASE_DIR${NC}"

if [ "$BUILD_MODE" == "integrated" ]; then
    echo -e "运行: ${YELLOW}cd $RELEASE_DIR && ./launcher${NC}"
fi