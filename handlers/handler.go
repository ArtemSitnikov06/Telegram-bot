package handlers

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sort"
)

var chats_user_msgCount map[int64]map[int64]int
var chats_login_msgCount map[int64]map[int64]string

func HandlerUpdates(update tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	chats_user_msgCount = make(map[int64]map[int64]int)
	chats_login_msgCount = make(map[int64]map[int64]string)
	for update := range update {
		if update.Message != nil {

			chat_id := update.Message.Chat.ID
			user_id := update.Message.From.ID
			user_login := update.Message.From.UserName

			if _, ok := chats_user_msgCount[chat_id]; !ok {
				chats_user_msgCount[chat_id] = make(map[int64]int)
			}

			if _, ok := chats_login_msgCount[chat_id]; !ok {
				chats_login_msgCount[chat_id] = make(map[int64]string)
			}

			chats_login_msgCount[chat_id][user_id] = user_login

			switch Parse(update.Message.Text) {

			case "/mystats":

				count := chats_user_msgCount[chat_id][user_id]
				text := fmt.Sprintf("@%s написал %d сообщений", user_login, count)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)
			case "/tagall":

				text := fmt.Sprintf("@%s этот долбаеб всех тегнул\n\n", user_login)

				userLogins := chats_login_msgCount[chat_id]

				for _, userLogin := range userLogins {
					text += fmt.Sprintf("@%s  \n", userLogin)
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)

			case "/topusers":

				userCounts := chats_user_msgCount[chat_id]
				var stats []userStat
				for id, count := range userCounts {
					stats = append(stats, userStat{id, count})
				}

				sort.Slice(stats, func(i, j int) bool {
					return stats[i].count > stats[j].count
				})

				text := "🏆 Топ активных пользователей:\n"
				limit := 10
				if len(stats) < 10 {
					limit = len(stats)
				}

				for i := 0; i < limit; i++ {
					login := chats_login_msgCount[chat_id][stats[i].userId]
					text += fmt.Sprintf("%d. @%s — %d сообщений\n", i+1, login, chats_user_msgCount[chat_id][stats[i].userId])
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)

			default:

				chats_user_msgCount[chat_id][user_id]++

			}

		}
	}
}

func Parse(text string) string {
	renSlice := []rune(text)
	var runReturn []rune
	for _, r := range renSlice {
		if string(r) == "@" {
			break
		}
		runReturn = append(runReturn, r)
	}
	return string(runReturn)

}

type userStat struct {
	userId int64
	count  int
}
