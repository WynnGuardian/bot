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

type VoteDenyCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewVoteDenyCase(s *discordgo.Session, i *discordgo.InteractionCreate) *VoteDenyCase {
	return &VoteDenyCase{
		session:     s,
		interaction: i,
	}
}

func (u *VoteDenyCase) Execute(input api.DenyVoteInput) {
	api.MustCallAndUnwrap(api.GetSurveyAPI().DenyVote, input, func(vote *entity.SurveyVote) {
		err := u.session.MessageReactionAdd(vote.Survey.ChannelID, vote.MessageID, "❌")
		if err != nil {
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}
		msg := embed.GetVoteMessage(*vote)
		_, err = u.session.ChannelMessageEditComplex(&discordgo.MessageEdit{
			Channel:    vote.Survey.ChannelID,
			ID:         vote.MessageID,
			Content:    &msg.Content,
			Components: &msg.Components,
			Embeds:     &msg.Embeds,
		})
		if err != nil {
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}
		err = util.MessageUser(u.session, vote.DiscordUserID, embed.GetVoteDeniedMessage(vote.Survey.ID, vote.Survey.ItemName))
		if err != nil {
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}
		response.WithMessage("Vote denied.", true, u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[entity.SurveyVote](u.session, u.interaction))
}
