package usecase

import (
	"victo/wynnguardian-bot/internal/domain/config"
	"victo/wynnguardian-bot/internal/infra/util"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
	"github.com/wynnguardian/common/response"
)

type SurveyEndedCase struct {
	s *discordgo.Session
}

func NewSurveyEndedCase(s *discordgo.Session) *SurveyEndedCase {
	return &SurveyEndedCase{
		s: s,
	}
}

func (u *SurveyEndedCase) Execute(input entity.SurveyResult) response.WGResponse {
	embed := embed.GetSurveyApprovalEmbed(&input)
	err := util.SendEmbedMessage(u.s, config.MainConfig.Discord.Channels.SurveyWaitingApproval, embed)

	if err != nil {
		return response.ErrInternalServerErr(err)
	}
	return response.Ok
}
