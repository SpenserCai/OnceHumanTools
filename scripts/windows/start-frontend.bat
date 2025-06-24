@echo off
chcp 65001 >nul

REM 前端服务启动脚本

REM 设置工作目录
cd /d "%~dp0"

REM 检查是否需要安装静态服务器
where serve >nul 2>&1
if errorlevel 1 (
    echo 安装静态文件服务器...
    npm install -g serve
    if errorlevel 1 (
        echo 错误：无法安装serve，请确保已安装Node.js和npm
        pause
        exit /b 1
    )
)

REM 设置端口
if not defined PORT set PORT=3000

echo 启动前端服务...
echo 地址: http://localhost:%PORT%
echo.

REM 启动静态文件服务
serve -s . -l %PORT%

pause 