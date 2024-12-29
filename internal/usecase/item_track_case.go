package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

type ItemTrackCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewItemTrackCase(s *discordgo.Session, i *discordgo.InteractionCreate) *ItemTrackCase {
	return &ItemTrackCase{
		session:     s,
		interaction: i,
	}
}

func (u *ItemTrackCase) Execute(input api.FindItemInput) {
	api.MustCallAndUnwrap(api.GetItemAPI().FindItem, input, func(t *entity.AuthenticatedItem) {
		api.MustCallAndUnwrap(api.GetItemAPI().FindCriteria, api.FindCriteriaInput{ItemName: t.Item}, func(t2 *entity.ItemCriteria) {
			response.WithEmbed(embed.GetItemTrackEmbed(t, t2), false, u.session, u.interaction)
		}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[entity.ItemCriteria](u.session, u.interaction))
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[entity.AuthenticatedItem](u.session, u.interaction))
}
