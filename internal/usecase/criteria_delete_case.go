package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"

	"github.com/bwmarrin/discordgo"
)

type DeleteCriteriaCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewDeleteCriteriaCase(s *discordgo.Session, i *discordgo.InteractionCreate) *DeleteCriteriaCase {
	return &DeleteCriteriaCase{
		session:     s,
		interaction: i,
	}
}

func (u *DeleteCriteriaCase) Execute(input api.DeleteCriteriaInput) {
	api.MustCallAndUnwrap(api.GetItemAPI().DeleteCriteria, input, func(t *any) {
		response.WithMessage("Criteria deleted successfully!", true, u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[any](u.session, u.interaction))
}
