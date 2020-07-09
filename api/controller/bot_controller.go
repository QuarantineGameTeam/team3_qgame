package controller

import (
	"log"
	"strings"

	"github.com/team3_qgame/api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const webHookPrefix = "/"

type Updater interface {
	CallbackQuery(update tgbotapi.Update)
	Messages(update tgbotapi.Update)
	SetUpdates(*tgbotapi.BotAPI, tgbotapi.UpdatesChannel)
}

type BotController struct {
	Bot *tgbotapi.BotAPI
}

func NewBotController(bot *api.Bot) *BotController {
	err := bot.InitiateBot()
	if err != nil {
		log.Fatalln("Initiate Bot failure. ", err.Error())
	}

	botAPI := bot.GetBotAPI()

	return &BotController{
		Bot: botAPI,
	}
}

func (b *BotController) StartWebHookListener(userUpd, fightUpd Updater) {
	updates := b.Bot.ListenForWebhook(webHookPrefix)

	userUpd.SetUpdates(b.Bot, updates)
	fightUpd.SetUpdates(b.Bot, updates)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				if strings.HasPrefix(update.Message.Command(), "f_") {
					fightUpd.Messages(update)
				} else {
					userUpd.Messages(update)
				}
			}
		}
		if update.CallbackQuery != nil {
			if strings.HasPrefix(update.CallbackQuery.Data, "f_") {
				fightUpd.CallbackQuery(update)
			} else {
				userUpd.CallbackQuery(update)
			}

		}
	}
}
