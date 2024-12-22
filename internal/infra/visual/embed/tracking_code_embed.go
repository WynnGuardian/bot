package embed

import (
	"fmt"
	"strings"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
)

func GetTrackingCodeSuccessEmbed(trackingCode string, userMsgErr error) *discordgo.MessageEmbed {

	notifyMsg := "Yes"
	if userMsgErr != nil {
		notifyMsg = fmt.Sprintf("No. Error: %s", userMsgErr.Error())
	}

	return &discordgo.MessageEmbed{
		Title:       "Item authenticated successfully!",
		Description: "If a valid Discord User ID was provided, the tracking code was automatically sent to the item owner.",
		Color:       0xb700ff,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Tracking code",
				Value:  fmt.Sprintf("``%s``", trackingCode),
				Inline: false,
			},
			{
				Name:   "Was user notified?",
				Value:  notifyMsg,
				Inline: false,
			},
		},
	}
}

func GetTrackingCodeDMMessage(resp api.AuthenticateItemResponse) *discordgo.MessageSend {

	length := 0
	for k := range resp.Item.Stats {
		if len(k) > length {
			length = len(k)
		}
	}

	table := make([]string, 0)
	table = append(table, fmt.Sprintf("``%s | %s | Overall  ``", util.PadText("ID", length), util.PadText("Value", 10)))
	table = append(table, fmt.Sprintf("``%s``", strings.Repeat("-", len(table[0]))))
	for id, val := range resp.Item.Stats {
		stat := resp.WynnItem.Stats[id]
		max, min := stat.Minimum, stat.Maximum
		pct := (float64(max-val) / float64(max-min)) * 100
		table = append(table, fmt.Sprintf("``%s | %s | %s%% ``", util.PadText(id, length), util.PadText(fmt.Sprintf(
			"%d", val), 10), util.PadText(fmt.Sprintf("%.1f", pct), 7)))
	}

	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Your item has been authenticated!",
				Description: "You can now track it's rank with /item track <tracking-code>",
				Color:       0xb700ff,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Tracking code",
						Value:  fmt.Sprintf("``%s``", resp.TrackingCode),
						Inline: false,
					},
					{
						Name:   resp.Item.Item,
						Value:  strings.Join(table[:], "\n"),
						Inline: false,
					},
				},
			},
		},
	}
}
