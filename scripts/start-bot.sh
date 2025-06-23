#!/bin/bash

# Discord机器人启动脚本

# 设置工作目录
cd "$(dirname "$0")"

# 检查环境变量文件
if [ ! -f .env ]; then
    echo "创建 .env 文件..."
    cp .env.example .env
    echo "请编辑 .env 文件配置Discord机器人Token"
    exit 1
fi

# 检查Token配置
if ! grep -q "DISCORD_BOT_TOKEN=" .env || grep -q "DISCORD_BOT_TOKEN=your_discord_bot_token_here" .env; then
    echo "错误：请在 .env 文件中配置 DISCORD_BOT_TOKEN"
    exit 1
fi

echo "启动Discord机器人..."

# 启动机器人
./bot