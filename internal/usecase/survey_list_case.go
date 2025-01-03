package usecase

import (
	"fmt"
	"log"
	"sync"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
	"github.com/wynnguardian/common/enums"
)

var (
	limit                 = 10
	surveyListPageControl = &sync.Map{}
)

type SurveyListCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewSurveyListCase(s *discordgo.Session, i *discordgo.InteractionCreate) *SurveyListCase {
	return &SurveyListCase{
		session:     s,
		interaction: i,
	}
}

func (u *SurveyListCase) Execute(input api.SurveyListInput, first bool) {
	page := int8(1)
	if input.MessageID != nil {
		p, ok := surveyListPageControl.Load(*input.MessageID)
		if !ok {
			p = int8(1)
		}
		page = (p.(int8))
	}
	if !first {
		if input.Prev {
			if page <= 1 {
				return
			}
			page--
		} else {
			if page >= 126 {
				return
			}
			page++
		}
	}
	in := api.FindSurveyInput{
		Status: int8(enums.SURVEY_OPEN),
		Limit:  int8(limit),
		Page:   page,
	}
	api.MustCallAndUnwrap(
		api.GetSurveyAPI().FindSurveys,
		in,
		func(t *[]entity.Survey) {
			if input.MessageID == nil {

				err := u.session.InteractionRespond(u.interaction.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				})

				if err != nil {
					response.ErrorResponse(err, true, u.session, u.interaction)
					return
				}

				msg, err := u.session.FollowupMessageCreate(u.interaction.Interaction, true, &discordgo.WebhookParams{
					Content: "Listing surveys:",
				})
				if err != nil {
					response.ErrorResponse(err, true, u.session, u.interaction)
					return
				}

				channelId := msg.ChannelID
				msgId := msg.ID
				edited := &discordgo.MessageEdit{
					Channel: channelId,
					ID:      msgId,
					Embeds:  &[]*discordgo.MessageEmbed{embed.GetSurveyListEmbed(*t)},
					Components: &[]discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label:    "Previous",
									Style:    discordgo.PrimaryButton,
									CustomID: fmt.Sprintf("previouspage_%s_%s", msgId, channelId),
								},
								discordgo.Button{
									Label:    "Next",
									Style:    discordgo.PrimaryButton,
									CustomID: fmt.Sprintf("nextpage_%s_%s", msgId, channelId),
								},
							},
						},
					},
				}

				_, err = u.session.ChannelMessageEditComplex(edited)
				if err != nil {
					log.Println("Erro ao editar a mensagem com botão:", err)
					return
				}

				surveyListPageControl.Store(msgId, int8(page))

				return
			}
			_, err := u.session.ChannelMessageEditEmbed(*input.ChannelID, *input.MessageID, embed.GetSurveyListEmbed(*t))
			if err != nil {
				response.ErrorResponse(err, true, u.session, u.interaction)
				return
			}
			surveyListPageControl.Store(*input.MessageID, int8(page))
			u.session.InteractionRespond(u.interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredMessageUpdate,
			})
		},
		cerrors.CatchAndLogInternal(u.session, u.interaction),
		cerrors.CatchAndLogAPIError[[]entity.Survey](u.session, u.interaction),
	)
}
