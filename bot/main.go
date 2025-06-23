package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/oncehuman/tools/bot/core"
	"github.com/oncehuman/tools/bot/platforms/discord"
	"github.com/oncehuman/tools/bot/platforms/discord/commands"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("未找到 .env 文件，使用系统环境变量")
	}

	// 创建机器人管理器
	manager := core.NewBotManager()

	// 初始化Discord机器人
	if discordToken := os.Getenv("DISCORD_BOT_TOKEN"); discordToken != "" {
		if err := initDiscordBot(manager, discordToken); err != nil {
			log.Fatalf("初始化Discord机器人失败: %v", err)
		}
	} else {
		log.Println("未配置 DISCORD_BOT_TOKEN，跳过Discord机器人")
	}

	// 检查是否有机器人被注册
	if len(manager.ListBots()) == 0 {
		log.Fatal("没有可用的机器人，请至少配置一个机器人")
	}

	// 创建上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动所有机器人
	go func() {
		if err := manager.Start(ctx); err != nil {
			log.Printf("机器人启动失败: %v", err)
			cancel()
		}
	}()

	// 等待中断信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	select {
	case <-sigChan:
		log.Println("收到中断信号，正在停止机器人...")
	case <-ctx.Done():
		log.Println("上下文已取消，正在停止机器人...")
	}

	// 停止所有机器人
	if err := manager.Stop(); err != nil {
		log.Printf("停止机器人时出错: %v", err)
	}

	log.Println("程序已退出")
}

// initDiscordBot 初始化Discord机器人
func initDiscordBot(manager *core.BotManager, token string) error {
	// 配置
	config := &discord.Config{
		Token:   token,
		GuildID: os.Getenv("DISCORD_GUILD_ID"), // 可选，用于开发测试
		DevMode: os.Getenv("BOT_DEV_MODE") == "true",
	}

	// 创建机器人
	bot, err := discord.NewDiscordBot(config)
	if err != nil {
		return err
	}

	// 注册命令
	bot.RegisterCommand(commands.CreateHelpCommand())
	bot.RegisterCommand(commands.CreateAffixCommand())
	bot.RegisterCommand(commands.CreateStrengthenCommand())

	// 注册到管理器
	return manager.Register(bot)
}