package usecase

import (
	"log"
	"sync"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

var (
	rankLimit           = 10
	rankViewPageControl = &sync.Map{}
)

type RankViewCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewRankViewCase(s *discordgo.Session, i *discordgo.InteractionCreate) *RankViewCase {
	return &RankViewCase{
		session:     s,
		interaction: i,
	}
}

func (u *RankViewCase) Execute(input api.RankListCaseInput, first bool) {

	page := int8(1)
	if input.MessageID != nil {
		p, ok := rankViewPageControl.Load(*input.MessageID)
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

	in := api.FindRankInput{
		ItemName: input.ItemName,
		Limit:    rankLimit,
		Page:     int(page),
	}

	api.MustCallAndUnwrap(
		api.GetItemAPI().GetRank,
		in,
		func(t *[]entity.AuthenticatedItem) {
			if input.MessageID == nil {

				err := u.session.InteractionRespond(u.interaction.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				})

				if err != nil {
					response.ErrorResponse(err, true, u.session, u.interaction)
					return
				}

				msg, err := u.session.FollowupMessageCreate(u.interaction.Interaction, true, &discordgo.WebhookParams{
					Content: "Listing rank:",
				})
				if err != nil {
					response.ErrorResponse(err, true, u.session, u.interaction)
					return
				}

				channelId := msg.ChannelID
				msgId := msg.ID

				messageSend := embed.GetRankListMessage(*t, in.ItemName, msgId, channelId, in.Limit, in.Page-1, u.session)

				edited := &discordgo.MessageEdit{
					Channel:    channelId,
					ID:         msgId,
					Embeds:     &messageSend.Embeds,
					Components: &messageSend.Components,
				}

				_, err = u.session.ChannelMessageEditComplex(edited)
				if err != nil {
					log.Println("Erro ao editar a mensagem com botão:", err)
					return
				}

				rankViewPageControl.Store(msgId, int8(page))

				return
			}
			_, err := u.session.ChannelMessageEditEmbed(*input.ChannelID, *input.MessageID, embed.GetRankListMessage(*t, in.ItemName, *input.MessageID, *input.ChannelID, in.Limit, in.Page-1, u.session).Embeds[0])
			if err != nil {
				response.ErrorResponse(err, true, u.session, u.interaction)
				return
			}
			rankViewPageControl.Store(*input.MessageID, int8(page))
			u.session.InteractionRespond(u.interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredMessageUpdate,
			})
		},
		cerrors.CatchAndLogInternal(u.session, u.interaction),
		cerrors.CatchAndLogAPIError[[]entity.AuthenticatedItem](u.session, u.interaction),
	)

}
