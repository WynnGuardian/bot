package server

import (
	"fmt"
	"log"
	"victo/wynnguardian-bot/internal/domain/config"

	"github.com/gin-gonic/gin"
	"github.com/wynnguardian/common/response"
)

func StartWebhookServer() {
	engine := gin.Default()

	engine.POST("/vote", parse(auth(handleVote)))
	engine.POST("/surveyEnd", parse(auth(handleSurveyEnd)))
	engine.POST("/surveyApproval", parse(auth(handleSendSurveyToApproval)))

	err := engine.Run(fmt.Sprintf(":%d", config.MainConfig.Server.Port))
	if err != nil {
		log.Fatalf("Couldn't start HTTP server: %s\n", err.Error())
		return
	}
	fmt.Println("Listening on port ", config.MainConfig.Server.Port)
}

func auth(hf func(ctx *gin.Context) response.WGResponse) func(ctx *gin.Context) response.WGResponse {
	return func(ctx *gin.Context) response.WGResponse {
		authorized := config.MainConfig.Private.Tokens.Whitelist
		for _, w := range authorized {
			if w == ctx.GetHeader("Authorization") {
				return hf(ctx)
			}
		}
		return response.ErrUnauthorized
	}
}

func parse(hf func(ctx *gin.Context) response.WGResponse) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := hf(ctx)
		if resp.Body == "" {
			resp.Body = "{}"
		}
		ctx.JSON(resp.Status, gin.H{"status": resp.Status, "message": resp.Message, "body": resp.Body})
	}
}
