package main

import (
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
}

const (
	BOT_TOKEN = "Token bot"
	BOT_PORT  = "8090"
	WEB_HOOK  = "https://xxx" // ngrok генерує
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(BOT_TOKEN)
)

func setWebhook(bot *tgbotapi.BotAPI) {
	bot.SetWebhook(tgbotapi.NewWebhook(WEB_HOOK))
}

var numberKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "1 value"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2 value"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3 value"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4 value"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5 value"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6 value"),
	),
)

func GetUser(msg *tgbotapi.Message) User {
	user := User{Id: msg.Chat.ID, FirstName: msg.Chat.FirstName}
	return user
}

func getUpdates(bot *tgbotapi.BotAPI) {
	setWebhook(bot)
	updates := bot.ListenForWebhook("/")
	for update := range updates {

		if update.CallbackQuery != nil {
			user := GetUser(update.CallbackQuery.Message)
			fmt.Print(user)
			bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
		}

		if update.Message != nil {

			user := GetUser(update.Message)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello,"+user.FirstName)

			switch update.Message.Text {
			case "/start":
				user = GetUser(update.Message)
				bot.Send(msg)
			case "/me":
				msg.Text = "Your rank is: 456"
				bot.Send(msg)
			case "open":
				msg.ReplyMarkup = numberKeyboard
				bot.Send(msg)
			case "close":
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

			}
		}

	}
}

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":"+BOT_PORT, nil))
	}()

	getUpdates(NewBot)

}
