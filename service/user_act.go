package service

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

/*
	In package service must be describe business logic.
	That is, here we will already directly describe the process of the duel and interaction with other players.
*/

type UserData struct {
	userPrepo *repository.UserRepository
	userPrepo *UserRepository
	bot       *tgbotapi.BotAPI
	update    tgbotapi.Update
	updates   tgbotapi.UpdatesChannel
	msg       tgbotapi.MessageConfig
}

func NewUseData(userPrepo *repository.UserRepository, bot *tgbotapi.BotAPI,	update tgbotapi.Update,
func NewUseData(userPrepo *UserRepository, bot *tgbotapi.BotAPI,	update tgbotapi.Update,
	updates tgbotapi.UpdatesChannel,
	msg tgbotapi.MessageConfig,
) *UserData {
	newdata := UserData{
		userPrepo,
		bot,
		update,
		updates,
		msg,
	}
	return &newdata
}

const helpMsg = "/register - bot register new user\n/rename - change user name\n/delete - delete user\n/me - shows your name\n/allusers - get every bot users"

func (u *UserData) UserInteraction() {

	switch u.update.Message.Command() {
	case "register":
		u.Register(u.bot, u.update, u.msg)
	case "rename":
		userCheck, _ := u.userPrepo.GetUserByID(u.update.Message.Chat.ID)
		if userCheck.ID == u.update.Message.Chat.ID {
			u.msg.Text = "Enter your new name"
			u.bot.Send(u.msg)
			for update := range u.updates {
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
		} else {
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
		for i, _ := range allUsers {
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
}

func (u *UserData) Register() {
	userCheck, _ := u.userPrepo.GetUserByID(u.update.Message.Chat.ID)
	if userCheck.ID != u.update.Message.Chat.ID {
		u.msg.Text = "Enter your name"
		u.bot.Send(u.msg)
		for update := range u.updates {
			if update.Message == nil {
				continue
			} else {
				userName := update.Message.Text
				userCheck.ID = update.Message.Chat.ID
				userCheck.Name = userName
				_ = u.userPrepo.NewUser(userCheck)
				u.msg.Text = "Welcome! Your username is " + userCheck.Name
				u.bot.Send(u.msg)
				break
			}
		}
	} else {
		u.msg.Text = "Your user is already exists"
		u.bot.Send(u.msg)
	}
}