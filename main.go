package main

import (
	"flag"
	"gihub.com/team3_qgame/api"
	"gihub.com/team3_qgame/database/repository"

	"log"
	"net/http"

	"gihub.com/team3_qgame/config"
	"gihub.com/team3_qgame/database"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)
var newDBConfig config.DBConfig
var newBotConfig config.BotConfig

func main() {

	NewDBConnection := database.NewDBConnection(&newDBConfig)
	err := NewDBConnection.Connect()
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	conn := NewDBConnection.GetConnection()

	userRepo := repository.NewUserRepository(conn)

	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			helpMsg := "/register - bot register new user\n/rename - change user name\n/delete - delete user\n/me - shows your name\n/allusers - get every bot users"
			switch update.Message.Command() {
			case "register":
				userCheck, _ := userRepo.GetUserByID(update.Message.Chat.ID)
				if userCheck.ID != update.Message.Chat.ID {
					msg.Text = "Enter your name"
					bot.Send(msg)
					for update := range updates {
						if update.Message == nil {
							continue
						} else {
							userName := update.Message.Text
							userCheck.ID = update.Message.Chat.ID
							userCheck.Name = userName
							_ = userRepo.NewUser(userCheck)
							msg.Text = "Welcome! Your username is " + userCheck.Name
							bot.Send(msg)
							break
						}
					}
				} else {
					msg.Text = "Your user is already exists"
					bot.Send(msg)
				}

			case "rename":
				userCheck, _ := userRepo.GetUserByID(update.Message.Chat.ID)
				if userCheck.ID == update.Message.Chat.ID {
					msg.Text = "Enter your new name"
					bot.Send(msg)
					for update := range updates {
						if update.Message == nil {
							continue
						} else {
							userName := update.Message.Text
							userCheck.Name = userName
							_ = userRepo.UpdateUser(userCheck)
							msg.Text = "Your new username is " + userCheck.Name
							bot.Send(msg)
							break
						}
					}
				} else{
					msg.Text = "You have no user yet"
					bot.Send(msg)
				}
			case "delete":
				_ = userRepo.DeleteUserByID(update.Message.Chat.ID)
				msg.Text = "Your user deleted"
				bot.Send(msg)
			case "me":
				userCheck, _ := userRepo.GetUserByID(update.Message.Chat.ID)
				if userCheck.ID == update.Message.Chat.ID {
					msg.Text = "Your username is " + userCheck.Name
					bot.Send(msg)
					break
				} else {
					msg.Text = "You have no user yet"
					bot.Send(msg)
				}
			case "allusers":
				allUsers, _ := userRepo.GetAllUsers()
				for i, _ := range allUsers{
					msg.Text = fmt.Sprintf("%+v", allUsers[i])
					bot.Send(msg)
				}
			case "help":
				msg.Text = helpMsg
				bot.Send(msg)
			default:
				msg.Text = "I don't know that command"
				bot.Send(msg)
				msg.Text = helpMsg
				bot.Send(msg)
			}

		} else{
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			bot.Send(msg)
		}

	}

}
