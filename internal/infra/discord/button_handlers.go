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

				if len(data) < 2 {
					return
				}
				MustBeMainGuild(MustBeMod(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					vote := usecase.NewConfirmVoteUsecase(s, i)
					go vote.Execute(api.ConfirmVoteUsecaseInput{
						Executer:  i.Member.User.ID,
						Token:     data[1],
						MessageID: i.Message.ID,
					})
				}))(s, i)
			}
			if data[0] == "denyvote" {
				if len(data) < 2 {
					return
				}
				MustBeMainGuild(MustBeMod(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					vote := usecase.NewVoteDenyCase(s, i)
					go vote.Execute(api.DenyVoteInput{
						Token: data[1],
					})
				}))(s, i)
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
				if len(data) < 4 {
					return
				}
				vote := usecase.NewRankViewCase(s, i)
				vote.Execute(api.RankListCaseInput{
					ItemName:  data[3],
					MessageID: &data[1],
					ChannelID: &data[2],
					Prev:      true,
				}, false)
			}

			if data[0] == "ranknextpage" {
				if len(data) < 3 {
					return
				}
				vote := usecase.NewRankViewCase(s, i)
				vote.Execute(api.RankListCaseInput{
					ItemName:  data[3],
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
				}, true)
			}
		}
	}
)
