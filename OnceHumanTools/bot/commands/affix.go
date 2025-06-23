package commands

import (
	"fmt"
	"strconv"
	"strings"
	
	"github.com/oncehuman/tools/internal/services"
)

// AffixCommand 词条概率计算命令
type AffixCommand struct {
	service *services.AffixProbabilityService
}

// NewAffixCommand 创建词条概率计算命令
func NewAffixCommand() *AffixCommand {
	return &AffixCommand{
		service: services.NewAffixProbabilityService(),
	}
}

// Name 命令名称
func (c *AffixCommand) Name() string {
	return "affix"
}

// Description 命令描述
func (c *AffixCommand) Description() string {
	return "计算模组词条概率"
}

// Usage 命令用法
func (c *AffixCommand) Usage() string {
	return "affix <词条数量> <目标词条ID列表>"
}

// Execute 执行命令
func (c *AffixCommand) Execute(args []string) (string, error) {
	if len(args) < 2 {
		return c.formatUsage(), nil
	}
	
	// 解析词条数量
	slotCount, err := strconv.Atoi(args[0])
	if err != nil || slotCount < 1 || slotCount > 10 {
		return "词条数量必须是1-10之间的数字", nil
	}
	
	// 解析目标词条ID
	targetIDs := []int{}
	idStrs := strings.Split(args[1], ",")
	for _, idStr := range idStrs {
		id, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err != nil || id < 1 || id > 10 {
			return fmt.Sprintf("无效的词条ID: %s", idStr), nil
		}
		targetIDs = append(targetIDs, id)
	}
	
	if len(targetIDs) == 0 {
		return "请提供至少一个目标词条ID", nil
	}
	
	// 计算概率
	result := c.service.CalculateProbability(slotCount, targetIDs, false)
	
	// 检查错误
	if result.Error != "" {
		return result.Error, nil
	}
	
	// 格式化结果
	return c.formatResult(slotCount, targetIDs, result), nil
}

// formatUsage 格式化用法说明
func (c *AffixCommand) formatUsage() string {
	var sb strings.Builder
	
	sb.WriteString("**词条概率计算器**\n")
	sb.WriteString("```\n")
	sb.WriteString("用法: !oh affix <词条数量> <目标词条ID列表>\n\n")
	sb.WriteString("参数说明:\n")
	sb.WriteString("  词条数量: 1-10之间的数字\n")
	sb.WriteString("  目标词条ID: 用逗号分隔的词条ID (1-10)\n\n")
	sb.WriteString("词条列表:\n")
	sb.WriteString("  1. 异常伤害\n")
	sb.WriteString("  2. 弹匣容量\n")
	sb.WriteString("  3. 换弹速度加成\n")
	sb.WriteString("  4. 对普通敌人伤害\n")
	sb.WriteString("  5. 对精英敌人伤害\n")
	sb.WriteString("  6. 对上位者伤害\n")
	sb.WriteString("  7. 最大生命值\n")
	sb.WriteString("  8. 头部受伤减免\n")
	sb.WriteString("  9. 枪械伤害减免\n")
	sb.WriteString(" 10. 异常伤害减免\n\n")
	sb.WriteString("示例: !oh affix 3 1,4,5,6\n")
	sb.WriteString("```")
	
	return sb.String()
}

// formatResult 格式化计算结果
func (c *AffixCommand) formatResult(slotCount int, targetIDs []int, result *services.AffixProbabilityResult) string {
	var sb strings.Builder
	
	// 获取词条名称映射
	affixNames := map[int]string{
		1: "异常伤害", 2: "弹匣容量", 3: "换弹速度加成",
		4: "对普通敌人伤害", 5: "对精英敌人伤害", 6: "对上位者伤害",
		7: "最大生命值", 8: "头部受伤减免", 9: "枪械伤害减免", 10: "异常伤害减免",
	}
	
	// 构建目标词条名称列表
	targetNames := []string{}
	for _, id := range result.TargetRange {
		if name, ok := affixNames[id]; ok {
			targetNames = append(targetNames, name)
		}
	}
	
	sb.WriteString("**词条概率计算结果**\n")
	sb.WriteString("```\n")
	sb.WriteString(fmt.Sprintf("词条数量: %d\n", slotCount))
	sb.WriteString(fmt.Sprintf("目标词条: %s\n", strings.Join(targetNames, ", ")))
	sb.WriteString("─────────────────────────\n")
	sb.WriteString(fmt.Sprintf("出现概率: %.4f%%\n", result.ProbabilityPercent))
	sb.WriteString(fmt.Sprintf("满足条件的组合数: %d\n", result.ValidCombinations))
	sb.WriteString(fmt.Sprintf("总组合数: %d\n", result.TotalCombinations))
	sb.WriteString(fmt.Sprintf("精确概率: %.6f\n", result.Probability))
	sb.WriteString("```")
	
	return sb.String()
}