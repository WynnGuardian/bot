package discord

import (
	"strings"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/usecase"

	"github.com/bwmarrin/discordgo"
)

var (
	buttonHandlers = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionMessageComponent {
			data := strings.Split(i.MessageComponentData().CustomID, "_")
			if data[0] == "confirmvote" {
				if len(data) < 3 {
					return
				}
				vote := usecase.NewConfirmVoteUsecase(s, i)
				vote.Execute(api.ConfirmVoteUsecaseInput{
					Executer:  i.Member.User.ID,
					UserID:    data[1],
					Survey:    data[2],
					MessageID: i.Message.ID,
				})
			}
			if data[0] == "nextpage" {
				if len(data) < 3 {
					return
				}
				vote := usecase.NewSurveyListCase(s, i)
				vote.Execute(api.SurveyListInput{
					MessageID: &data[1],
					ChannelID: &data[2],
					Prev:      false,
				}, false)
			}
			if data[0] == "previouspage" {
				if len(data) < 3 {
					return
				}
				vote := usecase.NewSurveyListCase(s, i)
				vote.Execute(api.SurveyListInput{
					MessageID: &data[1],
					ChannelID: &data[2],
					Prev:      true,
				}, false)
			}

			if data[0] == "rankpreviouspage" {
				if len(data) < 3 {
					return
				}
				vote := usecase.NewSurveyListCase(s, i)
				vote.Execute(api.SurveyListInput{
					MessageID: &data[1],
					ChannelID: &data[2],
					Prev:      true,
				}, false)
			}

			if data[0] == "ranknextpage" {
				if len(data) < 3 {
					return
				}
				vote := usecase.NewSurveyListCase(s, i)
				vote.Execute(api.SurveyListInput{
					MessageID: &data[1],
					ChannelID: &data[2],
					Prev:      false,
				}, false)
			}

			if data[0] == "surveyfill" {
				if len(data) < 2 {
					return
				}

				uc := usecase.NewStartVotingUsecase(s, i)
				uc.Execute(api.StartVotingUsecase{
					Item:   data[1],
					UserID: i.Member.User.ID,
				})
			}
		}
	}
)
