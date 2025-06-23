package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/oncehuman/tools/bot/commands"
)

// MessageHandler 消息处理器
type MessageHandler struct {
	cmdManager *commands.CommandManager
}

// NewMessageHandler 创建消息处理器
func NewMessageHandler(cmdManager *commands.CommandManager) *MessageHandler {
	return &MessageHandler{
		cmdManager: cmdManager,
	}
}

// HandleMessage 处理消息事件
func (h *MessageHandler) HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// 忽略机器人自己的消息
	if m.Author.ID == s.State.User.ID {
		return
	}
	
	// 忽略机器人的消息
	if m.Author.Bot {
		return
	}
	
	// 解析消息
	cmdName, args, ok := h.cmdManager.ParseMessage(m.Content)
	if !ok {
		return
	}
	
	// 获取命令
	cmd, ok := h.cmdManager.GetCommand(cmdName)
	if !ok {
		// 命令不存在，发送帮助信息
		h.sendMessage(s, m.ChannelID, "未知命令。使用 `"+h.cmdManager.GetPrefix()+" help` 查看帮助。")
		return
	}
	
	// 在频道中显示"正在输入"状态
	s.ChannelTyping(m.ChannelID)
	
	// 执行命令
	result, err := cmd.Execute(args)
	if err != nil {
		log.Printf("命令执行错误: %v", err)
		h.sendMessage(s, m.ChannelID, "命令执行失败: "+err.Error())
		return
	}
	
	// 发送结果
	h.sendMessage(s, m.ChannelID, result)
}

// sendMessage 发送消息
func (h *MessageHandler) sendMessage(s *discordgo.Session, channelID, content string) {
	// 如果消息太长，分段发送
	if len(content) > 2000 {
		parts := h.splitMessage(content, 2000)
		for _, part := range parts {
			_, err := s.ChannelMessageSend(channelID, part)
			if err != nil {
				log.Printf("发送消息失败: %v", err)
			}
		}
	} else {
		_, err := s.ChannelMessageSend(channelID, content)
		if err != nil {
			log.Printf("发送消息失败: %v", err)
		}
	}
}

// splitMessage 分割长消息
func (h *MessageHandler) splitMessage(content string, maxLength int) []string {
	if len(content) <= maxLength {
		return []string{content}
	}
	
	var parts []string
	lines := strings.Split(content, "\n")
	
	currentPart := ""
	for _, line := range lines {
		// 如果单行就超过最大长度，强制分割
		if len(line) > maxLength {
			if currentPart != "" {
				parts = append(parts, currentPart)
				currentPart = ""
			}
			
			// 分割长行
			for len(line) > maxLength {
				parts = append(parts, line[:maxLength])
				line = line[maxLength:]
			}
			if line != "" {
				currentPart = line
			}
			continue
		}
		
		// 如果加上这一行会超过最大长度，先保存当前部分
		if len(currentPart)+len(line)+1 > maxLength {
			if currentPart != "" {
				parts = append(parts, currentPart)
			}
			currentPart = line
		} else {
			if currentPart != "" {
				currentPart += "\n"
			}
			currentPart += line
		}
	}
	
	// 保存最后一部分
	if currentPart != "" {
		parts = append(parts, currentPart)
	}
	
	return parts
}

// HandleError 处理错误
func (h *MessageHandler) HandleError(s *discordgo.Session, channelID string, err error) {
	log.Printf("处理错误: %v", err)
	h.sendMessage(s, channelID, "❌ 发生错误: "+err.Error())
}