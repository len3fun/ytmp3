package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"ytmp3/pkg/telegram"
)

func main() {
	token, found := os.LookupEnv("BOT_TOKEN")
	if !found {
		log.Fatal("BOT_TOKEN variable hasn't been set")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
