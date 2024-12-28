package embed

import (
	"fmt"
	"strings"
	"time"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
	"github.com/wynnguardian/common/enums"
)

func GetVoteMessage(vote entity.SurveyVote) *discordgo.MessageSend {

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

	status := "```fix\nWAITING APPOVAL\n```"
	if vote.Status == enums.VOTE_CONTABILIZED {
		status = "```diff\n+ CONTABILIZED\n```"
	}
	if vote.Status == enums.VOTE_DENIED {
		status = "```diff\n- DENIED\n```"
	}

	msg := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{},
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Vote received",
				Description: "Confirm or deny the vote using the buttons bellow",
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
					{
						Name:   "Status",
						Value:  status,
						Inline: false,
					},
				},
			},
		},
	}

	if vote.Status == enums.VOTE_NOT_CONFIRMED {
		msg.Components = []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Emoji:    &discordgo.ComponentEmoji{Name: "✅"},
						Style:    discordgo.PrimaryButton,
						CustomID: fmt.Sprintf("confirmvote_%s", vote.Token),
					},
					discordgo.Button{
						Emoji:    &discordgo.ComponentEmoji{Name: "❌"},
						Style:    discordgo.PrimaryButton,
						CustomID: fmt.Sprintf("denyvote_%s", vote.Token),
					},
				},
			},
		}
	}

	return msg
}
