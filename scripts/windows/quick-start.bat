@echo off
chcp 65001 >nul

REM OnceHuman工具集 - 快速启动

REM 设置工作目录
cd /d "%~dp0"

echo ================================
echo OnceHuman工具集 - 快速启动
echo ================================
echo.
echo 正在启动服务，请稍候...
echo.

REM 直接启动启动器，不显示配置信息
launcher.exe

if errorlevel 1 (
    echo.
    echo 启动失败，请检查：
    echo 1. 是否有端口冲突
    echo 2. 是否有必要的权限
    echo 3. 查看错误信息
    pause
) 