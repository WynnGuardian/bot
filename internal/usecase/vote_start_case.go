package usecase

import (
	"fmt"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/util"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

type StartVotingUsecase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewStartVotingUsecase(s *discordgo.Session, i *discordgo.InteractionCreate) *StartVotingUsecase {
	return &StartVotingUsecase{
		session:     s,
		interaction: i,
	}
}

func (u *StartVotingUsecase) Execute(input api.StartVotingUsecase) {
	api.MustCallAndUnwrap(api.GetSurveyAPI().StartVoting, input, func(t *entity.SurveyVote) {
		url := fmt.Sprintf("http://localhost:5173/#/vote/%s?token=%s", t.Survey.ID, t.Token)
		err := util.MessageUser(u.session, u.interaction.Member.User.ID, embed.GetVoteCreateEmbed(url, input.Item))
		if err != nil {
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}
		response.WithMessage("Check DM to proceed.", u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[entity.SurveyVote](u.session, u.interaction))
}
