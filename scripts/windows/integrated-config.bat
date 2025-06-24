@echo off
chcp 65001 >nul

REM OnceHuman工具集 - 集成模式配置脚本

echo ================================
echo OnceHuman工具集 - 集成模式
echo ================================
echo.
echo 服务将在以下地址启动：
echo   主入口: http://localhost:9000
echo   API文档: http://localhost:9000/docs
echo.
echo 可用参数：
echo   launcher.exe -bot                     启用Discord机器人
echo   launcher.exe -port 8888               更改启动器端口
echo   launcher.exe -backend-port 8080       更改后端端口
echo   launcher.exe -backend-scheme http     使用HTTP协议（默认）
echo   launcher.exe -backend-scheme https    使用HTTPS协议
echo.
echo HTTPS模式额外参数：
echo   launcher.exe -backend-scheme https ^
echo                -tls-certificate cert.pem ^
echo                -tls-key key.pem ^
echo                [-tls-ca ca.pem]
echo.
echo 环境变量配置：
echo   请确保已配置 bot\.env 文件（如需使用机器人）
echo.

pause 