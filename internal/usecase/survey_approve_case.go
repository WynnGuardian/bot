package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/config"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
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
	api.MustCallAndUnwrap(api.GetSurveyAPI().ApproveSurvey, input, func(t *entity.Survey) {
		response.WithMessage("Survey approved successfully!", u.session, u.interaction)

		msg := embed.GetSurveyAnnounceMessage(t)
		edit := &discordgo.MessageEdit{
			ID:         t.AnnouncementMessageID,
			Channel:    config.MainConfig.Discord.Channels.SurveyPublicResults,
			Components: &msg.Components,
			Embeds:     &msg.Embeds,
		}

		u.session.ChannelMessageEditComplex(edit)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[entity.Survey](u.session, u.interaction))
}
