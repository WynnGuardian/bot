package cerrors

import (
	"errors"
	"fmt"
	"net/http"
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/response"

	"github.com/bwmarrin/discordgo"
)

func CatchAndLogInternal(s *discordgo.Session, i *discordgo.InteractionCreate) func(err error) {
	return func(err error) {
		if err != nil && response.ErrorResponse(err, true, s, i) != nil {
			fmt.Println("Couldn't log error: ", err.Error())
		}
	}
}

func CatchAndLogAPIError[T any](s *discordgo.Session, i *discordgo.InteractionCreate) func(err *api.APIResponse[T]) {
	return func(err *api.APIResponse[T]) {
		internal := err.Status == http.StatusInternalServerError
		if err.Status != 200 && !internal && response.ErrorResponse(errors.New(err.Message), internal, s, i) != nil {
			fmt.Println("Couldn't log error: ", err.Message)
		}
	}
}
