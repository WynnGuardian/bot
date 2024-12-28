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

type SurveyDiscardUsecase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewSurveyDiscardUsecase(s *discordgo.Session, i *discordgo.InteractionCreate) *SurveyDiscardUsecase {
	return &SurveyDiscardUsecase{
		session:     s,
		interaction: i,
	}
}

func (u *SurveyDiscardUsecase) Execute(input api.SurveyDiscardCaseInput) {
	s, i := u.session, u.interaction
	api.MustCallAndUnwrap(api.GetSurveyAPI().DiscardSurvey, input, func(t *entity.Survey) {
		response.WithMessage("Survey closed successfully!", true, s, i)

		msg := embed.GetSurveyAnnounceMessage(t)
		edit := discordgo.MessageEdit{
			ID:         t.AnnouncementMessageID,
			Channel:    config.MainConfig.Discord.Channels.SurveyAnnouncements,
			Components: &msg.Components,
			Embeds:     &msg.Embeds,
		}

		u.session.ChannelMessageEditComplex(&edit)

	}, cerrors.CatchAndLogInternal(s, i), cerrors.CatchAndLogAPIError[entity.Survey](s, i))
}
