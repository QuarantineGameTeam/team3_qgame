package controller

import (
	"gihub.com/team3_qgame/actions"
	"log"

	"gihub.com/team3_qgame/api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const webHookPrefix = "/"

type BotController struct {
	Bot         *tgbotapi.BotAPI
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

func (b *BotController) StartWebHookListener() {
	updates := b.Bot.ListenForWebhook(webHookPrefix)

	var uact actions.UserActions
	uact.Set(b.Bot, updates)

	for update := range updates {
		if update.CallbackQuery != nil {
			uact.CallbackQuery(update)
		}
		if update.Message != nil {
			uact.Messages(update)
		}

	}
}
