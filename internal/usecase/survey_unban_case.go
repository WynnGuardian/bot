package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"

	"github.com/bwmarrin/discordgo"
)

type SurveyUnbanCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewSurveyUnbanCase(s *discordgo.Session, i *discordgo.InteractionCreate) *SurveyUnbanCase {
	return &SurveyUnbanCase{
		session:     s,
		interaction: i,
	}
}

func (u *SurveyUnbanCase) Execute(input api.SurveyUnbanInput) {
	api.MustCallAndUnwrap(api.GetSurveyAPI().SurveyUnban, input, func(t *any) {
		response.WithMessage("User banned from surveys successfully!", true, u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[any](u.session, u.interaction))
}
