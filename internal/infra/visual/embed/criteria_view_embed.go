package embed

import (
	"fmt"
	"strings"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

func GetItemCriteriaEmbed(resp *entity.ItemCriteria) *discordgo.MessageEmbed {

	textLenght := util.HighestLength(util.KeySlice(resp.Modifiers))
	table := ""

	idTitle := util.PadText("Stat", textLenght)
	valueTitle := util.PadText("Weight %", 5)
	mask := "``%s | %s``\n"

	table += fmt.Sprintf(mask, idTitle, valueTitle)
	table += "``" + strings.Repeat("-", len(table)) + "``\n"

	for id, criteria := range resp.Modifiers {
		id := util.PadText(id, textLenght)
		weight := util.PadText(fmt.Sprintf("%.3f%%", criteria*100), 8)

		table += fmt.Sprintf(mask, id, weight)
	}

	return &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s weight criteria", resp.Item),
		Color: 0xb700ff,
		Fields: []*discordgo.MessageEmbedField{
			{
				Value:  table,
				Inline: false,
			},
		},
	}
}
