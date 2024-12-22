package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
)

type ItemAuthUsecase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewItemAuthUsecase(s *discordgo.Session, i *discordgo.InteractionCreate) *ItemAuthUsecase {
	return &ItemAuthUsecase{
		session:     s,
		interaction: i,
	}
}

func (u *ItemAuthUsecase) Execute(input api.AuthenticateItemInput) {
	api.MustCallAndUnwrap(api.GetItemAPI().AuthenticateItem, input, func(t *api.AuthenticateItemResponse) {
		err := util.MessageUser(u.session, input.DCOwnerUID, embed.GetTrackingCodeDMMessage(*t))
		if err != nil {
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}
		response.WithEmbed(embed.GetTrackingCodeSuccessEmbed(t.TrackingCode, err), u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[api.AuthenticateItemResponse](u.session, u.interaction))
}
