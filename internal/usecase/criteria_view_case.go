package usecase

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"
	"victo/wynnguardian-bot/internal/infra/cerrors"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/wynnguardian/common/entity"
)

type ViewCriteriaCase struct {
	session     *discordgo.Session
	interaction *discordgo.InteractionCreate
}

func NewViewCriteriaCase(s *discordgo.Session, i *discordgo.InteractionCreate) *ViewCriteriaCase {
	return &ViewCriteriaCase{
		session:     s,
		interaction: i,
	}
}

func (u *ViewCriteriaCase) Execute(input api.FindCriteriaInput) {
	api.MustCallAndUnwrap(api.GetItemAPI().FindCriteria, input, func(t *entity.ItemCriteria) {
		embed := embed.GetItemCriteriaEmbed(t)
		response.WithEmbed(embed, false, u.session, u.interaction)
	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[entity.ItemCriteria](u.session, u.interaction))
}
