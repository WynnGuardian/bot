package server

import (
	"time"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/infra/discord"
	"victo/wynnguardian-bot/internal/infra/util"
	"victo/wynnguardian-bot/internal/infra/visual/embed"

	"github.com/gin-gonic/gin"
	"github.com/wynnguardian/common/entity"
	"github.com/wynnguardian/common/response"
)

func handleVote(ctx *gin.Context) response.WGResponse {

	input := entity.SurveyVote{}
	if err := ctx.BindJSON(&input); err != nil {
		return response.ErrBadRequest
	}

	embed := embed.GetVoteEmbed(input)
	msg, err := util.SendVoteConfirmMessage(discord.Discord, input.Survey.ChannelID, input.Survey.ID, input.DiscordUserID, embed)
	if err != nil {
		return response.ErrInternalServerErr(err)
	}

	go func() {
		time.Sleep(500)
		_, err = api.GetSurveyAPI().DefineVoteMessage(api.DefineVoteMessageInput{
			SurveyID:  input.Survey.ID,
			UserID:    input.DiscordUserID,
			ChannelID: input.Survey.ChannelID,
			MessageID: msg,
		})

		if err != nil {
			util.LogError(err, input.Survey.ChannelID, "survey fill", discord.Discord)
		}
	}()

	return response.Ok
}
