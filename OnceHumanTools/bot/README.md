# OnceHuman工具集 - 机器人模块

这是OnceHuman工具集的机器人模块，支持多平台机器人接入。目前已实现Discord机器人，使用现代化的Slash Commands（斜杠命令）交互方式。

## 特性

- 🎮 **Discord支持**：使用交互式斜杠命令，提供更好的用户体验
- 🔌 **模块化设计**：易于扩展到其他平台（Telegram、Slack等）
- 🛠 **完整功能**：支持词条概率计算、强化概率计算等核心功能
- 📊 **美观展示**：使用Discord嵌入消息，信息展示清晰直观

## Discord机器人功能

### 交互式命令

1. **`/help`** - 显示帮助信息
   - 展示所有可用命令和使用方法

2. **`/affix`** - 词条概率计算
   - 参数：
     - `slots`：词条数量 (1-10)
     - `targets`：目标词条ID列表，逗号分隔
     - `show_combinations`：是否显示详细组合
   - 示例：`/affix slots:4 targets:1,4,5`

3. **`/strengthen single`** - 单词条强化概率
   - 参数：
     - `affix_id`：词条ID (1-10)
     - `current_level`：当前等级 (0-5)
     - `target_level`：目标等级 (1-5)
     - `slot_count`：词条数量
     - `tries`：强化次数
   - 示例：`/strengthen single affix_id:1 current_level:0 target_level:3 slot_count:4 tries:50`

4. **`/strengthen multi`** - 多词条强化概率
   - 参数：
     - `targets`：目标格式 ID:当前:目标
     - `slot_count`：词条数量
     - `tries`：强化次数
   - 示例：`/strengthen multi targets:1:0:3,4:1:5 slot_count:4 tries:100`

## 快速开始

### 1. 创建Discord机器人

1. 访问 [Discord Developer Portal](https://discord.com/developers/applications)
2. 创建新应用程序
3. 进入"Bot"页面，创建机器人
4. 复制机器人Token
5. 在"OAuth2 > URL Generator"中：
   - 勾选 `bot` 和 `applications.commands`
   - 选择所需权限（至少需要 `Send Messages` 和 `Use Slash Commands`）
   - 使用生成的URL邀请机器人到服务器

### 2. 配置环境变量

```bash
cp .env.example .env
```

编辑 `.env` 文件：
```env
DISCORD_BOT_TOKEN=你的机器人Token
DISCORD_GUILD_ID=你的服务器ID（可选，用于开发测试）
BOT_DEV_MODE=false
```

### 3. 安装依赖

```bash
go mod download
```

### 4. 运行机器人

```bash
go run main.go
```

## 架构设计

```
bot/
├── core/               # 核心接口定义
│   ├── interfaces.go   # Bot、Command等接口
│   └── manager.go      # 机器人管理器
├── platforms/          # 各平台实现
│   └── discord/        # Discord平台
│       ├── bot.go      # Discord机器人实现
│       └── commands/   # Discord命令实现
├── shared/             # 共享组件
└── main.go            # 主程序入口
```

## 扩展到其他平台

要添加新平台支持，只需：

1. 在 `platforms/` 下创建新平台目录
2. 实现 `core.Bot` 接口
3. 实现平台特定的命令处理
4. 在 `main.go` 中注册新平台

示例结构：
```go
type TelegramBot struct {
    // 平台特定字段
}

func (t *TelegramBot) Start(ctx context.Context) error {
    // 实现启动逻辑
}

func (t *TelegramBot) Stop() error {
    // 实现停止逻辑
}

// ... 其他接口方法
```

## 开发说明

### 命令注册时机

- **Guild命令**（指定服务器）：立即生效，适合开发测试
- **全局命令**：可能需要1小时才能在所有服务器生效

### 调试技巧

1. 设置 `BOT_DEV_MODE=true` 启用开发模式
2. 指定 `DISCORD_GUILD_ID` 进行快速测试
3. 查看日志了解命令注册和处理情况

## 词条ID对照表

| ID | 词条名称 |
|----|----------|
| 1  | 异常伤害 |
| 2  | 弹匣容量 |
| 3  | 换弹速度加成 |
| 4  | 对普通敌人伤害 |
| 5  | 对精英敌人伤害 |
| 6  | 对上位者伤害 |
| 7  | 最大生命值 |
| 8  | 头部受伤减免 |
| 9  | 枪械伤害减免 |
| 10 | 异常伤害减免 |

## 贡献指南

欢迎贡献新功能或平台支持！请确保：

1. 遵循现有的代码结构
2. 实现必要的接口
3. 添加适当的错误处理
4. 更新相关文档

## 许可证

MIT License