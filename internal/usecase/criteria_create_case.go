package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"

	"github.com/bwmarrin/discordgo"
)

type CreateCriteriaCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewCreateCriteriaCase(s *discordgo.Session, i *discordgo.InteractionCreate) *CreateCriteriaCase {
	return &CreateCriteriaCase{
		session:     s,
		interaction: i,
	}
}

func (u *CreateCriteriaCase) Execute(input api.CreateCriteriaInput) {
	api.MustCallAndUnwrap(api.GetItemAPI().CreateCriteria, input, func(t *any) {
		response.WithMessage("Criteria created successfully!", u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[any](u.session, u.interaction))
}
