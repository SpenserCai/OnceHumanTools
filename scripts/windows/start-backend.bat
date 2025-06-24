@echo off
chcp 65001 >nul

REM 后端服务启动脚本

REM 设置工作目录
cd /d "%~dp0"

REM 检查环境变量文件
if not exist .env (
    echo 创建 .env 文件...
    copy .env.example .env >nul
    echo 请编辑 .env 文件配置环境变量
    pause
    exit /b 1
)

REM 设置默认端口
if not defined PORT set PORT=8080

echo 启动后端服务...
echo 端口: %PORT%
echo API文档: http://localhost:%PORT%/docs
echo.

REM 启动服务
echo 使用 HTTP 模式启动服务（开发环境）...
server.exe --scheme=http --host=0.0.0.0 --port %PORT%

pause 