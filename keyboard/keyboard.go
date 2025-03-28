package keyboard

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CreateCommands() tgbotapi.ReplyKeyboardMarkup {
	commandKeyBoard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/start"),
			tgbotapi.NewKeyboardButton("/myStats"),
		),
	)
	return commandKeyBoard
}
