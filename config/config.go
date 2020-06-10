package config

import (
	"flag"
	"gihub.com/team3_qgame/config/bot"
	"gihub.com/team3_qgame/config/database"
	"log"
)

type Config struct {
	BotConfig bot.BConfig
	DBConfig  database.DBConfig
}

func (c *Config) InitConfig() {
	c.BotConfig.InitBotConfig()
	c.DBConfig.InitPgConfig()
	flag.Parse()

	log.Println("app starts with follow config: . . .")
	log.Printf("database configs:\n host=%s ,\n port=%d ,\n user_name=%s ,\n db_name=%s ;\n",
		c.DBConfig.Host, c.DBConfig.Port, c.DBConfig.User, c.DBConfig.DBName)
	log.Printf("telegram bot configs:\n webhooklink=%s ,\n port=%d ;\n",
		c.BotConfig.WebHookLink, c.BotConfig.Port)
}
