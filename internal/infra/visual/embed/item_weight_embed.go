package embed

import (
	"fmt"
	"strings"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
)

func GetItemWeightEmbed(resp *api.WeightResponse) *discordgo.MessageEmbed {

	textLenght := util.HighestLength(util.KeySlice(resp.StaticItem.Stats))
	table := ""

	idTitle := util.PadText("ID", textLenght)
	valueTitle := util.PadText("Value", 5)
	percentTitle := util.PadText("Value %", 7)
	criteriaTitle := util.PadText("Criteria", 8)
	weightedTitle := util.PadText("Weighted", 8)
	mask := "``%s | %s | %s | %s | %s``\n"

	table += fmt.Sprintf(mask, idTitle, valueTitle, percentTitle, criteriaTitle, weightedTitle)
	table += "``" + strings.Repeat("-", len(table)) + "``"

	for id, criteria := range resp.Criteria.Modifiers {
		item := resp.StaticItem

		norm := normalize(
			item.Stats[id],
			item.WynnItem.Stats[id].Minimum,
			item.WynnItem.Stats[id].Maximum)

		weighted := criteria * float64(norm)

		id := util.PadText(id, textLenght)
		val := util.PadText(fmt.Sprintf("%d", item.Stats[id]), 5)
		normalized := util.PadText(fmt.Sprintf("%.2f%%", norm*100), 7)
		criteria := util.PadText(fmt.Sprintf("%.3f", criteria), 8)
		weight := util.PadText(fmt.Sprintf("%.1f%%", weighted*100), 8)

		table += fmt.Sprintf(mask, id, val, normalized, criteria, weight)
	}

	return &discordgo.MessageEmbed{
		Title:       resp.StaticItem.WynnItem.Name,
		Description: fmt.Sprintf("Total Weight: %.3f%%", resp.Weight*100),
		Color:       int(getColor(int(resp.Weight))),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "",
				Value:  table,
				Inline: false,
			},
			{
				Name:   "Total Weight",
				Value:  fmt.Sprintf("```cpp\n%.2f%%```", resp.Weight*100),
				Inline: false,
			},
		},
	}
}

func normalize(val, min, max int) float64 {
	return float64((float64(val) - float64(min)) / (float64(max) - float64(min)))
}

func getColor(value int) int {
	if value < 0 || value > 1 {
		return 0x000000
	}
	startColor := [3]int{0, 255, 242}
	endColor := [3]int{255, 64, 64}

	r := (1-value)*startColor[0] + value*endColor[0]
	g := (1-value)*startColor[1] + value*endColor[1]
	b := (1-value)*startColor[2] + value*endColor[2]
	return (r << 16) | (g << 8) | b
}
