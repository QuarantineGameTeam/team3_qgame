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
	token       = "716959869:AAHHXSfUrr87c2ldUyb8G-ZjG4eDwqxO_tQ"
	webHookLink = "https://919167109a9a.ngrok.io"
	botPort     = 8443
)

func (b *BConfig) InitBotConfig() {
	flag.StringVar(&b.Token, "bot_token", token, "telegram bot token")
	flag.StringVar(&b.WebHookLink, "bot_web_hook_link", webHookLink, "telegram bot web hook link")
	flag.IntVar(&b.Port, "bot_port", botPort, "telegram bot port")
}
