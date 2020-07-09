package updater

import (
	"github.com/team3_qgame/actions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type UserUpdater struct {
	bot     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
	fight   *actions.Fight
}

func NewFightUpdater(fight *actions.Fight) *UserUpdater {
	return &UserUpdater{
		fight: fight,
	}
}

func (u *UserUpdater) SetUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	u.bot = bot
	u.updates = updates

	u.fight.SetUpdates(bot, updates)
}

func (u *UserUpdater) Messages(update tgbotapi.Update) {
	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "f_start":
			u.fight.Fight(update)
		default:
			u.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "WRONG COMMAND!"))
		}
	} else if update.Message != nil {
		u.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
	}
}

func (u *UserUpdater) CallbackQuery(update tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
	switch update.CallbackQuery.Data {
	case "4":
		msg.Text = "You hit the '4' button!"
	}
	u.bot.Send(msg)
}
