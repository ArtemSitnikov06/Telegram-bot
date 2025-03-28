﻿package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NewBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return bot, nil
}
