package bot

import (
	"flag"
)

type BConfig struct {
	Token       string
	WebHookLink string
	Port        int
}

const (
	token       = "1096824996:AAE-YL-XR97B6D1trRl06fSgGNX6H7gs1VE"
	webHookLink = "https://bc25db00211a.ngrok.io"
	botPort     = 8443
)

func (b *BConfig) InitBotConfig() {
	flag.StringVar(&b.Token, "bot_token", token, "telegram bot token")
	flag.StringVar(&b.WebHookLink, "bot_web_hook_link", webHookLink, "telegram bot web hook link")
	flag.IntVar(&b.Port, "bot_port", botPort, "telegram bot port")
}
