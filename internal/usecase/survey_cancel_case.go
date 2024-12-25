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

type SurveyCancelUsecase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewSurveyCancelUsecase(s *discordgo.Session, i *discordgo.InteractionCreate) *SurveyCancelUsecase {
	return &SurveyCancelUsecase{
		session:     s,
		interaction: i,
	}
}

func (u *SurveyCancelUsecase) Execute(input api.SurveyCancelUsecaseInput) {

	api.MustCallAndUnwrap(api.GetSurveyAPI().CancelSurvey, input, func(t *entity.Survey) {
		_, err := u.session.ChannelDelete(t.ChannelID)
		if err != nil {
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}
		response.WithMessage("Survey cancel successfully!", u.session, u.interaction)

		msg := embed.GetSurveyAnnounceMessage(t)
		edit := discordgo.MessageEdit{
			ID:         t.AnnouncementMessageID,
			Channel:    config.MainConfig.Discord.Channels.SurveyPublicResults,
			Components: &msg.Components,
			Embeds:     &msg.Embeds,
		}

		u.session.ChannelMessageEditComplex(&edit)

	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[entity.Survey](u.session, u.interaction))
}
