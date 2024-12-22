package util

import (
	"github.com/bwmarrin/discordgo"
)

func MessageUser(session *discordgo.Session, userId string, embed *discordgo.MessageSend) error {
	channel, err := session.UserChannelCreate(userId)
	if err != nil {
		return err
	}
	_, err = session.ChannelMessageSendComplex(channel.ID, embed)
	if err != nil {
		return err
	}
	return nil
}
