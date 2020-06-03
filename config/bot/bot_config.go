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
	token       = "1221820046:AAES49vWZ9cV0Fq3Rxkq9Yc24do1ONh_PbM"
	webHookLink = "https://0e203ad8e341.ngrok.io"
	botPort     = 8443
)

func (b *BConfig) InitBotConfig() {
	flag.StringVar(&b.Token, "bot_token", token, "telegram bot token")
	flag.StringVar(&b.WebHookLink, "bot_web_hook_link", webHookLink, "telegram bot web hook link")
	flag.IntVar(&b.Port, "bot_port", botPort, "telegram bot port")
}
