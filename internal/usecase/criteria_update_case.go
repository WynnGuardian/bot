package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"

	"github.com/bwmarrin/discordgo"
)

type UpdateCriteriaCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewUpdateCriteriaCase(s *discordgo.Session, i *discordgo.InteractionCreate) *UpdateCriteriaCase {
	return &UpdateCriteriaCase{
		session:     s,
		interaction: i,
	}
}

func (u *UpdateCriteriaCase) Execute(input api.UpdateCriteriaInput) {
	api.MustCallAndUnwrap(api.GetItemAPI().UpdateCriteria, input, func(t *any) {
		response.WithMessage("Criteria updated successfully!", u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[any](u.session, u.interaction))
}
