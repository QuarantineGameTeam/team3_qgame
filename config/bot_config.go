package config

import (
	"flag"
	"log"
)

type BotConfig struct {
	Token string
	Port  int
}

const (
	token = "localhost"
	botPort  = 8443
)

func (b *BotConfig) InitBotConfig() {
	flag.StringVar(&b.Token, "bot_token", token, "telegram bot token")
	flag.IntVar(&b.Port, "bot_port", botPort, "telegram bot port")
	flag.Parse()
	log.Printf("app starts whith bot configs:\n token=%s ,\n port=%d", b.Token, b.Port)
}
