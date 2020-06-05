package main

import (
	"log"

	"gihub.com/team3_qgame/actions"
	"gihub.com/team3_qgame/api"
	"gihub.com/team3_qgame/api/controller"
	"gihub.com/team3_qgame/api/updater"
	"gihub.com/team3_qgame/config"
	"gihub.com/team3_qgame/database"
	"gihub.com/team3_qgame/database/repository"
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

	// Create new update Manager
	updManager := updater.NewUpdateManager(userAct)

	//var useract user.UpdateManager
	//var uact user.UpdateManager
	botController.StartWebHookListener(updManager)
}
