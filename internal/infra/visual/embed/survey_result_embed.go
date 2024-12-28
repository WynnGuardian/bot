package embed

import (
	"fmt"
	"time"
	"victo/wynnguardian-bot/internal/domain/config"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

func GetSurveyResultMessage(result *entity.SurveyResult) *discordgo.MessageSend {
	len := util.HighestLength(util.KeySlice(result.Results))
	mask := "``%s | %s | %s``\n"
	idTitle := util.PadText("ID", len)
	sumTitle := util.PadText("Sum", 8)
	averageTitle := util.PadText("New weight %", 12)
	table := ""

	table += fmt.Sprintf(mask, idTitle, sumTitle, averageTitle)

	for id, val := range result.Results {
		padId := util.PadText(id, len)
		padVal := util.PadText(fmt.Sprintf("%.2f", val), 8)
		padAverage := util.PadText(fmt.Sprintf("%.2f %%", result.Results[id]*100), 12)
		table += fmt.Sprintf(mask, padId, padVal, padAverage)
	}

	return &discordgo.MessageSend{
		Content: fmt.Sprintf("<@&%s>", config.MainConfig.Discord.Roles.SurveyResults),
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       fmt.Sprintf("New criteria for item: %s", result.Survey.ItemName),
				Description: fmt.Sprintf("Approved at <t:%d>", time.Now().Unix()),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Total votes",
						Value:  fmt.Sprintf("``%d``", result.TotalVotes),
						Inline: true,
					},
					{
						Name:   "Results",
						Value:  table,
						Inline: false,
					},
				},
			},
		},
	}
}
