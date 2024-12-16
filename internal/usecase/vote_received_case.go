package usecase

import (
	"time"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/infra/util"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
	"github.com/wynnguardian/common/response"
)

type VoteReceivedCase struct {
	s *discordgo.Session
}

func NewVoteReceivedCase(s *discordgo.Session) *VoteReceivedCase {
	return &VoteReceivedCase{
		s: s,
	}
}

func (u *VoteReceivedCase) Execute(input entity.SurveyVote) response.WGResponse {
	embed := embed.GetVoteEmbed(input)
	msg, err := util.SendVoteConfirmMessage(u.s, input.Survey.ChannelID, input.Survey.ID, input.DiscordUserID, embed)
	if err != nil {
		return response.ErrBadRequest
	}

	go func() {
		time.Sleep(500)
		_, err = api.GetSurveyAPI().DefineVoteMessage(api.DefineVoteMessageInput{
			SurveyID:  input.Survey.ID,
			UserID:    input.DiscordUserID,
			ChannelID: input.Survey.ChannelID,
			MessageID: msg,
		})

		if err != nil {
			util.LogError(err, input.Survey.ChannelID, "survey fill", u.s)
		}
	}()

	return response.Ok
}