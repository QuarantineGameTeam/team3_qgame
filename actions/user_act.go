package actions

import (
	"database/sql"
	"fmt"
	"gihub.com/team3_qgame/database/repository"
	"gihub.com/team3_qgame/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	helpMsg = "/register - bot register new user" +
		"\n/rename - change user name" +
		"\n/delete - delete user" +
		"\n/me - shows your use data" +
		"\n/allusers - get every bot users" +
		"\n/changeteam - change or set your team"
	noTeamString string = "noteam"
)

type User struct {
	userRepo *repository.UserRepository
	bot      *tgbotapi.BotAPI
	updates  tgbotapi.UpdatesChannel
	enemy    *model.User
}

type Turn struct {
	param1 float64
	param2 float64
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
				break
			}
		}
	} else {
		msg.Text = "You have no user yet"
		u.bot.Send(msg)
	}
}

func (u *User) CHelp(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = helpMsg
	u.bot.Send(msg)
}
func (u *User) CStartTeamSelection(update tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Chose your team")
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "TEAM_1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "TEAM_2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "TEAM_3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("No team", noTeamString),
		),
	)

	msg.ReplyMarkup = &replyMarkup
	u.bot.Send(msg)
}

func (u *User) CStartFightKb(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Are you sure")
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В БІЙ!!!", "Fight"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "Back"),
		),
	)

	msg.ReplyMarkup = &replyMarkup
	u.bot.Send(msg)
}

func (u *User) KbAttack(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Attack")
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Attack⚔️", "strength"),
			tgbotapi.NewInlineKeyboardButtonData("Attack💫", "intellect"),
		),
	)

	msg.ReplyMarkup = &replyMarkup
	u.bot.Send(msg)
}

func (u *User) KbDefence(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(u.enemy.ID, "Defence")
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Defence🛡", "strength"),
			tgbotapi.NewInlineKeyboardButtonData("Defence🔮", "intellect"),
		),
	)

	msg.ReplyMarkup = &replyMarkup
	u.bot.Send(msg)
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NullString() sql.NullString {
	return sql.NullString{
		String: "",
		Valid:  false,
	}
}

func (u *User) TeamChange(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
	var teamName string
	for update := range u.updates {
		if update.CallbackQuery.Data != "" && update.CallbackQuery.Data != noTeamString {
			teamName = update.CallbackQuery.Data
			userCheck.Team = NewNullString(teamName)
			_ = u.userRepo.UpdateUser(userCheck)
			msg.Text = "Your team is " + userCheck.Team.String
			break
		} else {
			userCheck.Team = NullString()
			_ = u.userRepo.UpdateUser(userCheck)
			msg.Text = "You are not in the team"
			break
		}
	}
	u.bot.Send(msg)
}

func (u *User) StartFight(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	for update := range u.updates {
		switch update.CallbackQuery.Data {
		case "Fight":
			msg.Text = "Fight started"
			u.bot.Send(msg)
		case "Back":
			msg.Text = "Retreat"
			u.bot.Send(msg)
		}
		break
	}
}

func (u *User) AttackCallBack(update tgbotapi.Update) Turn {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
	var attackerTurn Turn
	for update := range u.updates {
		switch update.CallbackQuery.Data {
		case "strength":
			attackerTurn = Turn{userCheck.Strength, 0}
			msg.Text = "Attack with bow 🏹"
		case "intellect":
			attackerTurn = Turn{0, userCheck.Intellect}
			msg.Text = "Attack with rainbow 🏳️‍🌈"
		}
		break
	}
	u.bot.Send(msg)
	return attackerTurn
}

/*func (u *User) DefenceCallBack(update tgbotapi.Update) Turn {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
	var defenderTurn Turn
	for update := range u.updates {
		switch update.CallbackQuery.Data {
		case "strength":
			defenderTurn = Turn {userCheck.Defence, 0}
			msg.Text = "Use Shield"
			break
		case "intellect":
			defenderTurn = Turn {0, userCheck.Defence}
			msg.Text = "Become invisible"
			break
		}
		u.bot.Send(msg)
	}
	return defenderTurn
}*/
