package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/SpenserCai/OnceHumanTools/bot/platforms/discord"
)

// CreateHelpCommand 创建帮助命令
func CreateHelpCommand() *discord.SlashCommand {
	return &discord.SlashCommand{
		Command: &discordgo.ApplicationCommand{
			Name:        "help",
			Description: "显示OnceHuman工具集帮助信息",
		},
		Handler: handleHelpCommand,
	}
}

// handleHelpCommand 处理帮助命令
func handleHelpCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resp := discord.CreateResponse(s, i)

	embed := &discordgo.MessageEmbed{
		Title:       "🛠 OnceHuman工具集帮助",
		Description: "欢迎使用OnceHuman游戏工具集！以下是可用的命令：",
		Color:       0x00D4FF,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "📊 /affix - 词条概率计算",
				Value: `计算模组出现特定词条的概率
**参数：**
• \`slots\` - 词条数量 (1-10)
• \`targets\` - 目标词条ID，逗号分隔
• \`show_combinations\` - 显示详细组合

**示例：** \`/affix slots:4 targets:1,4,5\``,
				Inline: false,
			},
			{
				Name: "🎯 /strengthen single - 单词条强化",
				Value: `计算单个词条强化到目标等级的概率
**参数：**
• \`affix_id\` - 词条ID (1-10)
• \`current_level\` - 当前等级 (0-5)
• \`target_level\` - 目标等级 (1-5)
• \`slot_count\` - 词条数量
• \`tries\` - 强化次数

**示例：** \`/strengthen single affix_id:1 current_level:0 target_level:3 slot_count:4 tries:50\``,
				Inline: false,
			},
			{
				Name: "🎯 /strengthen multi - 多词条强化",
				Value: `计算多个词条同时强化的概率
**参数：**
• \`targets\` - 格式: ID:当前:目标
• \`slot_count\` - 词条数量
• \`tries\` - 强化次数

**示例：** \`/strengthen multi targets:1:0:3,4:1:5 slot_count:4 tries:100\``,
				Inline: false,
			},
			{
				Name: "📖 词条ID对照表",
				Value: `1=异常伤害 | 2=弹匣容量 | 3=换弹速度
4=对普通敌人伤害 | 5=对精英敌人伤害
6=对上位者伤害 | 7=最大生命值
8=头部受伤减免 | 9=枪械伤害减免
10=异常伤害减免`,
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "OnceHuman工具集 - 更多功能开发中...",
			IconURL: "https://i.imgur.com/AfFp7pu.png",
		},
		Timestamp: discordgo.NowTimestamp(),
	}

	resp.SendEmbed(embed)
}