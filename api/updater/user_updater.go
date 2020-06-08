package updater

import (
	"gihub.com/team3_qgame/actions"
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

type UpdateManager struct {
	bot     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
	user    *actions.User
}

func NewUpdateManager(user *actions.User) *UpdateManager {
	return &UpdateManager{
		user: user,
	}
}

func (u *UpdateManager) SetUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	u.bot = bot
	u.updates = updates

	u.user.SetUpdates(bot, updates)
}

func (u *UpdateManager) Messages(update tgbotapi.Update) {
	if update.Message.Text != "" {
		switch update.Message.Text {
		case "/start":
			u.user.MSStart(update)
		case "/register":
			u.user.MSRegistration(update)
		case "/information":
			u.user.MSIformation(update)
		case "/plus":
			u.user.StartClanSelection(update)
			u.user.ProcessClanSelection(update)
		default:
			u.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "WRONG COMMAND!"))
		}
	}
}

func (u *UpdateManager) CallbackQuery(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "dfasdf")
	switch update.CallbackQuery.Data {
	case "4":
		msg.Text = "You hit the '4' button!"
	}
	u.bot.Send(msg)
}
