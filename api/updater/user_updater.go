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
	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			u.user.CStart(update)
		case "register":
			u.user.CRegistration(update)
		case "delete":
			u.user.CDelete(update)
		case "me":
			u.user.CGetUserInfo(update)
		case "allusers":
			u.user.CGetAllUsers(update)
		case "help":
			u.user.CHelp(update)
		case "rename":
			u.user.CNameUpdate(update)
		case "changeteam":
			u.user.CStartTeamSelection(update)
			u.user.TeamChange(update)
		case "startfight":
			u.user.CStartFightKb(update)
			u.user.StartFight(update)
			u.user.KbAttack(update)
			u.user.AttackCallBack(update)
		default:
			u.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "WRONG COMMAND!"))
		}
	} else if update.Message != nil {
		u.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
	}
}

func (u *UpdateManager) CallbackQuery(update tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
	switch update.CallbackQuery.Data {
	case "4":
		msg.Text = "You hit the '4' button!"
	}
	u.bot.Send(msg)
}
