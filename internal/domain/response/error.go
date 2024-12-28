package response

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ErrorResponse(err error, internal bool, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	msg := "Something went wrong. This error was logged and is already being handled."
	if !internal {
		msg = err.Error()
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("``Error: %s``", msg),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func UnauthorizedResponse(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "``You do not have permission to perform this action``",
		},
	})
}
