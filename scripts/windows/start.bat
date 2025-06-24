@echo off
chcp 65001 >nul

REM OnceHuman工具集 - 主启动脚本

REM 设置工作目录
cd /d "%~dp0"

echo ================================
echo OnceHuman工具集
echo ================================
echo.
echo 正在启动集成服务...
echo.

REM 显示配置信息
call integrated-config.bat

REM 启动集成启动器
echo 启动中...
launcher.exe

pause 