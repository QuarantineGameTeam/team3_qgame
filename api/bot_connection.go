package api

import (
	"fmt"
	"gihub.com/team3_qgame/config/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
)

type Bot struct {
	botAPI *tgbotapi.BotAPI
	config *bot.BConfig
}

func NewBot(botConfig *bot.BConfig) *Bot {
	return &Bot{
		config: botConfig,
	}
}

func (b *Bot) InitiateBot() error {
	err := b.newBotAPI()
	if err != nil {
		return err
	}

	_, err = b.setWebHook()
	if err != nil {
		return err
	}

	go b.startBotServer()

	return nil
}

func (b *Bot) GetBotAPI() *tgbotapi.BotAPI {
	return b.botAPI
}

func (b *Bot) newBotAPI() error {
	bot, err := tgbotapi.NewBotAPI(b.config.Token)
	if err != nil {
		return err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	b.botAPI = bot

	return nil
}

func (b *Bot) setWebHook() (tgbotapi.APIResponse, error) {
	newWebHook := tgbotapi.NewWebhook(b.config.WebHookLink)
	APIResponse, err := b.botAPI.SetWebhook(newWebHook)
	if err != nil {
		return APIResponse, err
	}
	return APIResponse, err
}

func (b *Bot) startBotServer() {
	log.Fatalln(http.ListenAndServe(fmt.Sprint(":", b.config.Port), nil))
}
