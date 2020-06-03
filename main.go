package main

import (
	"fmt"
	"log"
	"net/http"

	"gihub.com/team3_qgame/api"
	"gihub.com/team3_qgame/api/controller"
	"gihub.com/team3_qgame/config"
)

/*
	In the main file we collect all function that needed for running app.
*/

func main() {

	// Initiate program configuration
	var appConfig config.Config
	appConfig.InitConfig()


	bot := api.NewBot(&appConfig.BotConfig)

	botController := controller.NewBotController(bot)
	botController.StartWebHookListener()


	//fmt.Println(bot, dbConn)
	fmt.Println("CONFIG", appConfig)

	log.Println("PORT!!!!!!", appConfig.BotConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", appConfig.BotConfig.Port), nil))
}
