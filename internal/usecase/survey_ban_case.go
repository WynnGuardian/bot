package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"

	"github.com/bwmarrin/discordgo"
)

type SurveyBanCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewSurveyBanCase(s *discordgo.Session, i *discordgo.InteractionCreate) *SurveyBanCase {
	return &SurveyBanCase{
		session:     s,
		interaction: i,
	}
}

func (u *SurveyBanCase) Execute(input api.SurveyBanInput) {
	api.MustCallAndUnwrap(api.GetSurveyAPI().SurveyBan, input, func(t *any) {
		response.WithMessage("User banned from surveys successfully!", true, u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[any](u.session, u.interaction))
}
