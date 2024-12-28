package util

import (
	"github.com/bwmarrin/discordgo"
)

func SendEmbedMessage(s *discordgo.Session, channel string, embed *discordgo.MessageEmbed) error {
	_, err := s.ChannelMessageSendEmbed(channel, embed)
	return err
}

func SendMessage(s *discordgo.Session, channel string, msg *discordgo.MessageSend) error {
	_, err := s.ChannelMessageSendComplex(channel, msg)
	return err
}
