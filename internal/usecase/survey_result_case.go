package usecase

import (
	"fmt"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/config"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

type SurveyResultCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewSurveyResultUsecase(s *discordgo.Session, i *discordgo.InteractionCreate) *SurveyResultCase {
	return &SurveyResultCase{
		session:     s,
		interaction: i,
	}
}

func (u *SurveyResultCase) Execute(input api.OpenSurveyInput) {
	api.MustCallAndUnwrap(api.GetSurveyAPI().OpenSurvey, input, func(t *entity.Survey) {

		msg := embed.GetSurveyAnnounceMessage(t)

		announce, err := u.session.ChannelMessageSendComplex(config.MainConfig.Discord.Channels.SurveyAnnouncements, msg)
		if err != nil {
			err = fmt.Errorf("error while creating announcement message: %s", err.Error())
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}

		m, err := u.session.ChannelMessageSend(config.MainConfig.Discord.Channels.VotesWaitingApproval, fmt.Sprintf("Starting new voting thread for survey %s (%s)...", t.ItemName, t.ID))
		if err != nil {
			err = fmt.Errorf("error while creating thread message: %s", err.Error())
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}

		thread, err := u.session.MessageThreadStartComplex(m.ChannelID, m.ID, &discordgo.ThreadStart{
			Name:                fmt.Sprintf("%s (%s)", t.ItemName, t.ID),
			AutoArchiveDuration: 4320,
			Invitable:           false,
			RateLimitPerUser:    0,
		})

		if err != nil {
			err = fmt.Errorf("error while creating vote thread: %s", err.Error())
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}

		defineChanIn := api.DefineSurveyInfoInput{
			Survey:            t.ID,
			ChannelID:         thread.ID,
			AnnouncementMsgID: announce.ID,
		}

		api.MustCallAndUnwrap(api.GetSurveyAPI().DefineSurveyInfo, defineChanIn, func(_ *any) {

			_, err := u.session.ChannelMessageSend(thread.ID, fmt.Sprintf("Voting thread for survey %s create successfully!", t.ID))
			if err != nil {
				response.ErrorResponse(err, true, u.session, u.interaction)
				return
			}

			response.WithMessage("Survey opened successfully!", true, u.session, u.interaction)

		}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[any](u.session, u.interaction))

	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[entity.Survey](u.session, u.interaction))

}
