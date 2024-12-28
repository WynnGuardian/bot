package server

import (
	"victo/wynnguardian-bot/internal/infra/discord"
	"victo/wynnguardian-bot/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/wynnguardian/common/entity"
	"github.com/wynnguardian/common/response"
)

func handleVote(ctx *gin.Context) response.WGResponse {

	input := entity.SurveyVote{}
	if err := ctx.BindJSON(&input); err != nil {
		return response.ErrBadRequest
	}

	c := usecase.NewVoteReceivedCase(discord.Discord)
	return c.Execute(input)

	return response.Ok
}
