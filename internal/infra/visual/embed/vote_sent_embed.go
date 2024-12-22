package embed

import (
	"fmt"
	"strings"
	"time"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

func GetVoteEmbed(vote entity.SurveyVote) *discordgo.MessageEmbed {

	textLenght := util.HighestLength(util.KeySlice(vote.Votes))
	table := ""
	mask := "``%s | %s``\n"
	idTitle := util.PadText("ID", textLenght)
	weightTitle := util.PadText("Weight %", 10)

	table += fmt.Sprintf(mask, idTitle, weightTitle)
	table += ("``" + strings.Repeat("-", len(table)) + "``\n")
	for id, value := range vote.Votes {
		id := util.PadText(id, textLenght)
		val := util.PadText(fmt.Sprintf("%.2f", value*100), 10)
		table += fmt.Sprintf(mask, id, val)
	}

	return &discordgo.MessageEmbed{
		Title:       "Vote received",
		Description: fmt.Sprintf("Confirm with the button below or type ``/survey confirm %s``", vote.Token),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "User ID",
				Value:  fmt.Sprintf("<@%s>", vote.DiscordUserID),
				Inline: false,
			},
			{
				Name:   "Sent at",
				Value:  fmt.Sprintf("<t:%d>", time.Now().Unix()),
				Inline: false,
			},
			{
				Name:   "Votes",
				Value:  table,
				Inline: false,
			},
		},
	}
}
