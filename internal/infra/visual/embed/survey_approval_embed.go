package embed

import (
	"fmt"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

func GetSurveyApprovalEmbed(result *entity.SurveyResult) *discordgo.MessageEmbed {

	len := util.HighestLength(util.KeySlice(result.Results))
	mask := "``%s | %s | %s``\n"
	idTitle := util.PadText("ID", len)
	sumTitle := util.PadText("Sum", 8)
	averageTitle := util.PadText("Average", 10)
	table := ""

	table += fmt.Sprintf(mask, idTitle, sumTitle, averageTitle)

	for id, val := range result.Results {
		padId := util.PadText(id, len)
		padVal := util.PadText(fmt.Sprintf("%.2f", val), 8)
		padAverage := util.PadText(fmt.Sprintf("%.2f", result.Results[id]), 10)
		table += fmt.Sprintf(mask, padId, padVal, padAverage)
	}

	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Item criteria survey ended for item: %s", result.ItemName),
		Description: fmt.Sprintf("Type ``/survey approve %s`` to approve results.", result.SurveyID),
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
	}
}
