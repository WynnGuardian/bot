package usecase

import (
	"victo/wynnguardian-bot/internal/domain/config"
	"victo/wynnguardian-bot/internal/infra/util"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
	"github.com/wynnguardian/common/response"
)

type SurveyApprovalCase struct {
	s *discordgo.Session
}

func NewSurveyApprovalCase(s *discordgo.Session) *SurveyApprovalCase {
	return &SurveyApprovalCase{
		s: s,
	}
}

func (u *SurveyApprovalCase) Execute(input entity.SurveyResult) response.WGResponse {

	msg := embed.GetSurveyAnnounceMessage(input.Survey)
	edit := &discordgo.MessageEdit{
		ID:         input.Survey.AnnouncementMessageID,
		Channel:    config.MainConfig.Discord.Channels.SurveyAnnouncements,
		Components: &msg.Components,
		Embeds:     &msg.Embeds,
	}

	u.s.ChannelMessageEditComplex(edit)

	embed := embed.GetSurveyApprovalEmbed(&input)

	err := util.SendEmbedMessage(u.s, config.MainConfig.Discord.Channels.SurveyWaitingApproval, embed)
	if err != nil {
		return response.ErrInternalServerErr(err)
	}
	return response.Ok
}
