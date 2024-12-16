package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
)

type ItemWeighUsecase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewItemWeighUsecase(s *discordgo.Session, i *discordgo.InteractionCreate) *ItemWeighUsecase {
	return &ItemWeighUsecase{
		session:     s,
		interaction: i,
	}
}

func (u *ItemWeighUsecase) Execute(input api.WeightItemInput) {
	api.MustCallAndUnwrap(api.GetItemAPI().WeightItem, input, func(t *api.WeightResponse) {
		response.WithEmbed(embed.GetItemWeightEmbed(t), u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[api.WeightResponse](u.session, u.interaction))
}
