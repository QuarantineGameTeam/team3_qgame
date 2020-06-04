package main

import (
	"fmt"
	"log"
	"net/http"
	"rgb/conf"
	"rgb/db"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(conf.BOT_TOKEN)
)

func setWebhook(bot *tgbotapi.BotAPI) {
	bot.SetWebhook(tgbotapi.NewWebhook(conf.WEB_HOOK))
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1", "http://golang.org"),
		tgbotapi.NewInlineKeyboardButtonSwitch("2", "open 2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "33"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func GetUser(msg *tgbotapi.Message) conf.User {
	user := conf.User{Id: msg.Chat.ID, FirstName: msg.Chat.FirstName}
	fmt.Println(user)
	return user
}

func getUpdates(bot *tgbotapi.BotAPI) {
	setWebhook(bot)
	updates := bot.ListenForWebhook("/")
	for update := range updates {
		var user conf.User
		if update.CallbackQuery != nil {

			fmt.Print(user)

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			switch update.CallbackQuery.Data {
			case "4":
				msg.Text = "You hit the '4' button!"
			}
			bot.Send(msg)
		}
		if update.Message != nil {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "<empty>")

			fmt.Println(user)

			if update.Message.Text != "/start" {
				user = db.GetUser(strconv.Itoa(int(update.Message.Chat.ID)))
				msg.Text = "Привіт, " + user.FirstName
			}

			switch update.Message.Text {
			case "/start":
				user = GetUser(update.Message)
				db.SaveUser(&user)
			case "/me":
				msg.Text = "Your rank is: 567"
			case "open":
				msg.ReplyMarkup = numericKeyboard
			case "close":
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			case "plus":
				user.Counter++
				db.SaveUser(&user)
				msg.Text = strconv.Itoa(user.Counter)
			}
			bot.Send(msg)
		}
		db.SaveUser(&user)

	}
}

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":"+conf.BOT_PORT, nil))
	}()

	getUpdates(NewBot)
}
