# OnceHuman工具集

一个功能强大的OnceHuman游戏工具集，提供Web界面、REST API和Discord机器人多种使用方式。

## 🌟 特性

- **模组词条概率计算器** - 精确计算特定词条组合出现的概率
- **模组强化概率计算器** - 计算词条强化到目标等级的成功率
- **炫酷科幻UI** - 采用Vue3打造的沉浸式科幻风格界面
- **RESTful API** - 基于Go和Swagger的高性能后端服务
- **Discord机器人** - 在Discord中直接使用计算工具
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

### Discord机器人

1. 创建 `.env` 文件：
```env
DISCORD_TOKEN=your_discord_bot_token
```

2. 运行机器人：
```bash
cd bot
go mod tidy
go run main.go
```

## 📖 使用指南

### Web界面

1. 访问首页，选择需要的工具
2. 输入参数，点击计算
3. 查看详细的计算结果和可视化图表

### Discord命令

- `!oh help` - 显示帮助信息
- `!oh affix <词条数量> <目标词条ID列表>` - 计算词条概率
- `!oh strengthen <初始等级> <目标等级> [模式]` - 计算强化概率

示例：
```
!oh affix 3 1,4,5,6
!oh strengthen 1,1,1,1 2,2,2,2
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
├── bot/            # Discord机器人
│   ├── commands/   # 命令实现
│   └── handlers/   # 事件处理
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

### Discord机器人
- discordgo
- 命令模式架构

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