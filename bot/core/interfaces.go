package core

import "context"

// Bot 机器人接口
type Bot interface {
	// Start 启动机器人
	Start(ctx context.Context) error
	// Stop 停止机器人
	Stop() error
	// GetName 获取机器人名称
	GetName() string
	// IsRunning 是否正在运行
	IsRunning() bool
}

// Command 命令接口
type Command interface {
	// GetName 获取命令名称
	GetName() string
	// GetDescription 获取命令描述
	GetDescription() string
	// Execute 执行命令
	Execute(ctx *CommandContext) error
}

// CommandContext 命令上下文
type CommandContext struct {
	// 用户ID
	UserID string
	// 用户名
	Username string
	// 频道ID
	ChannelID string
	// 服务器ID
	GuildID string
	// 命令参数
	Args map[string]interface{}
	// 原始数据（平台特定）
	Raw interface{}
}

// Response 响应接口
type Response interface {
	// SendText 发送文本消息
	SendText(text string) error
	// SendEmbed 发送嵌入消息
	SendEmbed(embed *Embed) error
	// SendError 发送错误消息
	SendError(err error) error
	// Defer 延迟响应
	Defer() error
	// FollowUp 跟进消息
	FollowUp(text string) error
}

// Embed 嵌入消息
type Embed struct {
	Title       string
	Description string
	Color       int
	Fields      []EmbedField
	Footer      *EmbedFooter
	Thumbnail   *EmbedImage
	Image       *EmbedImage
}

// EmbedField 嵌入字段
type EmbedField struct {
	Name   string
	Value  string
	Inline bool
}

// EmbedFooter 嵌入页脚
type EmbedFooter struct {
	Text    string
	IconURL string
}

// EmbedImage 嵌入图片
type EmbedImage struct {
	URL    string
	Width  int
	Height int
}

// Manager 机器人管理器接口
type Manager interface {
	// Register 注册机器人
	Register(bot Bot) error
	// Start 启动所有机器人
	Start(ctx context.Context) error
	// Stop 停止所有机器人
	Stop() error
	// GetBot 获取机器人
	GetBot(name string) (Bot, error)
	// ListBots 列出所有机器人
	ListBots() []string
}