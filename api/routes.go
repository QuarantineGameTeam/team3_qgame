package api

import (
	"flag"
	"gihub.com/team3_qgame/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

/*
	Routing or router in web development is a mechanism where requests are routed to the code that handles them.
  	To put simply, in the Router you determine what should happen when a user visits a certain page or ask certain
  	information from our bot.
*/

type Bot struct {
	botAPI *tgbotapi.BotAPI
	config *config.BotConfig
}

func NewBot(botConfig *config.BotConfig) *Bot {
	botConfig.InitBotConfig()
	log.Printf("app starts whith bot configs:\n token=%s ,\n port=%d ,\n url=%s\n" ,
		botConfig.BotToken, botConfig.BotPort, botConfig.BotWebHookUrl)
	flag.Parse()
	return &Bot{
		config: botConfig,
	}
}

func (b *Bot) GetNewBotAPI() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(b.config.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}


func (b *Bot) GetUpdate(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	bot.SetWebhook(tgbotapi.NewWebhook(b.config.BotWebHookUrl))
	updates := bot.ListenForWebhook("/")
	return updates
}

