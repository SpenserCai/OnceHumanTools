# OnceHuman工具集

一个功能强大的OnceHuman游戏工具集，提供Web界面、REST API和Discord机器人多种使用方式。

## 🌟 特性

- **模组词条概率计算器** - 精确计算特定词条组合出现的概率
- **模组强化概率计算器** - 计算词条强化到目标等级的成功率
- **炫酷科幻UI** - 采用Vue3打造的沉浸式科幻风格界面
- **RESTful API** - 基于Go和Swagger的高性能后端服务
- **Discord机器人** - 使用现代化的Slash Commands交互方式
- **多平台支持** - 模块化设计，易于扩展到Telegram、Slack等平台
- **可扩展架构** - 轻松添加新的游戏工具

## 🚀 快速开始

### 后端服务

```bash
cd backend
go mod tidy
make run
```

API文档将在 http://localhost:8080/docs 可用

### 前端应用

```bash
cd frontend
npm install
npm run dev
```

访问 http://localhost:3000 查看应用

### Discord机器人（交互式命令）

1. 创建 `.env` 文件（参考 `bot/.env.example`）：
```env
DISCORD_BOT_TOKEN=your_discord_bot_token
DISCORD_GUILD_ID=your_guild_id（可选，用于开发测试）
BOT_DEV_MODE=false
```

2. 运行机器人：
```bash
cd bot
go mod download
go run main.go
```

## 📖 使用指南

### Web界面

1. 访问首页，选择需要的工具
2. 输入参数，点击计算
3. 查看详细的计算结果和可视化图表

### Discord交互式命令（Slash Commands）

- `/help` - 显示帮助信息
- `/affix` - 计算词条概率
  - `slots`: 词条数量 (1-10)
  - `targets`: 目标词条ID列表，逗号分隔
  - `show_combinations`: 是否显示详细组合
- `/strengthen single` - 计算单个词条强化概率
  - `affix_id`: 词条ID (1-10)
  - `current_level`: 当前等级 (0-5)
  - `target_level`: 目标等级 (1-5)
  - `slot_count`: 词条数量
  - `tries`: 强化次数
- `/strengthen multi` - 计算多个词条强化概率
  - `targets`: 格式 ID:当前:目标，逗号分隔
  - `slot_count`: 词条数量
  - `tries`: 强化次数

示例：
```
/affix slots:4 targets:1,4,5
/strengthen single affix_id:1 current_level:0 target_level:3 slot_count:4 tries:50
/strengthen multi targets:1:0:3,4:1:5 slot_count:4 tries:100
```

### API接口

#### 健康检查
```
GET /api/v1/health
```

#### 获取词条列表
```
GET /api/v1/mod/affix/list
```

#### 计算词条概率
```
POST /api/v1/mod/affix/probability
{
  "slotCount": 3,
  "targetAffixIds": [1, 4, 5, 6],
  "showCombinations": true
}
```

#### 计算强化概率
```
POST /api/v1/mod/strengthen/probability
{
  "initialLevels": [1, 2, 3, 1],
  "targetLevels": [3, 4, 5, 2],
  "orderIndependent": true,
  "showPaths": false
}
```

## 🏗️ 项目结构

```
OnceHumanTools/
├── backend/          # Go后端服务
│   ├── api/         # Swagger API定义
│   ├── cmd/         # 主程序入口
│   ├── internal/    # 内部实现
│   └── pkg/         # 公共包
├── frontend/        # Vue3前端应用
│   ├── src/
│   │   ├── api/     # API调用
│   │   ├── views/   # 页面组件
│   │   └── styles/  # 样式文件
├── bot/            # 多平台机器人模块
│   ├── core/       # 核心接口定义
│   ├── platforms/  # 各平台实现
│   │   └── discord/# Discord平台
│   └── shared/     # 共享组件
└── docs/          # 文档
```

## 🛠️ 技术栈

### 后端
- Go 1.21+
- go-swagger
- Gorilla Mux
- CORS中间件

### 前端
- Vue 3
- Vite
- Element Plus
- Chart.js
- Three.js
- SCSS

### 机器人模块
- discordgo (Discord平台)
- 交互式命令 (Slash Commands)
- 模块化架构，支持多平台扩展
- 接口驱动设计

## 📊 游戏机制说明

### 词条系统
- 游戏中共有10种不同的词条
- 同一个模组中，相同词条不会重复出现
- 每次随机都是从剩余的词条池中选择

### 强化系统
- 一个模组总共有4个词条
- 每个词条最低1级，最高5级
- 总共有5次强化机会
- 每次强化随机强化4个词条中的一个
- 如果有词条到达5级，下次强化时该词条不会在随机范围内

## 🤝 贡献

欢迎提交Issue和Pull Request！

## 📄 许可

MIT License