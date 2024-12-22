package server

import (
	"victo/wynnguardian-bot/internal/infra/discord"
	"victo/wynnguardian-bot/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/wynnguardian/common/entity"
	"github.com/wynnguardian/common/response"
)

func handleSurveyEnd(ctx *gin.Context) response.WGResponse {
	input := entity.SurveyResult{}
	if err := ctx.BindJSON(&input); err != nil {
		return response.ErrBadRequest
	}
	return usecase.NewSurveyEndedCase(discord.Discord).Execute(input)
}
