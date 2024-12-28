package usecase

import (
	"fmt"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/config"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/util"

	"github.com/bwmarrin/discordgo"
)

type RankUpdateCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewRankUpdateCase(s *discordgo.Session, i *discordgo.InteractionCreate) *RankUpdateCase {
	return &RankUpdateCase{
		session:     s,
		interaction: i,
	}
}

func (u *RankUpdateCase) Execute(input api.RankUpdateCaseInput) {
	api.MustCallAndUnwrap(api.GetItemAPI().RankUpdate, input, func(t *any) {
		err := util.SendMessage(u.session, config.MainConfig.Discord.Channels.RankUpdates, &discordgo.MessageSend{
			Content: fmt.Sprintf("<@&%s>: Rank for item ``%s`` has been updated!", config.MainConfig.Discord.Roles.Ranks, input.ItemName),
		})
		if err != nil {
			response.ErrorResponse(err, true, u.session, u.interaction)
			return
		}
		response.WithMessage("Rank updated successfully!", true, u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[any](u.session, u.interaction))
}
