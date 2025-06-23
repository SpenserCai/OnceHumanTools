package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/oncehuman/tools/bot/platforms/discord"
	"github.com/oncehuman/tools/internal/services"
)

// CreateAffixCommand 创建词条概率计算命令
func CreateAffixCommand() *discord.SlashCommand {
	return &discord.SlashCommand{
		Command: &discordgo.ApplicationCommand{
			Name:        "affix",
			Description: "计算模组词条概率",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "slots",
					Description: "词条数量 (1-10)",
					Required:    true,
					MinValue:    &[]float64{1}[0],
					MaxValue:    10,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "targets",
					Description: "目标词条ID列表，用逗号分隔 (例如: 1,4,5,6)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "show_combinations",
					Description: "是否显示详细组合",
					Required:    false,
				},
			},
		},
		Handler: handleAffixCommand,
	}
}

// handleAffixCommand 处理词条概率计算命令
func handleAffixCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resp := discord.CreateResponse(s, i)
	
	// 延迟响应，因为计算可能需要时间
	if err := resp.Defer(); err != nil {
		return
	}

	// 获取参数
	options := i.ApplicationCommandData().Options
	var slotCount int
	var targetStr string
	showCombinations := false
	
	for _, opt := range options {
		switch opt.Name {
		case "slots":
			slotCount = int(opt.IntValue())
		case "targets":
			targetStr = opt.StringValue()
		case "show_combinations":
			if opt.BoolValue() != nil {
				showCombinations = *opt.BoolValue()
			}
		}
	}

	// 解析目标词条ID
	targetIDs := parseTargetIDs(targetStr)
	if len(targetIDs) == 0 {
		resp.SendError(fmt.Errorf("无效的目标词条ID格式"))
		return
	}

	// 计算概率
	service := services.NewAffixProbabilityService()
	result := service.CalculateProbability(slotCount, targetIDs, showCombinations)

	// 检查错误
	if result.Error != "" {
		resp.SendError(fmt.Errorf(result.Error))
		return
	}

	// 构建响应
	embed := buildAffixResultEmbed(slotCount, result)
	resp.SendEmbed(embed)
}

// parseTargetIDs 解析目标词条ID
func parseTargetIDs(str string) []int {
	parts := strings.Split(str, ",")
	var ids []int
	
	for _, part := range parts {
		part = strings.TrimSpace(part)
		var id int
		if _, err := fmt.Sscanf(part, "%d", &id); err == nil && id >= 1 && id <= 10 {
			ids = append(ids, id)
		}
	}
	
	return ids
}

// buildAffixResultEmbed 构建结果嵌入消息
func buildAffixResultEmbed(slotCount int, result *services.AffixProbabilityResult) *discordgo.MessageEmbed {
	// 词条名称映射
	affixNames := map[int]string{
		1: "异常伤害", 2: "弹匣容量", 3: "换弹速度加成",
		4: "对普通敌人伤害", 5: "对精英敌人伤害", 6: "对上位者伤害",
		7: "最大生命值", 8: "头部受伤减免", 9: "枪械伤害减免", 10: "异常伤害减免",
	}

	// 构建目标词条名称列表
	var targetNames []string
	for _, id := range result.TargetRange {
		if name, ok := affixNames[id]; ok {
			targetNames = append(targetNames, name)
		}
	}

	// 选择颜色
	color := 0x00FF88 // 绿色
	if result.ProbabilityPercent < 10 {
		color = 0xFF0044 // 红色
	} else if result.ProbabilityPercent < 30 {
		color = 0xFFAA00 // 橙色
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📊 词条概率计算结果",
		Description: fmt.Sprintf("计算 %d 个词条位中出现指定词条的概率", slotCount),
		Color:       color,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📌 目标词条",
				Value:  strings.Join(targetNames, ", "),
				Inline: false,
			},
			{
				Name:   "🎲 出现概率",
				Value:  fmt.Sprintf("**%.4f%%**", result.ProbabilityPercent),
				Inline: true,
			},
			{
				Name:   "📈 精确概率",
				Value:  fmt.Sprintf("%.6f", result.Probability),
				Inline: true,
			},
			{
				Name:   "🔢 满足条件的组合数",
				Value:  fmt.Sprintf("%d", result.ValidCombinations),
				Inline: true,
			},
			{
				Name:   "🔢 总组合数",
				Value:  fmt.Sprintf("%d", result.TotalCombinations),
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "OnceHuman工具集",
		},
		Timestamp: discordgo.NowTimestamp(),
	}

	// 添加组合示例
	if len(result.Combinations) > 0 && len(result.Combinations) <= 10 {
		var comboStrs []string
		for i, combo := range result.Combinations {
			if i >= 5 { // 最多显示5个
				comboStrs = append(comboStrs, fmt.Sprintf("... 还有 %d 种组合", len(result.Combinations)-5))
				break
			}
			var names []string
			for _, id := range combo {
				if name, ok := affixNames[id]; ok {
					names = append(names, name)
				}
			}
			comboStrs = append(comboStrs, fmt.Sprintf("%d. %s", i+1, strings.Join(names, " + ")))
		}
		
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "📝 可能的组合",
			Value:  strings.Join(comboStrs, "\n"),
			Inline: false,
		})
	}

	return embed
}

// GetAffixListChoices 获取词条选择列表（用于自动完成）
func GetAffixListChoices() []*discordgo.ApplicationCommandOptionChoice {
	return []*discordgo.ApplicationCommandOptionChoice{
		{Name: "异常伤害", Value: "1"},
		{Name: "弹匣容量", Value: "2"},
		{Name: "换弹速度加成", Value: "3"},
		{Name: "对普通敌人伤害", Value: "4"},
		{Name: "对精英敌人伤害", Value: "5"},
		{Name: "对上位者伤害", Value: "6"},
		{Name: "最大生命值", Value: "7"},
		{Name: "头部受伤减免", Value: "8"},
		{Name: "枪械伤害减免", Value: "9"},
		{Name: "异常伤害减免", Value: "10"},
	}
}