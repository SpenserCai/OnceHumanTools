package discord

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/bwmarrin/discordgo"
)

// DiscordBot Discord机器人
type DiscordBot struct {
	session  *discordgo.Session
	config   *Config
	commands map[string]*SlashCommand
	handlers map[string]CommandHandler
	running  bool
	mu       sync.RWMutex
}

// Config Discord配置
type Config struct {
	Token    string
	AppID    string
	GuildID  string // 为空则注册全局命令
	DevMode  bool   // 开发模式，只在特定服务器注册命令
}

// CommandHandler 命令处理器
type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

// SlashCommand 斜杠命令定义
type SlashCommand struct {
	Command *discordgo.ApplicationCommand
	Handler CommandHandler
}

// NewDiscordBot 创建Discord机器人
func NewDiscordBot(config *Config) (*DiscordBot, error) {
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, fmt.Errorf("创建Discord会话失败: %w", err)
	}

	bot := &DiscordBot{
		session:  dg,
		config:   config,
		commands: make(map[string]*SlashCommand),
		handlers: make(map[string]CommandHandler),
	}

	// 设置事件处理器
	dg.AddHandler(bot.handleInteraction)
	dg.AddHandler(bot.handleReady)

	return bot, nil
}

// Start 启动机器人
func (b *DiscordBot) Start(ctx context.Context) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.running {
		return fmt.Errorf("机器人已在运行")
	}

	// 打开连接
	if err := b.session.Open(); err != nil {
		return fmt.Errorf("打开Discord连接失败: %w", err)
	}

	b.running = true
	log.Println("Discord机器人已启动")

	// 注册命令
	if err := b.registerCommands(); err != nil {
		log.Printf("注册命令失败: %v", err)
	}

	// 等待上下文取消
	<-ctx.Done()
	return b.Stop()
}

// Stop 停止机器人
func (b *DiscordBot) Stop() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.running {
		return nil
	}

	// 清理命令
	b.cleanupCommands()

	// 关闭连接
	if err := b.session.Close(); err != nil {
		return fmt.Errorf("关闭Discord连接失败: %w", err)
	}

	b.running = false
	log.Println("Discord机器人已停止")
	return nil
}

// GetName 获取机器人名称
func (b *DiscordBot) GetName() string {
	return "discord"
}

// IsRunning 是否正在运行
func (b *DiscordBot) IsRunning() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.running
}

// RegisterCommand 注册命令
func (b *DiscordBot) RegisterCommand(cmd *SlashCommand) {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	b.commands[cmd.Command.Name] = cmd
	b.handlers[cmd.Command.Name] = cmd.Handler
}

// registerCommands 注册所有命令到Discord
func (b *DiscordBot) registerCommands() error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, cmd := range b.commands {
		var err error
		if b.config.GuildID != "" {
			// 注册到特定服务器（立即生效）
			_, err = b.session.ApplicationCommandCreate(b.session.State.User.ID, b.config.GuildID, cmd.Command)
		} else {
			// 注册全局命令（可能需要1小时生效）
			_, err = b.session.ApplicationCommandCreate(b.session.State.User.ID, "", cmd.Command)
		}
		
		if err != nil {
			log.Printf("注册命令 %s 失败: %v", cmd.Command.Name, err)
		} else {
			log.Printf("成功注册命令: %s", cmd.Command.Name)
		}
	}

	return nil
}

// cleanupCommands 清理命令
func (b *DiscordBot) cleanupCommands() {
	if b.config.DevMode {
		// 开发模式下清理命令
		var guildID string
		if b.config.GuildID != "" {
			guildID = b.config.GuildID
		}

		cmds, err := b.session.ApplicationCommands(b.session.State.User.ID, guildID)
		if err != nil {
			log.Printf("获取命令列表失败: %v", err)
			return
		}

		for _, cmd := range cmds {
			err := b.session.ApplicationCommandDelete(b.session.State.User.ID, guildID, cmd.ID)
			if err != nil {
				log.Printf("删除命令 %s 失败: %v", cmd.Name, err)
			}
		}
	}
}

// handleReady 处理就绪事件
func (b *DiscordBot) handleReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Discord机器人已登录: %s#%s", s.State.User.Username, s.State.User.Discriminator)
	
	// 设置状态
	s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "OnceHuman工具集",
				Type: discordgo.ActivityTypeGame,
			},
		},
		Status: string(discordgo.StatusOnline),
	})
}

// handleInteraction 处理交互事件
func (b *DiscordBot) handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// 只处理应用命令
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	// 获取命令数据
	data := i.ApplicationCommandData()
	
	// 查找处理器
	b.mu.RLock()
	handler, ok := b.handlers[data.Name]
	b.mu.RUnlock()

	if !ok {
		log.Printf("未找到命令处理器: %s", data.Name)
		return
	}

	// 执行处理器
	handler(s, i)
}

// CreateResponse 创建响应助手
func CreateResponse(s *discordgo.Session, i *discordgo.InteractionCreate) *InteractionResponse {
	return &InteractionResponse{
		session:     s,
		interaction: i,
	}
}

// InteractionResponse 交互响应助手
type InteractionResponse struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
	deferred    bool
}

// Defer 延迟响应
func (r *InteractionResponse) Defer() error {
	err := r.session.InteractionRespond(r.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err == nil {
		r.deferred = true
	}
	return err
}

// SendText 发送文本响应
func (r *InteractionResponse) SendText(content string) error {
	if r.deferred {
		_, err := r.session.InteractionResponseEdit(r.interaction.Interaction, &discordgo.WebhookEdit{
			Content: &content,
		})
		return err
	}

	return r.session.InteractionRespond(r.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

// SendEmbed 发送嵌入消息
func (r *InteractionResponse) SendEmbed(embed *discordgo.MessageEmbed) error {
	if r.deferred {
		_, err := r.session.InteractionResponseEdit(r.interaction.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{embed},
		})
		return err
	}

	return r.session.InteractionRespond(r.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

// SendError 发送错误消息
func (r *InteractionResponse) SendError(err error) error {
	embed := &discordgo.MessageEmbed{
		Title:       "❌ 错误",
		Description: err.Error(),
		Color:       0xFF0000,
	}
	return r.SendEmbed(embed)
}

// FollowUp 发送跟进消息
func (r *InteractionResponse) FollowUp(content string) error {
	_, err := r.session.FollowupMessageCreate(r.interaction.Interaction, true, &discordgo.WebhookParams{
		Content: content,
	})
	return err
}