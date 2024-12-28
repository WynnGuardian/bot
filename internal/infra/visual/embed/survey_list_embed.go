package embed

import (
	"fmt"
	"strings"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
	"github.com/victorbetoni/go-streams/streams"
	"github.com/wynnguardian/common/entity"
)

func GetSurveyListEmbed(resp []entity.Survey) *discordgo.MessageEmbed {

	itemNames := *streams.Map(streams.StreamOf(resp...), func(r entity.Survey) string {
		return r.ItemName
	}).ToSlice()
	textLenght := util.HighestLength(itemNames)
	mask := "``%s | %s | ``%s\n"
	table := ""

	idTitle := util.PadText("ID", 10)
	itemTitle := util.PadText("Item", textLenght)
	endsTitle := util.PadText("Deadline", 16)

	table += fmt.Sprintf("``%s | %s | %s``\n", idTitle, itemTitle, endsTitle)
	table += fmt.Sprintf("``%s``\n", strings.Repeat("-", len(table)))
	for _, s := range resp {
		id := util.PadText(s.ID, 10)
		item := util.PadText(s.ItemName, textLenght)
		endsIn := util.PadText(fmt.Sprintf("<t:%d:R>", s.Deadline.Unix()), 16)
		table += fmt.Sprintf(mask, id, item, endsIn)
	}

	return &discordgo.MessageEmbed{
		Title: "Currently open surveys",
		Color: 0xb700ff,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Surveys",
				Value:  table,
				Inline: false,
			},
		},
	}
}
