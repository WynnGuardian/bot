package api

type OpenSurveyInput struct {
	ItemName     string `json:"item_name"`
	DurationDays int    `json:"deadline"`
}

type SurveyVoteInput struct {
	Survey string            `json:"survey_id"`
	User   string            `json:"user_id"`
	Votes  []SurveyVoteEntry `json:"votes"`
}

type SurveyVoteEntry struct {
	Stat  string  `json:"stat"`
	Value float64 `json:"value"`
}

type FindSurveyInput struct {
	Id       *string `json:"id"`
	ItemName *string `json:"item_name"`
	Status   int8    `json:"status"`
	Limit    int8    `json:"limit"`
	Page     int8    `json:"page"`
}

type WeightItemInput struct {
	ItemUTF16 string `json:"item_utf16"`
}

type AuthenticateItemInput struct {
	Item       string `json:"item_utf16"`
	MCOwnerUID string `json:"owner_mc_uid"`
	DCOwnerUID string `json:"owner_dc_uid"`
	Public     bool   `json:"public_info"`
}

type FindCriteriaInput struct {
	ItemName string `json:"item_name"`
}

type ItemUTF16Input struct {
	Item string `json:""`
}

type StartVotingUsecase struct {
	Item   string `json:"item_name"`
	UserID string `json:"user_dc_id"`
}

type DefineSurveyInfoInput struct {
	Survey            string `json:"survey_id"`
	ChannelID         string `json:"channel_id"`
	AnnouncementMsgID string `json:"announcement_message_id"`
}

type ConfirmVoteUsecaseInput struct {
	Executer  string `json:"executer"`
	UserID    string `json:"user_dc_id"`
	Survey    string `json:"survey_id"`
	MessageID string `json:"message_id"`
	ChannelID string `json:"channel_id"`
}

type DefineVoteMessageInput struct {
	SurveyID  string `json:"survey_id"`
	UserID    string `json:"user_dc_id"`
	ChannelID string `json:"channel_id"`
	MessageID string `json:"message_id"`
}

type SurveyCloseUsecaseInput struct {
	ItemName string `json:"item_name"`
}

type SurveyCancelUsecaseInput struct {
	ItemName string `json:"item_name"`
}

type SurveyListInput struct {
	MessageID *string `json:"message_id"`
	ChannelID *string `json:"channel_id"`
	Prev      bool    `json:"previous"`
}

type SurveyApproveCaseInput struct {
	SurveyID string `json:"survey_id"`
}

type SurveyDiscardCaseInput struct {
	SurveyID string `json:"survey_id"`
}

type RankUpdateCaseInput struct {
	ItemName string `json:"item_name"`
}

type RankListCaseInput struct {
	ItemName  string  `json:"item_name"`
	MessageID *string `json:"message_id"`
	ChannelID *string `json:"channel_id"`
	Prev      bool    `json:"previous"`
}

type FindRankInput struct {
	ItemName string `json:"item_name"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}

type CreateCriteriaInput struct {
	ItemName   string  `json:"item_name"`
	Default    float64 `json:"default"`
	CriteriaId string  `json:"criteria_id"`
}

type DeleteCriteriaInput struct {
	ItemName   string `json:"item_name"`
	CriteriaId string `json:"criteria_id"`
}

type FindCriteriaByNameInput struct {
	ItemName string `json:"item_name"`
}

type UpdateCriteriaInput struct {
	ItemName   string `json:"item_name"`
	CriteriaId string `json:"criteria_id"`
	Value      int    `json:"value"`
}
