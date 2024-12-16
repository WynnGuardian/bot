package embed

import (
	"fmt"
	"victo/wynnguardian-bot/internal/domain/config"

	"github.com/wynnguardian/common/enums"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

func GetSurveyAnnounceEmbed(survey *entity.Survey) *discordgo.MessageEmbed {

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

	return &discordgo.MessageEmbed{
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
	}
}

func GetVoteCreateEmbed(url, item string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Click below to fill %s survey", item),
		Description: fmt.Sprintf("[Vote URL](%s)", url),
		Color:       0xb700ff,
	}
}

func GetVoteConfirmedEmbed(survey, item string) *discordgo.MessageEmbed {

	return &discordgo.MessageEmbed{
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
	}
}
