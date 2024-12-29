package embed

import (
	"fmt"
	"strings"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

func GetItemTrackEmbed(resp *entity.AuthenticatedItem, crit *entity.ItemCriteria) *discordgo.MessageEmbed {
	textLenght := util.HighestLength(util.KeySlice(resp.WynnItem.Stats))
	table := ""

	idTitle := util.PadText("ID", textLenght)
	valueTitle := util.PadText("Value", 8)
	percentTitle := util.PadText("Value %", 7)
	criteriaTitle := util.PadText("Criteria", 8)
	weightedTitle := util.PadText("Weighted", 8)
	mask := "``%s | %s | %s | %s | %s``\n"

	table += fmt.Sprintf(mask, idTitle, valueTitle, percentTitle, criteriaTitle, weightedTitle)
	table += "``" + strings.Repeat("-", len(table)) + "``\n"

	for id, criteria := range crit.Modifiers {
		item := resp.WynnItem

		norm := normalize(
			resp.Stats[id],
			item.Stats[id].Minimum,
			item.Stats[id].Maximum)

		weighted := criteria * float64(norm)

		sprinted := fmt.Sprintf("%d", int32(resp.Stats[id]))

		id := util.PadText(id, textLenght)
		val := util.PadText(sprinted, 8)
		normalized := util.PadText(fmt.Sprintf("%.2f%%", norm*100), 7)
		criteria := util.PadText(fmt.Sprintf("%.3f", criteria), 8)
		weight := util.PadText(fmt.Sprintf("%.1f%%", weighted*100), 8)

		table += fmt.Sprintf(mask, id, val, normalized, criteria, weight)
	}

	return &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Tracking %s (%s)", resp.Item, resp.TrackingCode),
		Color: int(getColor(int(resp.Weight))),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "",
				Value:  table,
				Inline: false,
			},
			{
				Name:   "Total Weight",
				Value:  fmt.Sprintf("```cpp\n%.2f%%```", resp.Weight),
				Inline: false,
			},
			{
				Name:   "Rank Position",
				Value:  fmt.Sprintf("```cpp\n%d```", resp.Position),
				Inline: false,
			},
		},
	}
}
