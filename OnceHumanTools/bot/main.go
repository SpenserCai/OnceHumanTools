package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/oncehuman/tools/bot/commands"
	"github.com/oncehuman/tools/bot/handlers"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("未找到.env文件，使用系统环境变量")
	}

	// 获取Discord token
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("请设置DISCORD_TOKEN环境变量")
	}

	// 创建Discord会话
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("创建Discord会话失败:", err)
	}

	// 注册命令
	cmdManager := commands.NewCommandManager()
	cmdManager.RegisterCommands()

	// 创建消息处理器
	msgHandler := handlers.NewMessageHandler(cmdManager)

	// 添加事件处理器
	dg.AddHandler(msgHandler.HandleMessage)
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("机器人已登录: %v#%v", s.State.User.Username, s.State.User.Discriminator)
		
		// 设置机器人状态
		s.UpdateGameStatus(0, "!oh help | OnceHuman工具集")
	})

	// 打开连接
	err = dg.Open()
	if err != nil {
		log.Fatal("打开Discord连接失败:", err)
	}

	// 等待中断信号
	fmt.Println("机器人正在运行，按 Ctrl+C 停止")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// 关闭连接
	dg.Close()
}