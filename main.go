package main

import (
	"TelegramBot/bot"
	"TelegramBot/handlers"
	"log"
)
import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func main() {
	bot1, err := bot.NewBot("7549289800:AAG8za91NbEgucyenDWx01PsF3P1594lWkA")
	if err != nil {
		log.Panic(err)
	}

	bot1.Debug = true

	log.Printf("Authorized on account %s", bot1.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot1.GetUpdatesChan(u)

	commands := []tgbotapi.BotCommand{
		{Command: "/mystats", Description: "Посмотреть свою статистику"},
		{Command: "/topusers", Description: "Посмотреть топ пользователей"},
		{Command: "/tagall", Description: "Отметить всех"},
	}

	cfg := tgbotapi.NewSetMyCommands(commands...)

	_, err1 := bot1.Request(cfg)
	if err1 != nil {
		log.Panic(err)
	}
	handlers.HandlerUpdates(updates, bot1)

}
