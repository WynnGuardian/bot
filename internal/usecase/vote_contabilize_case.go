package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/util"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

type ConfirmVoteUsecase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewConfirmVoteUsecase(s *discordgo.Session, i *discordgo.InteractionCreate) *ConfirmVoteUsecase {
	return &ConfirmVoteUsecase{
		session:     s,
		interaction: i,
	}
}

func (u *ConfirmVoteUsecase) Execute(input api.ConfirmVoteUsecaseInput) {
	api.MustCallAndUnwrap(api.GetSurveyAPI().ConfirmVote, input, func(t *entity.SurveyVote) {
		err := u.session.MessageReactionAdd(t.Survey.ChannelID, t.MessageID, "✅")
		if err != nil {
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}
		msg := embed.GetVoteMessage(*t)
		_, err = u.session.ChannelMessageEditComplex(&discordgo.MessageEdit{
			Channel:    t.Survey.ChannelID,
			ID:         t.MessageID,
			Content:    &msg.Content,
			Components: &msg.Components,
			Embeds:     &msg.Embeds,
		})
		if err != nil {
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}
		err = util.MessageUser(u.session, t.DiscordUserID, embed.GetVoteConfirmedMessage(t.Survey.ID, t.Survey.ItemName))
		if err != nil {
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}

		response.WithMessage("Vote contabilized!", true, u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[entity.SurveyVote](u.session, u.interaction))
}
