package util

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SendEmbedMessage(s *discordgo.Session, channel string, embed *discordgo.MessageEmbed) error {
	_, err := s.ChannelMessageSendEmbed(channel, embed)
	return err
}

func SendVoteConfirmMessage(s *discordgo.Session, channel, user, survey string, embed *discordgo.MessageEmbed) (string, error) {
	msg, err := s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{embed},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						CustomID: fmt.Sprintf("confirmvote_%s_%s", survey, user),
						Label:    "Contabilize",
					},
				},
			},
		},
	})
	return msg.ID, err
}
