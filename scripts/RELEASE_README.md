# OnceHuman工具集 - 发布版本

欢迎使用OnceHuman工具集！本发布包包含了完整的应用程序。

## 📦 包含内容

```
release/
├── backend/          # 后端服务
│   ├── server       # 可执行文件
│   ├── api/         # API定义
│   └── start-backend.sh
├── frontend/        # 前端应用
│   ├── index.html   # 入口文件
│   ├── assets/      # 静态资源
│   └── start-frontend.sh
├── bot/             # Discord机器人
│   ├── bot          # 可执行文件
│   ├── .env.example # 配置示例
│   └── start-bot.sh
└── launcher         # 统一启动器
```

## 🚀 快速开始

### 方式一：使用启动器（推荐）

```bash
# 启动所有服务（前端+后端）
./launcher

# 启动所有服务（包括Discord机器人）
./launcher -bot

# 自定义端口
./launcher -port 8888
```

### 方式二：分别启动各服务

1. **启动后端服务**
   ```bash
   cd backend
   ./start-backend.sh
   ```

2. **启动前端服务**
   ```bash
   cd frontend
   ./start-frontend.sh
   ```

3. **启动Discord机器人**（可选）
   ```bash
   cd bot
   # 首次运行需要配置 .env 文件
   cp .env.example .env
   # 编辑 .env 文件，填入Discord Bot Token
   ./start-bot.sh
   ```

## 🔧 配置说明

### 后端配置
- 配置文件：`backend/.env`
- 默认端口：8080
- API文档：http://localhost:8080/docs

### 前端配置
- 默认端口：3000
- 访问地址：http://localhost:3000

### Discord机器人配置
- 配置文件：`bot/.env`
- 必需配置：
  - `DISCORD_BOT_TOKEN`: Discord机器人令牌
  - `DISCORD_GUILD_ID`: 服务器ID（可选，用于开发测试）

## 📊 系统要求

- 操作系统：Linux/macOS/Windows
- 内存：至少512MB可用内存
- 端口：确保8080、3000、9000端口未被占用

## 🛠 故障排除

### 端口被占用
修改启动参数指定其他端口：
```bash
./launcher -port 8888 -backend-port 8081 -frontend-port 3001
```

### Discord机器人无法启动
1. 检查 `.env` 文件中的 `DISCORD_BOT_TOKEN` 是否正确
2. 确保机器人已被邀请到服务器
3. 检查网络连接

### 前端无法访问后端
1. 确保后端服务已启动
2. 检查防火墙设置
3. 查看后端日志排查问题

## 📝 更新日志

### v1.0.0
- 初始发布版本
- 支持词条概率计算
- 支持强化概率计算
- Discord机器人集成
- 炫酷科幻UI界面

## 📞 支持

如遇到问题，请访问项目主页提交Issue。

---
© 2024 OnceHuman工具集. MIT License.