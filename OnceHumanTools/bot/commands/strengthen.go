package commands

import (
	"fmt"
	"strconv"
	"strings"
	
	"github.com/oncehuman/tools/internal/services"
)

// StrengthenCommand 强化概率计算命令
type StrengthenCommand struct {
	service *services.StrengthenProbabilityService
}

// NewStrengthenCommand 创建强化概率计算命令
func NewStrengthenCommand() *StrengthenCommand {
	return &StrengthenCommand{
		service: services.NewStrengthenProbabilityService(),
	}
}

// Name 命令名称
func (c *StrengthenCommand) Name() string {
	return "strengthen"
}

// Description 命令描述
func (c *StrengthenCommand) Description() string {
	return "计算模组强化概率"
}

// Usage 命令用法
func (c *StrengthenCommand) Usage() string {
	return "strengthen <初始等级> <目标等级> [模式]"
}

// Execute 执行命令
func (c *StrengthenCommand) Execute(args []string) (string, error) {
	if len(args) < 2 {
		return c.formatUsage(), nil
	}
	
	// 解析初始等级
	initialLevels, err := c.parseLevels(args[0])
	if err != nil {
		return fmt.Sprintf("初始等级格式错误: %s", err.Error()), nil
	}
	
	// 解析目标等级
	targetLevels, err := c.parseLevels(args[1])
	if err != nil {
		return fmt.Sprintf("目标等级格式错误: %s", err.Error()), nil
	}
	
	// 解析模式（可选）
	orderIndependent := true
	if len(args) > 2 {
		mode := strings.ToLower(args[2])
		if mode == "strict" || mode == "位置对应" {
			orderIndependent = false
		}
	}
	
	// 计算概率
	result := c.service.CalculateProbability(initialLevels, targetLevels, orderIndependent, false)
	
	// 检查错误
	if result.Error != "" {
		return result.Error, nil
	}
	
	// 格式化结果
	return c.formatResult(initialLevels, targetLevels, orderIndependent, result), nil
}

// parseLevels 解析等级字符串
func (c *StrengthenCommand) parseLevels(levelStr string) ([]int, error) {
	parts := strings.Split(levelStr, ",")
	if len(parts) != 4 {
		return nil, fmt.Errorf("必须提供4个等级值")
	}
	
	levels := make([]int, 4)
	for i, part := range parts {
		level, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil || level < 1 || level > 5 {
			return nil, fmt.Errorf("等级必须是1-5之间的数字")
		}
		levels[i] = level
	}
	
	return levels, nil
}

// formatUsage 格式化用法说明
func (c *StrengthenCommand) formatUsage() string {
	var sb strings.Builder
	
	sb.WriteString("**强化概率计算器**\n")
	sb.WriteString("```\n")
	sb.WriteString("用法: !oh strengthen <初始等级> <目标等级> [模式]\n\n")
	sb.WriteString("参数说明:\n")
	sb.WriteString("  初始等级: 4个词条的初始等级，用逗号分隔 (1-5)\n")
	sb.WriteString("  目标等级: 4个词条的目标等级，用逗号分隔 (1-5)\n")
	sb.WriteString("  模式(可选): \n")
	sb.WriteString("    - 留空或'free': 顺序无关模式（默认）\n")
	sb.WriteString("    - 'strict': 位置对应模式\n\n")
	sb.WriteString("游戏规则:\n")
	sb.WriteString("  - 一个模组有4个词条\n")
	sb.WriteString("  - 每个词条最低1级，最高5级\n")
	sb.WriteString("  - 总共有5次强化机会\n")
	sb.WriteString("  - 每次随机强化一个未满级的词条\n\n")
	sb.WriteString("示例:\n")
	sb.WriteString("  !oh strengthen 1,1,1,1 2,2,2,2\n")
	sb.WriteString("  !oh strengthen 1,2,3,1 3,4,5,2 strict\n")
	sb.WriteString("```")
	
	return sb.String()
}

// formatResult 格式化计算结果
func (c *StrengthenCommand) formatResult(initialLevels, targetLevels []int, orderIndependent bool, result *services.StrengthenProbabilityResult) string {
	var sb strings.Builder
	
	modeStr := "顺序无关"
	if !orderIndependent {
		modeStr = "位置对应"
	}
	
	sb.WriteString("**强化概率计算结果**\n")
	sb.WriteString("```\n")
	sb.WriteString(fmt.Sprintf("初始等级: [%d, %d, %d, %d]\n", 
		initialLevels[0], initialLevels[1], initialLevels[2], initialLevels[3]))
	sb.WriteString(fmt.Sprintf("目标等级: [%d, %d, %d, %d]\n", 
		targetLevels[0], targetLevels[1], targetLevels[2], targetLevels[3]))
	sb.WriteString(fmt.Sprintf("判断模式: %s\n", modeStr))
	sb.WriteString("─────────────────────────\n")
	sb.WriteString(fmt.Sprintf("成功概率: %.4f%%\n", result.ProbabilityPercent))
	sb.WriteString(fmt.Sprintf("成功路径数: %d\n", result.SuccessfulOutcomes))
	sb.WriteString(fmt.Sprintf("总路径数: %d\n", result.TotalOutcomes))
	sb.WriteString(fmt.Sprintf("精确概率: %.6f\n", result.Probability))
	sb.WriteString("```")
	
	// 添加提示
	if result.ProbabilityPercent < 10 {
		sb.WriteString("\n💡 提示: 成功率较低，建议调整目标或准备更多资源")
	} else if result.ProbabilityPercent > 75 {
		sb.WriteString("\n✨ 提示: 成功率很高，祝你好运！")
	}
	
	return sb.String()
}