package discord

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/config"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/usecase"

	"github.com/bwmarrin/discordgo"
)

var (
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"criteria": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			switch options[0].Name {
			case "create":
				MustBeMainGuild(MustBeManager(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewCreateCriteriaCase(s, i)
					go uc.Execute(api.CreateCriteriaInput{
						ItemName:   options[0].Options[0].StringValue(),
						CriteriaId: options[0].Options[1].StringValue(),
						Default:    0,
					})
				}))(s, i)

			case "view":
				uc := usecase.NewViewCriteriaCase(s, i)
				go uc.Execute(api.FindCriteriaInput{
					ItemName: options[0].Options[0].StringValue(),
				})

			case "delete":
				MustBeMainGuild(MustBeManager(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewDeleteCriteriaCase(s, i)
					go uc.Execute(api.DeleteCriteriaInput{
						ItemName:   options[0].Options[0].StringValue(),
						CriteriaId: options[0].Options[1].StringValue(),
					})
				}))(s, i)

			case "update":
				MustBeMainGuild(MustBeManager(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewUpdateCriteriaCase(s, i)
					go uc.Execute(api.UpdateCriteriaInput{
						Value:      int(options[0].Options[2].IntValue()),
						ItemName:   options[0].Options[0].StringValue(),
						CriteriaId: options[0].Options[1].StringValue(),
					})
				}))(s, i)

			}
		},
		"rank": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			switch options[0].Name {

			case "update":
				MustBeMainGuild(MustBeMod(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewRankUpdateCase(s, i)
					go uc.Execute(api.RankUpdateCaseInput{
						ItemName: options[0].Options[0].StringValue(),
					})
				}))(s, i)

			case "view":
				uc := usecase.NewRankViewCase(s, i)
				go uc.Execute(api.RankListCaseInput{
					ItemName:  options[0].Options[0].StringValue(),
					MessageID: nil,
					ChannelID: &i.ChannelID,
					Prev:      false,
				}, true)
			}
		},
		"item": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			switch options[0].Name {

			case "weight":
				uc := usecase.NewItemWeighUsecase(s, i)
				go uc.Execute(api.WeightItemInput{
					ItemUTF16: options[0].Options[0].StringValue(),
				})

			case "authenticate":
				MustBeMainGuild(MustBeMod(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewItemAuthUsecase(s, i)
					go uc.Execute(api.AuthenticateItemInput{
						Item:       options[0].Options[0].StringValue(),
						MCOwnerUID: options[0].Options[1].StringValue(),
						DCOwnerUID: options[0].Options[2].StringValue(),
						Public:     options[0].Options[3].BoolValue(),
						Force:      options[0].Options[4].BoolValue(),
					})
				}))(s, i)

			case "track":
				uc := usecase.NewItemTrackCase(s, i)
				go uc.Execute(api.FindItemInput{
					TrackingCode: options[0].Options[0].StringValue(),
				})
			}
		},

		"survey": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			switch options[0].Name {

			case "list":
				uc := usecase.NewSurveyListCase(s, i)
				go uc.Execute(api.SurveyListInput{
					MessageID: nil,
					ChannelID: &i.ChannelID,
					Prev:      false,
				}, true)

			case "open":
				MustBeMainGuild(MustBeMod(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewSurveyOpenUsecase(s, i)
					go uc.Execute(api.OpenSurveyInput{
						ItemName:     options[0].Options[0].StringValue(),
						DurationDays: int(options[0].Options[1].IntValue()),
					})
				}))(s, i)

			case "approve":
				MustBeMainGuild(MustBeManager(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewSurveyApproveCase(s, i)
					go uc.Execute(api.SurveyApproveCaseInput{
						SurveyID: options[0].Options[0].StringValue(),
					})
				}))(s, i)

			case "discard":
				MustBeMainGuild(MustBeManager(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewSurveyDiscardUsecase(s, i)
					go uc.Execute(api.SurveyDiscardCaseInput{
						SurveyID: options[0].Options[0].StringValue(),
					})
				}))(s, i)

			case "close":
				MustBeMainGuild(MustBeManager(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewSurveyCloseUsecase(s, i)
					go uc.Execute(api.SurveyCloseUsecaseInput{
						ItemName: options[0].Options[0].StringValue(),
					})
				}))(s, i)

			case "cancel":
				MustBeMainGuild(MustBeManager(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewSurveyCancelUsecase(s, i)
					go uc.Execute(api.SurveyCancelUsecaseInput{
						ItemName: options[0].Options[0].StringValue(),
					})
				}))(s, i)

			case "fill":
				uc := usecase.NewStartVotingUsecase(s, i)
				go uc.Execute(api.StartVotingUsecase{
					UserID: i.Member.User.ID,
					Item:   options[0].Options[0].StringValue(),
				}, false)

			case "ban":
				uc := usecase.NewSurveyBanCase(s, i)
				MustBeMainGuild(MustBeMod(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					go uc.Execute(api.SurveyBanInput{
						UserID: options[0].Options[0].UserValue(s).ID,
						Reason: options[0].Options[1].StringValue(),
					})
				}))(s, i)

			case "unban":
				MustBeMainGuild(MustBeMod(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
					uc := usecase.NewSurveyUnbanCase(s, i)
					go uc.Execute(api.SurveyUnbanInput{
						UserID: options[0].Options[0].UserValue(s).ID,
					})
				}))(s, i)
			}
		},
	}
)

type Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)

func MustHaveRole(role string, next Handler) Handler {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		for _, r := range i.Member.Roles {
			if r == role {
				next(s, i)
				return
			}
		}
		response.UnauthorizedResponse(s, i)
	}
}

func MustBeManager(next Handler) Handler {
	return MustHaveRole(config.MainConfig.Discord.Roles.Manager, next)
}

func MustBeMod(next Handler) Handler {
	return MustHaveRole(config.MainConfig.Discord.Roles.Moderator, next)
}

func MustBeMainGuild(next Handler) Handler {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.GuildID == config.MainConfig.Discord.MainGuild {
			next(s, i)
			return
		}
		response.UnauthorizedResponse(s, i)
	}
}
