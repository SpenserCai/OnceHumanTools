package core

import (
	"context"
	"fmt"
	"sync"
)

// BotManager 机器人管理器
type BotManager struct {
	bots map[string]Bot
	mu   sync.RWMutex
}

// NewBotManager 创建新的机器人管理器
func NewBotManager() *BotManager {
	return &BotManager{
		bots: make(map[string]Bot),
	}
}

// Register 注册机器人
func (m *BotManager) Register(bot Bot) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	name := bot.GetName()
	if _, exists := m.bots[name]; exists {
		return fmt.Errorf("机器人 %s 已存在", name)
	}

	m.bots[name] = bot
	return nil
}

// Start 启动所有机器人
func (m *BotManager) Start(ctx context.Context) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var wg sync.WaitGroup
	errChan := make(chan error, len(m.bots))
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for name, bot := range m.bots {
		wg.Add(1)
		go func(n string, b Bot) {
			defer wg.Done()
			if err := b.Start(ctx); err != nil {
				errChan <- fmt.Errorf("启动机器人 %s 失败: %w", n, err)
				cancel() // 取消其他机器人
			}
		}(name, bot)
	}

	// 等待所有机器人启动完成或出错
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// 检查错误
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

// Stop 停止所有机器人
func (m *BotManager) Stop() error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var wg sync.WaitGroup
	var firstErr error
	var errMu sync.Mutex

	for name, bot := range m.bots {
		wg.Add(1)
		go func(n string, b Bot) {
			defer wg.Done()
			if err := b.Stop(); err != nil {
				errMu.Lock()
				if firstErr == nil {
					firstErr = fmt.Errorf("停止机器人 %s 失败: %w", n, err)
				}
				errMu.Unlock()
			}
		}(name, bot)
	}

	wg.Wait()
	return firstErr
}

// GetBot 获取机器人
func (m *BotManager) GetBot(name string) (Bot, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	bot, exists := m.bots[name]
	if !exists {
		return nil, fmt.Errorf("机器人 %s 不存在", name)
	}

	return bot, nil
}

// ListBots 列出所有机器人
func (m *BotManager) ListBots() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	names := make([]string, 0, len(m.bots))
	for name := range m.bots {
		names = append(names, name)
	}

	return names
}