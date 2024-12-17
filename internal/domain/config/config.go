package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wynnguardian/common/utils"
)

type SurveyStatusConfig struct {
	Color   int32  `json:"color"`
	Icon    string `json:"icon"`
	Message string `json:"string"`
}

type SurveyEmbedsConfig struct {
	StatusConfig struct {
		Open     SurveyStatusConfig `json:"open"`
		Denied   SurveyStatusConfig `json:"denied"`
		Waiting  SurveyStatusConfig `json:"waiting"`
		Approved SurveyStatusConfig `json:"approved"`
	} `json:"status"`
	MinVotes int `json:"min_votes"`
}

type PrivateConfig struct {
	Tokens struct {
		Discord   string   `json:"discord"`
		Self      string   `json:"self"`
		Whitelist []string `json:"whitelist"`
	} `json:"tokens"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type HostsConfig struct {
	ItemsAPI  string `json:"items"`
	SurveyAPI string `json:"surveys"`
}

type DiscordConfig struct {
	Channels struct {
		SurveyAnnouncements   string `json:"survey_announcements"`
		SurveyPublicResults   string `json:"survey_public_results"`
		SurveyWaitingApproval string `json:"survey_waiting_approval"`
		VotesWaitingApproval  string `json:"votes_waiting_approval"`
		ErrorLog              string `json:"error_log"`
	} `json:"channels"`
	Roles struct {
		Admin string `json:"admin"`
	} `json:"roles"`
	MainGuild string `json:"main_guild"`
}

type Config struct {
	SurveyEmbeds SurveyEmbedsConfig
	Private      PrivateConfig
	Server       ServerConfig
	Hosts        HostsConfig
	Discord      DiscordConfig
}

var MainConfig *Config = &Config{}

func Load() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	mustReadAndAssign(pwd, "config/embed/survey_embeds.json", &MainConfig.SurveyEmbeds)
	mustReadAndAssign(pwd, "config/private.json", &MainConfig.Private)
	mustReadAndAssign(pwd, "config/hosts.json", &MainConfig.Hosts)
	mustReadAndAssign(pwd, "config/discord.json", &MainConfig.Discord)
	mustReadAndAssign(pwd, "config/server.json", &MainConfig.Server)
}

func mustReadAndAssign(pwd, relativeDir string, target interface{}) {
	f := utils.MustVal(os.ReadFile(fmt.Sprintf("%s/%s", pwd, relativeDir)))
	utils.Must(json.Unmarshal(f, &target))
}
