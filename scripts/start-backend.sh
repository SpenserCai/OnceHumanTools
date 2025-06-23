#!/bin/bash

# 后端服务启动脚本

# 设置 Go 环境变量
export GOPATH=$(go env GOPATH 2>/dev/null || echo "$HOME/go")
export PATH="$GOPATH/bin:$PATH"

# 设置工作目录
cd "$(dirname "$0")"

# 检查环境变量文件
if [ ! -f .env ]; then
    echo "创建 .env 文件..."
    cp .env.example .env
    echo "请编辑 .env 文件配置环境变量"
    exit 1
fi

# 加载环境变量
export $(cat .env | grep -v '^#' | xargs)

# 设置默认端口
PORT=${PORT:-8080}

echo "启动后端服务..."
echo "端口: $PORT"
echo "API文档: http://localhost:$PORT/docs"

# 启动服务
echo "使用 HTTP 模式启动服务（开发环境）..."
./server --scheme=http --host=0.0.0.0 --port $PORT