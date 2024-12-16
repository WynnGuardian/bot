package main

import (
	"victo/wynnguardian-bot/internal/domain/api"
	"victo/wynnguardian-bot/internal/domain/config"
	"victo/wynnguardian-bot/internal/infra/discord"
	http "victo/wynnguardian-bot/server"
)

func main() {
	config.Load()
	api.Setup()
	go http.StartWebhookServer()
	discord.Init()
}
