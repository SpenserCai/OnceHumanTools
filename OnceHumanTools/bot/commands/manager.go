package commands

import (
	"strings"
)

// Command 命令接口
type Command interface {
	Name() string
	Description() string
	Usage() string
	Execute(args []string) (string, error)
}

// CommandManager 命令管理器
type CommandManager struct {
	commands map[string]Command
	prefix   string
}

// NewCommandManager 创建命令管理器
func NewCommandManager() *CommandManager {
	return &CommandManager{
		commands: make(map[string]Command),
		prefix:   "!oh",
	}
}

// RegisterCommands 注册所有命令
func (m *CommandManager) RegisterCommands() {
	// 注册帮助命令
	m.Register(NewHelpCommand(m))
	
	// 注册词条概率命令
	m.Register(NewAffixCommand())
	
	// 注册强化概率命令
	m.Register(NewStrengthenCommand())
}

// Register 注册命令
func (m *CommandManager) Register(cmd Command) {
	m.commands[strings.ToLower(cmd.Name())] = cmd
}

// GetCommand 获取命令
func (m *CommandManager) GetCommand(name string) (Command, bool) {
	cmd, ok := m.commands[strings.ToLower(name)]
	return cmd, ok
}

// GetAllCommands 获取所有命令
func (m *CommandManager) GetAllCommands() map[string]Command {
	return m.commands
}

// GetPrefix 获取命令前缀
func (m *CommandManager) GetPrefix() string {
	return m.prefix
}

// ParseMessage 解析消息
func (m *CommandManager) ParseMessage(content string) (cmdName string, args []string, ok bool) {
	// 检查前缀
	if !strings.HasPrefix(content, m.prefix) {
		return "", nil, false
	}
	
	// 移除前缀并分割
	content = strings.TrimPrefix(content, m.prefix)
	content = strings.TrimSpace(content)
	
	if content == "" {
		return "", nil, false
	}
	
	parts := strings.Fields(content)
	if len(parts) == 0 {
		return "", nil, false
	}
	
	cmdName = parts[0]
	if len(parts) > 1 {
		args = parts[1:]
	}
	
	return cmdName, args, true
}