@echo off
chcp 65001 >nul

REM OnceHuman工具集 - Windows使用说明

echo ================================================
echo             OnceHuman工具集
echo            Windows版本使用说明
echo ================================================
echo.
echo 🚀 快速开始：
echo   双击 quick-start.bat        快速启动（推荐）
echo   双击 start.bat             标准启动（显示配置）
echo.
echo 📁 手动启动各个组件：
echo   start-backend.bat          仅启动后端服务
echo   start-frontend.bat         仅启动前端界面
echo   start-bot.bat              仅启动Discord机器人
echo.
echo 🔧 服务地址：
echo   主入口: http://localhost:9000
echo   API文档: http://localhost:9000/docs
echo   前端界面: http://localhost:3000 （独立运行时）
echo   后端API: http://localhost:8080 （独立运行时）
echo.
echo 💡 使用提示：
echo   1. 首次运行会自动创建配置文件
echo   2. 确保端口9000、8080、3000未被占用
echo   3. 如需使用Discord机器人，请配置bot/.env文件
echo   4. 建议使用 quick-start.bat 进行日常使用
echo.
echo 🛠️ 系统要求：
echo   - Windows 7/8/10/11 （64位）
echo   - Node.js 16+ （用于前端静态服务）
echo   - 可选：Discord机器人Token（用于机器人功能）
echo.
echo 📋 故障排除：
echo   - 端口冲突：修改环境变量或关闭占用端口的程序
echo   - 权限问题：以管理员身份运行
echo   - 网络问题：检查防火墙设置
echo.
echo 📞 技术支持：
echo   GitHub: https://github.com/SpenserCai/OnceHumanTools
echo   Issues: 如有问题请在GitHub提交Issue
echo.
echo ================================================
echo 按任意键关闭...
pause >nul 