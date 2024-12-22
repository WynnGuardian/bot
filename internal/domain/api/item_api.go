package api

import (
	"victo/wynnguardian-bot/internal/domain/config"

	"github.com/wynnguardian/common/entity"
)

type ItemsAPI interface {
	API
	WeightItem(input WeightItemInput) (*APIResponse[WeightResponse], error)
	AuthenticateItem(input AuthenticateItemInput) (*APIResponse[AuthenticateItemResponse], error)
	FindCriteria(input FindCriteriaInput) (*APIResponse[entity.ItemCriteria], error)
	RankUpdate(input RankUpdateCaseInput) (*APIResponse[any], error)
	GetRank(input FindRankInput) (*APIResponse[[]entity.AuthenticatedItem], error)
}

type ItemsAPIImpl struct {
	ItemsAPI
}

func (a *ItemsAPIImpl) CallData() CallData {
	return CallData{
		Token: config.MainConfig.Private.Tokens.Self,
		Host:  config.MainConfig.Hosts.ItemsAPI,
	}
}

func (a *ItemsAPIImpl) WeightItem(input WeightItemInput) (*APIResponse[WeightResponse], error) {
	return NewCall[WeightItemInput, WeightResponse](a.CallData(), "itemWeigh", input).Post()
}

func (a *ItemsAPIImpl) AuthenticateItem(input AuthenticateItemInput) (*APIResponse[AuthenticateItemResponse], error) {
	return NewCall[AuthenticateItemInput, AuthenticateItemResponse](a.CallData(), "itemAuthenticate", input).Post()
}

func (a *ItemsAPIImpl) FindCriteria(input FindCriteriaInput) (*APIResponse[entity.ItemCriteria], error) {
	return NewCall[FindCriteriaInput, entity.ItemCriteria](a.CallData(), "findCriteria", input).Post()
}

func (a *ItemsAPIImpl) RankUpdate(input RankUpdateCaseInput) (*APIResponse[any], error) {
	return NewCall[RankUpdateCaseInput, any](a.CallData(), "rankUpdate", input).Post()
}

func (a *ItemsAPIImpl) GetRank(input FindRankInput) (*APIResponse[[]entity.AuthenticatedItem], error) {
	return NewCall[FindRankInput, []entity.AuthenticatedItem](a.CallData(), "getRank", input).Post()
}
