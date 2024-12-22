package usecase

import (
	"sync"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/infra/cerrors"

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

	in := api.FindRankInput{
		ItemName: input.ItemName,
		Limit:    limit,
		Page:     int(page),
	}

	api.MustCallAndUnwrap(api.GetItemAPI().GetRank, in, func(t *[]entity.AuthenticatedItem) {

	}, cerrors.CatchAndLogInternal(u.session, u.interaction), cerrors.CatchAndLogAPIError[[]entity.AuthenticatedItem](u.session, u.interaction))
}
