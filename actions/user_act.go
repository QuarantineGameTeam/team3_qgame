package actions

import (
	"fmt"
	"gihub.com/team3_qgame/database/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	helpMsg = "/register - bot register new user" +
		//		"\n/rename - change user name" +
		"\n/delete - delete user" +
		"\n/me - shows your use data" +
		"\n/allusers - get every bot users"
)

type User struct {
	userRepo *repository.UserRepository
	bot      *tgbotapi.BotAPI
	updates  tgbotapi.UpdatesChannel
}

func NewUser(userRepo *repository.UserRepository) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) SetUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	u.bot = bot
	u.updates = updates
}

func (u *User) CStart(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome tho the game! Chose registration")
	_, _ = u.bot.Send(msg)
}

func (u *User) CRegistration(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
	if userCheck.ID != update.Message.Chat.ID {
		msg.Text = "Enter your name"
		u.bot.Send(msg)
		for update := range u.updates {
			if update.Message.Text != "" {
				userName := update.Message.Text
				userCheck.ID = update.Message.Chat.ID
				userCheck.Name = userName
				_ = u.userRepo.NewUser(userCheck)
				msg.Text = "Welcome! Your username is " + userCheck.Name
				u.bot.Send(msg)
			}
		}
	} else {
		msg.Text = "Your user is already exists"
		u.bot.Send(msg)
	}
}
func (u *User) CDelete(update tgbotapi.Update) {
	_ = u.userRepo.DeleteUserByID(update.Message.Chat.ID)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Your user deleted"
	u.bot.Send(msg)
}

func (u *User) CGetUserInfo(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
	if userCheck.ID == update.Message.Chat.ID {
		msg.Text = "Your user info:" + fmt.Sprintf("\n%+v", userCheck)
		u.bot.Send(msg)
	} else {
		msg.Text = "You have no user yet"
		u.bot.Send(msg)
	}
}

func (u *User) CGetAllUsers(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	allUsers, _ := u.userRepo.GetAllUsers()
	for i, _ := range allUsers {
		msg.Text = fmt.Sprintf("%+v", allUsers[i])
		u.bot.Send(msg)
	}
}

func (u *User) CNameUpdate(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
	if userCheck.ID == update.Message.Chat.ID {
		msg.Text = "Enter your new name"
		u.bot.Send(msg)
		for update := range u.updates {
			if update.Message != nil {
				userName := update.Message.Text
				userCheck.Name = userName
				_ = u.userRepo.UpdateUser(userCheck)
				msg.Text = "Your new username is " + userCheck.Name
				u.bot.Send(msg)
			}
		}
	} else{
		msg.Text = "You have no user yet"
		u.bot.Send(msg)
	}
}

func (u *User) CHelp(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = helpMsg
	u.bot.Send(msg)
}
