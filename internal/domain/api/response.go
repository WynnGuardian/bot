package api

import "github.com/wynnguardian/common/entity"

type WeightResponse struct {
	StaticItem entity.ItemInstance `json:"item"`
	Criteria   entity.ItemCriteria `json:"criteria"`
	Weight     float64             `json:"weight"`
}

type AuthenticateItemResponse struct {
	TrackingCode string               `json:"tracking_code"`
	WynnItem     *entity.WynnItem     `json:"wynn_item"`
	Weight       float64              `json:"weight"`
	Item         *entity.ItemInstance `json:"item"`
}

type CreateVoteResponse struct {
	Token    string `json:"token"`
	SurveyID string `json:"survey_id"`
}
