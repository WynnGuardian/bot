package util

import (
	"encoding/json"
	"fmt"
	"os"
)

type SurveyStatusConfig struct {
	Color   int32  `json:"color"`
	Icon    string `json:"icon"`
	Message string `json:"string"`
}

type SurveyConfig struct {
	StatusConfig struct {
		Open     SurveyStatusConfig `json:"open"`
		Denied   SurveyStatusConfig `json:"denied"`
		Waiting  SurveyStatusConfig `json:"waiting"`
		Approved SurveyStatusConfig `json:"approved"`
	} `json:"status"`
	MinVotes int `json:"min_votes"`
}

type Config struct {
	SurveyConfig SurveyConfig `json:"surveys"`
}

var MainConfig *Config

func LoadConfig() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := os.ReadFile(fmt.Sprintf("%s/%s", pwd, "config.json"))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(f, &MainConfig)
	if err != nil {
		panic(err)
	}

}
