package actions

import (
	"fmt"
	"gihub.com/team3_qgame/database/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

const (
	helpMsg = "/register - bot register new user" +
		//		"\n/rename - change user name" +
		"\n/delete - delete user" +
		"\n/me - shows your use data" +
		"\n/allusers - get every bot users"
)

const(
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
			if update.Message.Text == "" {
				continue
			} else {
				userName := update.Message.Text
				userCheck.ID = update.Message.Chat.ID
				userCheck.Name = userName
				_ = u.userRepo.NewUser(userCheck)
				msg.Text = "Welcome! Your username is " + userCheck.Name
				u.bot.Send(msg)
				break
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
	for i, _ := range allUsers{
		msg.Text = fmt.Sprintf("%+v", allUsers[i])
		u.bot.Send(msg)
	}
}

func (u *User) CHelp(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = helpMsg
	u.bot.Send(msg)
}


func (u *User) CGetAllUsers(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	allUsers, _ := u.userRepo.GetAllUsers()
	for i, _ := range allUsers {
		msg.Text = fmt.Sprintf("%+v", allUsers[i])
		u.bot.Send(msg)
	}
}

func (u *User) CUpdate(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
	var ParamKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Name", "Name"),
			tgbotapi.NewInlineKeyboardButtonData("Team", "Team"),
			tgbotapi.NewInlineKeyboardButtonData("Role", "Role"),
			tgbotapi.NewInlineKeyboardButtonData("Health", "Health"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Strength", "Strength"),
			tgbotapi.NewInlineKeyboardButtonData("Defence", "Defence"),
			tgbotapi.NewInlineKeyboardButtonData("Intellect", "Intellect"),
			tgbotapi.NewInlineKeyboardButtonData("Level", "Level"),
		),
	)
	msg.ReplyMarkup = ParamKeyboard
	u.bot.Send(msg)
	if userCheck.ID != update.Message.Chat.ID {
		for update := range u.updates {
			if update.CallbackQuery != nil {
				switch update.CallbackQuery.Data {
				case "Name":
					userName := update.Message.Text
					userCheck.Name = userName
					_ = u.userRepo.UpdateUser(userCheck)
				case "Team":
					userTeam := update.Message.Text
					userCheck.Team.String = userTeam
					_ = u.userRepo.UpdateUser(userCheck)
				case "Role":
					userRole := update.Message.Text
					userCheck.Role.String = userRole
					_ = u.userRepo.UpdateUser(userCheck)
				case "Health":
					userHealth := update.Message.Text
					userCheck.Health, _ = strconv.ParseFloat(userHealth, 32)
					_ = u.userRepo.UpdateUser(userCheck)
				case "Strength":
					userStrength := update.Message.Text
					userCheck.Strength, _ = strconv.ParseFloat(userStrength, 32)
					_ = u.userRepo.UpdateUser(userCheck)
				case "Defence":
					userDefence := update.Message.Text
					userCheck.Defence, _ = strconv.ParseFloat(userDefence, 32)
					_ = u.userRepo.UpdateUser(userCheck)
				case "Intellect":
					userIntellect := update.Message.Text
					userCheck.Intellect, _ = strconv.ParseFloat(userIntellect, 32)
					_ = u.userRepo.UpdateUser(userCheck)
				case "Level":
					userLevel := update.Message.Text
					userCheck.Level, _ = strconv.ParseFloat(userLevel, 32)
					_ = u.userRepo.UpdateUser(userCheck)
				}
			}
		}
	}
}

func (u *User) CHelp(update tgbotapi.Update) {
msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
msg.Text = helpMsg
u.bot.Send(msg)
}

