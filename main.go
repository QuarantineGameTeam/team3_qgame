package main

import (
	"log"

	"github.com/team3_qgame/actions"
	"github.com/team3_qgame/api"
	"github.com/team3_qgame/api/controller"
	"github.com/team3_qgame/api/updater"
	"github.com/team3_qgame/config"
	"github.com/team3_qgame/database"
	"github.com/team3_qgame/database/repository"
)

/*
	In the main file we collect all function that needed for running app.
*/

func main() {

	// Initiate program configuration
	var appConfig config.Config
	appConfig.InitConfig()

	// Initiate connection to database
	dbConn := database.NewDBConnection(&appConfig.DBConfig)
	conn, err := dbConn.GetConnection()
	if err != nil {
		log.Println("DB connection failure, error:", err.Error())
	}

	// Create new instance of user repository
	userRepo := repository.NewUserRepository(conn)

	// Initiate new bot connection
	bot := api.NewBot(&appConfig.BotConfig)
	botController := controller.NewBotController(bot)

	// Initiate new user action instance
	userAct := actions.NewUser(userRepo)

	// Initiate new user action instance
	fighting := actions.NewFight(userRepo)

	// Create new user update
	updUser := updater.NewUserUpdate(userAct)

	//
	updFight := updater.NewFightUpdater(fighting)

	//var useract user.UserUpdater
	//var uact user.UserUpdater
	botController.StartWebHookListener(updUser, updFight)
}
