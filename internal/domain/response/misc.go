package response

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func WithMessage(msg string, ephemeral bool, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	d := &discordgo.InteractionResponseData{
		Content: fmt.Sprintf("``%s``", msg),
	}
	if ephemeral {
		d.Flags = discordgo.MessageFlagsEphemeral
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: d,
	})
}

func WithEmbed(embed *discordgo.MessageEmbed, ephemeral bool, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	d := &discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{embed},
	}
	if ephemeral {
		d.Flags = discordgo.MessageFlagsEphemeral
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: d,
	})
}
