package service

import (
	"fmt"
	"gihub.com/team3_qgame/database/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

/*
	In package service must be describe business logic.
	That is, here we will already directly describe the process of the duel and interaction with other players.
*/

type UserData struct {
	userRepo *repository.UserRepository
	bot      *tgbotapi.BotAPI
}

func NewUserData(
	userRepo *repository.UserRepository,
	bot *tgbotapi.BotAPI,
) *UserData {
	newdata := UserData{
		userRepo,
		bot,
	}
	return &newdata
}


func (u *UserData) UserActions(updates tgbotapi.UpdatesChannel) {

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			helpMsg := "/register - bot register new user\n" +
				"/rename - change user name\n" +
				"/delete - delete user\n" +
				"/me - shows your name\n" +
				"/allusers - get every bot users"
			switch update.Message.Command() {
			case "register":
				userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
				if userCheck.ID != update.Message.Chat.ID {
					msg.Text = "Enter your name"
					u.bot.Send(msg)
					userName := update.Message.Text
					userCheck.ID = update.Message.Chat.ID
					userCheck.Name = userName
					_ = u.userRepo.NewUser(userCheck)
					msg.Text = "Welcome! Your username is " + userCheck.Name
					u.bot.Send(msg)
				} else {
					msg.Text = "Your user is already exists"
					u.bot.Send(msg)
				}

			case "rename":
				userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
				if userCheck.ID == update.Message.Chat.ID {
					msg.Text = "Enter your new name"
					u.bot.Send(msg)

					userName := update.Message.Text
					userCheck.Name = userName
					_ = u.userRepo.UpdateUser(userCheck)
					msg.Text = "Your new username is " + userCheck.Name
					u.bot.Send(msg)

				} else {
					msg.Text = "You have no user yet"
					u.bot.Send(msg)
				}
			case "delete":
				_ = u.userRepo.DeleteUserByID(update.Message.Chat.ID)
				msg.Text = "Your user deleted"
				u.bot.Send(msg)
			case "me":
				userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
				if userCheck.ID == update.Message.Chat.ID {
					msg.Text = "Your username is " + userCheck.Name + "Your user role is " + userCheck.Role.String
					u.bot.Send(msg)
				} else {
					msg.Text = "You have no user yet"
					u.bot.Send(msg)
				}
			case "allusers":
				allUsers, _ := u.userRepo.GetAllUsers()
				for i, _ := range allUsers {
					msg.Text = fmt.Sprintf("%+v", allUsers[i])
					u.bot.Send(msg)
				}
			case "help":
				msg.Text = helpMsg
				u.bot.Send(msg)
			default:
				msg.Text = "I don't know that command"
				u.bot.Send(msg)
				msg.Text = helpMsg
				u.bot.Send(msg)
			}

		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			u.bot.Send(msg)
		}

	}
}
