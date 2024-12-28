package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/config"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
)

type SurveyApproveCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewSurveyApproveCase(s *discordgo.Session, i *discordgo.InteractionCreate) *SurveyApproveCase {
	return &SurveyApproveCase{
		session:     s,
		interaction: i,
	}
}

func (u *SurveyApproveCase) Execute(input api.SurveyApproveCaseInput) {
	api.MustCallAndUnwrap(api.GetSurveyAPI().ApproveSurvey, input, func(t *api.SurveyApproveResponse) {
		response.WithMessage("Survey approved successfully!", true, u.session, u.interaction)

		msg := embed.GetSurveyAnnounceMessage(&t.Survey)
		edit := &discordgo.MessageEdit{
			ID:         t.Survey.AnnouncementMessageID,
			Channel:    config.MainConfig.Discord.Channels.SurveyAnnouncements,
			Components: &msg.Components,
			Embeds:     &msg.Embeds,
		}

		u.session.ChannelMessageEditComplex(edit)

		u.session.ChannelMessageSendComplex(config.MainConfig.Discord.Channels.SurveyPublicResults, embed.GetSurveyResultMessage(&t.Result))

	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[api.SurveyApproveResponse](u.session, u.interaction))
}
