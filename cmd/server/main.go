package main

import (
	"context"
	"github.com/krespix/cities-game/internal/bot"
	"github.com/krespix/cities-game/internal/processors/message"
	"github.com/krespix/cities-game/internal/services/cities"
	"log"
)

func main() {
	cityService := cities.New()
	cityService.InitCitiesMap()
	msgProc := message.New()
	tgbot, err := bot.New("5830837372:AAFeG4jdvnK6V7shuQW7S4YL5B6P9QiK2ao", msgProc)
	if err != nil {
		log.Fatal("catn init tg bot")
	}
	tgbot.Start(context.Background())
}
