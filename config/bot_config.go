package config

import (
	"flag"
	"log"
)

type BotConfig struct {
	BotToken string
	BotPort  int
	BotWebHookUrl string
}

const (
	token   = "1218266837:AAF-Z-gn4JlWpv5Fq-x1ReiHb8nhfZhm7aY"
	botPort = 8443
	webHookUrl = "https://fe149b8861de.ngrok.io"
)

func (b *BotConfig) InitBotConfig() {
	flag.StringVar(&b.BotToken, "bot_token", token, "telegram bot token")
	flag.IntVar(&b.BotPort, "bot_port", botPort, "telegram bot port")
	flag.StringVar(&b.BotWebHookUrl, "bot_url", webHookUrl, "telegram bot url")
	log.Printf("app starts whith bot configs:\n token=%s ,\n port=%d ,\n url=%s\n" ,
		b.BotToken, b.BotPort, b.BotWebHookUrl)
}
