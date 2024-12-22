package api

import (
	"victo/wynnguardian-bot/internal/domain/config"

	"github.com/wynnguardian/common/entity"
)

type SurveyAPI interface {
	OpenSurvey(input OpenSurveyInput) (*APIResponse[entity.Survey], error)
	SendSurveyVote(input SurveyVoteInput) (*APIResponse[any], error)
	FindSurveys(surveyId FindSurveyInput) (*APIResponse[[]entity.Survey], error)
	DefineSurveyInfo(input DefineSurveyInfoInput) (*APIResponse[any], error)
	CloseSurvey(input SurveyCloseUsecaseInput) (*APIResponse[entity.Survey], error)
	CancelSurvey(input SurveyCancelUsecaseInput) (*APIResponse[entity.Survey], error)
	ApproveSurvey(input SurveyApproveCaseInput) (*APIResponse[entity.Survey], error)
	DiscardSurvey(input SurveyDiscardCaseInput) (*APIResponse[entity.Survey], error)
	StartVoting(input StartVotingUsecase) (*APIResponse[entity.SurveyVote], error)
	ConfirmVote(input ConfirmVoteUsecaseInput) (*APIResponse[entity.SurveyVote], error)
	DefineVoteMessage(input DefineVoteMessageInput) (*APIResponse[any], error)
}

type SurveyAPIImpl struct {
	SurveyAPI
}

func (a *SurveyAPIImpl) CallData() CallData {
	return CallData{
		Token: config.MainConfig.Private.Tokens.Self,
		Host:  config.MainConfig.Hosts.SurveyAPI,
	}
}

func (a *SurveyAPIImpl) OpenSurvey(input OpenSurveyInput) (*APIResponse[entity.Survey], error) {
	return NewCall[OpenSurveyInput, entity.Survey](a.CallData(), "surveyCreate", input).Post()
}

func (a *SurveyAPIImpl) FindSurveys(input FindSurveyInput) (*APIResponse[[]entity.Survey], error) {
	return NewCall[FindSurveyInput, []entity.Survey](a.CallData(), "findOpenSurvey", input).Post()
}

func (a *SurveyAPIImpl) SendSurveyVote(input SurveyVoteInput) (*APIResponse[any], error) {
	return NewCall[SurveyVoteInput, any](a.CallData(), "surveyVote", input).Post()
}

func (a *SurveyAPIImpl) StartVoting(input StartVotingUsecase) (*APIResponse[entity.SurveyVote], error) {
	return NewCall[StartVotingUsecase, entity.SurveyVote](a.CallData(), "createVote", input).Post()
}

func (a *SurveyAPIImpl) DefineSurveyInfo(input DefineSurveyInfoInput) (*APIResponse[any], error) {
	return NewCall[DefineSurveyInfoInput, any](a.CallData(), "defineSurveyInfo", input).Post()
}

func (a *SurveyAPIImpl) ConfirmVote(input ConfirmVoteUsecaseInput) (*APIResponse[entity.SurveyVote], error) {
	return NewCall[ConfirmVoteUsecaseInput, entity.SurveyVote](a.CallData(), "confirmVote", input).Post()
}

func (a *SurveyAPIImpl) DefineVoteMessage(input DefineVoteMessageInput) (*APIResponse[any], error) {
	return NewCall[DefineVoteMessageInput, any](a.CallData(), "defineVoteMessage", input).Post()
}

func (a *SurveyAPIImpl) CloseSurvey(input SurveyCloseUsecaseInput) (*APIResponse[entity.Survey], error) {
	return NewCall[SurveyCloseUsecaseInput, entity.Survey](a.CallData(), "closeSurvey", input).Post()
}

func (a *SurveyAPIImpl) CancelSurvey(input SurveyCancelUsecaseInput) (*APIResponse[entity.Survey], error) {
	return NewCall[SurveyCancelUsecaseInput, entity.Survey](a.CallData(), "cancelSurvey", input).Post()
}

func (a *SurveyAPIImpl) ApproveSurvey(input SurveyApproveCaseInput) (*APIResponse[entity.Survey], error) {
	return NewCall[SurveyApproveCaseInput, entity.Survey](a.CallData(), "approveSurvey", input).Post()
}

func (a *SurveyAPIImpl) DiscardSurvey(input SurveyDiscardCaseInput) (*APIResponse[entity.Survey], error) {
	return NewCall[SurveyDiscardCaseInput, entity.Survey](a.CallData(), "discardSurvey", input).Post()
}
