package embed

import (
	"fmt"
	"strings"
	"victo/wynnguardian-bot/internal/domain/config"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

func GetRankListMessage(resp []entity.AuthenticatedItem, itemName, msg, channel string, offset, limit int, session *discordgo.Session) *discordgo.MessageSend {

	mask := "`` %s | %s | %s | %s ``\n"
	table := ""

	posTitle := util.PadText("Pos", 3)
	weightTitle := util.PadText("Weight", 6)
	tcodeTitle := util.PadText("Tracking Code", 16)
	ownerTitle := util.PadText("Owner", 18)

	table += fmt.Sprintf(mask, posTitle, weightTitle, tcodeTitle, ownerTitle)
	table += fmt.Sprintf("``%s``\n", strings.Repeat("-", len(table)))

	for ind, s := range resp {
		pos := util.PadText(fmt.Sprintf("%d", ind+offset*limit+1), 3)
		weight := util.PadText(fmt.Sprintf("%.2f%%", s.Weight), 6)
		trackingCode := util.PadText(s.TrackingCode, 16)

		owner := "HIDDEN"
		if s.PublicOwner {
			user, err := session.GuildMember(config.MainConfig.Discord.MainGuild, s.OwnerDC)
			if err == nil {
				owner = user.User.Username
			}
		}
		owner = util.PadText(owner, 37)

		table += fmt.Sprintf(mask, pos, weight, trackingCode, owner)
	}

	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: fmt.Sprintf("Top 100 %s", itemName),
				Color: 0xb700ff,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Rank",
						Value:  table,
						Inline: false,
					},
				},
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Emoji:    &discordgo.ComponentEmoji{Name: "◀"},
						Style:    discordgo.PrimaryButton,
						CustomID: fmt.Sprintf("rankpreviouspage_%s_%s_%s", msg, channel, itemName),
					},
					discordgo.Button{
						Emoji:    &discordgo.ComponentEmoji{Name: "▶"},
						Style:    discordgo.PrimaryButton,
						CustomID: fmt.Sprintf("ranknextpage_%s_%s_%s", msg, channel, itemName),
					},
				},
			},
		},
	}
}
