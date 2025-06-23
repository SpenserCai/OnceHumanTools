#!/bin/bash

# 前端服务启动脚本

# 设置工作目录
cd "$(dirname "$0")"

# 检查是否需要安装静态服务器
if ! command -v serve &> /dev/null; then
    echo "安装静态文件服务器..."
    npm install -g serve
fi

# 设置端口
PORT=${PORT:-3000}

echo "启动前端服务..."
echo "地址: http://localhost:$PORT"

# 启动静态文件服务
serve -s . -l $PORT