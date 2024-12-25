package embed

import (
	"fmt"
	"victo/wynnguardian-bot/internal/domain/config"

	"github.com/wynnguardian/common/enums"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

func GetSurveyAnnounceMessage(survey *entity.Survey) *discordgo.MessageSend {

	var status config.SurveyStatusConfig
	switch survey.Status {
	case enums.SURVEY_APPROVED:
		status = config.MainConfig.SurveyEmbeds.StatusConfig.Approved
	case enums.SURVEY_DENIED:
		status = config.MainConfig.SurveyEmbeds.StatusConfig.Denied
	case enums.SURVEY_WAITING_APPROVAL:
		status = config.MainConfig.SurveyEmbeds.StatusConfig.Waiting
	case enums.SURVEY_OPEN:
		status = config.MainConfig.SurveyEmbeds.StatusConfig.Open
	}

	msg := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       fmt.Sprintf("Criteria Survey: %s", survey.ItemName),
				Description: fmt.Sprintf("Survey ends <t:%d:R>\nYou can fill it by typing ``/survey fill %s`` or clicking the button bellow.", survey.Deadline.Unix(), survey.ItemName),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Status",
						Value:  status.Message,
						Inline: true,
					},
				},
				Color: int(status.Color),
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: status.Icon,
				},
			},
		},
	}

	if survey.Status == enums.SURVEY_OPEN {
		msg.Components = []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						CustomID: fmt.Sprintf("surveyfill_%s", survey.ItemName),
						Label:    "Vote",
						Style:    discordgo.PrimaryButton,
					},
				},
			},
		}
	}

	return msg
}

func GetVoteCreateMessage(url, item string) *discordgo.MessageSend {
	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: fmt.Sprintf("Click the button to fill the ``%s`` survey", item),
				Color: 0xb700ff,
				Image: &discordgo.MessageEmbedImage{
					URL:    "http://198.7.123.203/weapons/dagger.water2.webp",
					Width:  300,
					Height: 300,
				},
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "Vote",
						Style: discordgo.LinkButton,
						URL:   url,
					},
				},
			},
		},
	}
}

func GetVoteConfirmedMessage(survey, item string) *discordgo.MessageSend {
	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: "Your vote was contabilized!",
				Color: 0xb700ff,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Item",
						Value:  item,
						Inline: false,
					},
					{
						Name:   "Survey",
						Value:  survey,
						Inline: true,
					},
				},
			},
		},
	}
}
