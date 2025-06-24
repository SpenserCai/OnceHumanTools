@echo off
chcp 65001 >nul

REM Discord机器人启动脚本

REM 设置工作目录
cd /d "%~dp0"

REM 检查环境变量文件
if not exist .env (
    echo 创建 .env 文件...
    copy .env.example .env >nul
    echo 请编辑 .env 文件配置Discord机器人Token
    pause
    exit /b 1
)

REM 检查Token配置
findstr /C:"DISCORD_BOT_TOKEN=" .env >nul
if errorlevel 1 (
    echo 错误：请在 .env 文件中配置 DISCORD_BOT_TOKEN
    pause
    exit /b 1
)

findstr /C:"DISCORD_BOT_TOKEN=your_discord_bot_token_here" .env >nul
if not errorlevel 1 (
    echo 错误：请在 .env 文件中配置有效的 DISCORD_BOT_TOKEN
    pause
    exit /b 1
)

echo 启动Discord机器人...
echo.

REM 启动机器人
bot.exe

pause 