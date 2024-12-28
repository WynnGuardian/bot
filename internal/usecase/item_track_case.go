package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
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

func (u *ItemTrackCase) Execute(input api.WeightItemInput) {
	api.MustCallAndUnwrap(api.GetItemAPI().WeightItem, input, func(t *api.WeightResponse) {
		response.WithEmbed(embed.GetItemWeightEmbed(t), false, u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[api.WeightResponse](u.session, u.interaction))
}
