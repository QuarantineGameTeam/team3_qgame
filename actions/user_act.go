package actions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

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


type UserActions struct {
	bot     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
}

func (u *UserActions) Set(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	u.bot = bot
	u.updates = updates
}

func (u *UserActions) Messages(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "<empty>")

	switch update.Message.Text {
	case "/me":
		msg.Text = "Your rank is: 567"
	case "open":
		msg.ReplyMarkup = numericKeyboard
	case "close":
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	case "plus":
	}
	_, _ = u.bot.Send(msg)
}

func (u *UserActions) CallbackQuery(update tgbotapi.Update) {
	//fmt.Print(user)
	//
	//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
	//switch update.CallbackQuery.Data {
	//case "4":
	//	msg.Text = "You hit the '4' button!"
	//}
	//bot.Send(msg)
}


