package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"

	"github.com/bwmarrin/discordgo"
)

type RankUpdateCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewRankUpdateCase(s *discordgo.Session, i *discordgo.InteractionCreate) *RankUpdateCase {
	return &RankUpdateCase{
		session:     s,
		interaction: i,
	}
}

func (u *RankUpdateCase) Execute(input api.RankUpdateCaseInput) {
	api.MustCallAndUnwrap(api.GetItemAPI().RankUpdate, input, func(t *any) {
		response.WithMessage("Rank updated successfully!", u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[any](u.session, u.interaction))
}
