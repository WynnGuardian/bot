package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/infra/cerrors"

	"github.com/bwmarrin/discordgo"
)

type DefineVoteMessageUsecase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewDefineVoteMessageUsecase(s *discordgo.Session, i *discordgo.InteractionCreate) *DefineVoteMessageUsecase {
	return &DefineVoteMessageUsecase{
		session:     s,
		interaction: i,
	}
}

func (u *DefineVoteMessageUsecase) Execute(input api.DefineVoteMessageInput) {
	api.MustCallAndUnwrap(api.GetSurveyAPI().DefineVoteMessage, input, func(_ *any) {},
		cerrors.CatchAndLogInternal(u.session, u.interaction),
		cerrors.CatchAndLogAPIError[any](u.session, u.interaction))
}
