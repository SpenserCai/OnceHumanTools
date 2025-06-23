package commands

import (
	"fmt"
	"strings"
)

// HelpCommand 帮助命令
type HelpCommand struct {
	manager *CommandManager
}

// NewHelpCommand 创建帮助命令
func NewHelpCommand(manager *CommandManager) *HelpCommand {
	return &HelpCommand{
		manager: manager,
	}
}

// Name 命令名称
func (c *HelpCommand) Name() string {
	return "help"
}

// Description 命令描述
func (c *HelpCommand) Description() string {
	return "显示帮助信息"
}

// Usage 命令用法
func (c *HelpCommand) Usage() string {
	return "help [命令名]"
}

// Execute 执行命令
func (c *HelpCommand) Execute(args []string) (string, error) {
	if len(args) > 0 {
		// 显示特定命令的帮助
		cmdName := args[0]
		if cmd, ok := c.manager.GetCommand(cmdName); ok {
			return c.formatCommandHelp(cmd), nil
		}
		return fmt.Sprintf("未找到命令: %s", cmdName), nil
	}
	
	// 显示所有命令的帮助
	return c.formatAllCommands(), nil
}

// formatCommandHelp 格式化单个命令的帮助信息
func (c *HelpCommand) formatCommandHelp(cmd Command) string {
	var sb strings.Builder
	
	sb.WriteString("```\n")
	sb.WriteString(fmt.Sprintf("命令: %s %s\n", c.manager.GetPrefix(), cmd.Name()))
	sb.WriteString(fmt.Sprintf("描述: %s\n", cmd.Description()))
	sb.WriteString(fmt.Sprintf("用法: %s %s\n", c.manager.GetPrefix(), cmd.Usage()))
	sb.WriteString("```")
	
	return sb.String()
}

// formatAllCommands 格式化所有命令的帮助信息
func (c *HelpCommand) formatAllCommands() string {
	var sb strings.Builder
	
	sb.WriteString("**OnceHuman工具集 - Discord机器人**\n")
	sb.WriteString("```\n")
	sb.WriteString("可用命令:\n\n")
	
	commands := c.manager.GetAllCommands()
	for name, cmd := range commands {
		sb.WriteString(fmt.Sprintf("%s %-15s - %s\n", 
			c.manager.GetPrefix(), 
			name, 
			cmd.Description()))
	}
	
	sb.WriteString("\n使用 ")
	sb.WriteString(c.manager.GetPrefix())
	sb.WriteString(" help <命令名> 查看详细用法\n")
	sb.WriteString("```")
	
	return sb.String()
}