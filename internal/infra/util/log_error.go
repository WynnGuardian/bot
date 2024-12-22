package util

import (
	"runtime"
	"time"
	"victo/wynnguardian-bot/internal/domain/config"

	"github.com/bwmarrin/discordgo"
)

func LogError(err error, fromChan, cmd string, session *discordgo.Session) {
	channel := config.MainConfig.Discord.Channels.ErrorLog
	b := make([]byte, 2048)
	n := runtime.Stack(b, false)
	s := string(b[:n])
	embed := &discordgo.MessageEmbed{
		Title: "Error",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Time",
				Value:  time.Now().Format(time.RFC3339),
				Inline: true,
			},
			{
				Name:   "Message",
				Value:  err.Error(),
				Inline: false,
			},
			{
				Name:   "Command",
				Value:  cmd,
				Inline: true,
			},
			{
				Name:   "Channel ID:",
				Value:  fromChan,
				Inline: true,
			},
			{
				Name:   "Stacktrace",
				Value:  s,
				Inline: false,
			},
		},
	}
	SendEmbedMessage(session, channel, embed)
}
