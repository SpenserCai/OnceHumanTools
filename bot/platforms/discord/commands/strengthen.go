package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SpenserCai/OnceHumanTools/backend/internal/services"
	"github.com/SpenserCai/OnceHumanTools/bot/platforms/discord"
	"github.com/bwmarrin/discordgo"
)

// CreateStrengthenCommand 创建强化概率计算命令
func CreateStrengthenCommand() *discord.SlashCommand {
	return &discord.SlashCommand{
		Command: &discordgo.ApplicationCommand{
			Name:        "strengthen",
			Description: "计算词条强化概率",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "single",
					Description: "计算单个词条强化概率",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "affix_id",
							Description: "词条ID (1-10)",
							Required:    true,
							MinValue:    &[]float64{1}[0],
							MaxValue:    10,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "current_level",
							Description: "当前等级 (0-5)",
							Required:    true,
							MinValue:    &[]float64{0}[0],
							MaxValue:    5,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "target_level",
							Description: "目标等级 (1-5)",
							Required:    true,
							MinValue:    &[]float64{1}[0],
							MaxValue:    5,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "slot_count",
							Description: "词条数量 (1-10)",
							Required:    true,
							MinValue:    &[]float64{1}[0],
							MaxValue:    10,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "tries",
							Description: "强化次数 (1-999)",
							Required:    true,
							MinValue:    &[]float64{1}[0],
							MaxValue:    999,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "multi",
					Description: "计算多个词条强化概率",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "targets",
							Description: "目标格式: 词条ID:当前等级:目标等级，用逗号分隔 (例如: 1:0:3,4:1:5)",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "slot_count",
							Description: "词条数量 (1-10)",
							Required:    true,
							MinValue:    &[]float64{1}[0],
							MaxValue:    10,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "tries",
							Description: "强化次数 (1-999)",
							Required:    true,
							MinValue:    &[]float64{1}[0],
							MaxValue:    999,
						},
					},
				},
			},
		},
		Handler: handleStrengthenCommand,
	}
}

// handleStrengthenCommand 处理强化概率计算命令
func handleStrengthenCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resp := discord.CreateResponse(s, i)

	// 延迟响应
	if err := resp.Defer(); err != nil {
		return
	}

	options := i.ApplicationCommandData().Options
	subCommand := options[0]

	switch subCommand.Name {
	case "single":
		handleSingleStrengthen(resp, subCommand.Options)
	case "multi":
		handleMultiStrengthen(resp, subCommand.Options)
	}
}

// handleSingleStrengthen 处理单个词条强化
func handleSingleStrengthen(resp *discord.InteractionResponse, options []*discordgo.ApplicationCommandInteractionDataOption) {
	// 解析参数
	var affixID, currentLevel, targetLevel, slotCount, tries int

	for _, opt := range options {
		switch opt.Name {
		case "affix_id":
			affixID = int(opt.IntValue())
		case "current_level":
			currentLevel = int(opt.IntValue())
		case "target_level":
			targetLevel = int(opt.IntValue())
		case "slot_count":
			slotCount = int(opt.IntValue())
		case "tries":
			tries = int(opt.IntValue())
		}
	}

	// 验证参数
	if targetLevel <= currentLevel {
		resp.SendError(fmt.Errorf("目标等级必须大于当前等级"))
		return
	}

	// 创建目标
	targets := []services.StrengthenTarget{
		{
			AffixID:      affixID,
			CurrentLevel: currentLevel,
			TargetLevel:  targetLevel,
		},
	}

	// 计算概率
	service := services.NewStrengthenProbabilityService()
	result := service.CalculateStrengthenProbability(targets, slotCount, tries)

	// 构建响应
	embed := buildSingleStrengthenResultEmbed(result, slotCount, tries)
	resp.SendEmbed(embed)
}

// handleMultiStrengthen 处理多个词条强化
func handleMultiStrengthen(resp *discord.InteractionResponse, options []*discordgo.ApplicationCommandInteractionDataOption) {
	// 解析参数
	var targetsStr string
	var slotCount, tries int

	for _, opt := range options {
		switch opt.Name {
		case "targets":
			targetsStr = opt.StringValue()
		case "slot_count":
			slotCount = int(opt.IntValue())
		case "tries":
			tries = int(opt.IntValue())
		}
	}

	// 解析目标
	targets, err := parseStrengthenTargets(targetsStr)
	if err != nil {
		resp.SendError(fmt.Errorf("解析目标失败: %v", err))
		return
	}

	if len(targets) == 0 {
		resp.SendError(fmt.Errorf("至少需要一个目标"))
		return
	}

	// 计算概率
	service := services.NewStrengthenProbabilityService()
	result := service.CalculateStrengthenProbability(targets, slotCount, tries)

	// 构建响应
	embed := buildMultiStrengthenResultEmbed(result, slotCount, tries)
	resp.SendEmbed(embed)
}

// parseStrengthenTargets 解析强化目标
func parseStrengthenTargets(str string) ([]services.StrengthenTarget, error) {
	var targets []services.StrengthenTarget
	parts := strings.Split(str, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		values := strings.Split(part, ":")
		if len(values) != 3 {
			return nil, fmt.Errorf("无效的格式: %s", part)
		}

		affixID, err := strconv.Atoi(values[0])
		if err != nil || affixID < 1 || affixID > 10 {
			return nil, fmt.Errorf("无效的词条ID: %s", values[0])
		}

		currentLevel, err := strconv.Atoi(values[1])
		if err != nil || currentLevel < 0 || currentLevel > 5 {
			return nil, fmt.Errorf("无效的当前等级: %s", values[1])
		}

		targetLevel, err := strconv.Atoi(values[2])
		if err != nil || targetLevel < 1 || targetLevel > 5 || targetLevel <= currentLevel {
			return nil, fmt.Errorf("无效的目标等级: %s", values[2])
		}

		targets = append(targets, services.StrengthenTarget{
			AffixID:      affixID,
			CurrentLevel: currentLevel,
			TargetLevel:  targetLevel,
		})
	}

	return targets, nil
}

// buildSingleStrengthenResultEmbed 构建单个词条强化结果
func buildSingleStrengthenResultEmbed(result *services.StrengthenProbabilityResult, slotCount, tries int) *discordgo.MessageEmbed {
	if len(result.Results) == 0 {
		return &discordgo.MessageEmbed{
			Title:       "❌ 错误",
			Description: "计算失败",
			Color:       0xFF0000,
		}
	}

	// 词条名称映射
	affixNames := map[int]string{
		1: "异常伤害", 2: "弹匣容量", 3: "换弹速度加成",
		4: "对普通敌人伤害", 5: "对精英敌人伤害", 6: "对上位者伤害",
		7: "最大生命值", 8: "头部受伤减免", 9: "枪械伤害减免", 10: "异常伤害减免",
	}

	target := result.Results[0]
	affixName := affixNames[target.AffixID]

	// 选择颜色
	color := 0x00FF88 // 绿色
	if target.SuccessRate < 0.1 {
		color = 0xFF0044 // 红色
	} else if target.SuccessRate < 0.3 {
		color = 0xFFAA00 // 橙色
	}

	// 构建路径描述
	var pathDesc []string
	for _, step := range target.Path {
		pathDesc = append(pathDesc, fmt.Sprintf("Lv%d → Lv%d (%.2f%%)",
			step.FromLevel, step.ToLevel, step.Probability*100))
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🎯 强化概率计算结果",
		Description: fmt.Sprintf("词条: **%s**", affixName),
		Color:       color,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "📊 基础信息",
				Value: fmt.Sprintf("当前等级: Lv%d\n目标等级: Lv%d\n词条数量: %d\n强化次数: %d",
					target.CurrentLevel, target.TargetLevel, slotCount, tries),
				Inline: true,
			},
			{
				Name:   "🎲 成功率",
				Value:  fmt.Sprintf("**%.4f%%**", target.SuccessRate*100),
				Inline: true,
			},
			{
				Name:   "📈 强化路径",
				Value:  strings.Join(pathDesc, "\n"),
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "OnceHuman工具集",
		},
		Timestamp: discordgo.NowTimestamp(),
	}

	// 添加期望值
	if target.ExpectedStrengthens > 0 {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "💰 期望强化次数",
			Value:  fmt.Sprintf("%.2f 次", target.ExpectedStrengthens),
			Inline: true,
		})
	}

	return embed
}

// buildMultiStrengthenResultEmbed 构建多个词条强化结果
func buildMultiStrengthenResultEmbed(result *services.StrengthenProbabilityResult, slotCount, tries int) *discordgo.MessageEmbed {
	// 词条名称映射
	affixNames := map[int]string{
		1: "异常伤害", 2: "弹匣容量", 3: "换弹速度加成",
		4: "对普通敌人伤害", 5: "对精英敌人伤害", 6: "对上位者伤害",
		7: "最大生命值", 8: "头部受伤减免", 9: "枪械伤害减免", 10: "异常伤害减免",
	}

	// 选择颜色
	color := 0x00FF88 // 绿色
	if result.TotalProbability < 0.1 {
		color = 0xFF0044 // 红色
	} else if result.TotalProbability < 0.3 {
		color = 0xFFAA00 // 橙色
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🎯 多词条强化概率计算结果",
		Description: fmt.Sprintf("同时达成所有目标的概率"),
		Color:       color,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📊 总体成功率",
				Value:  fmt.Sprintf("**%.4f%%**", result.TotalProbability*100),
				Inline: false,
			},
			{
				Name:   "🎰 词条数量",
				Value:  fmt.Sprintf("%d", slotCount),
				Inline: true,
			},
			{
				Name:   "🔄 强化次数",
				Value:  fmt.Sprintf("%d", tries),
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "OnceHuman工具集",
		},
		Timestamp: discordgo.NowTimestamp(),
	}

	// 添加各词条详情
	var details []string
	for _, target := range result.Results {
		affixName := affixNames[target.AffixID]
		details = append(details, fmt.Sprintf("• **%s**: Lv%d → Lv%d (%.2f%%)",
			affixName, target.CurrentLevel, target.TargetLevel, target.SuccessRate*100))
	}

	if len(details) > 0 {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "📋 词条详情",
			Value:  strings.Join(details, "\n"),
			Inline: false,
		})
	}

	return embed
}
