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
	noTeamString string  = "noteam"
	mult         float64 = 25
)

type User struct {
	userRepo      *repository.UserRepository
	bot           *tgbotapi.BotAPI
	updates       tgbotapi.UpdatesChannel
	user          model.User
	enemy         model.User
	attackersTurn Turn
	defendersTurn Turn
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

func (u *User) KbAttack(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Attack")
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Attack⚔️", "strength"),
			tgbotapi.NewInlineKeyboardButtonData("Attack💫", "intellect"),
		),
	)

	msg.ReplyMarkup = &replyMarkup
	u.bot.Send(msg)
}

func (u *User) KbDefence(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Defence")
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
	msg4u := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	u.enemy, _ = u.userRepo.GetRandomUser()
	u.user, _ = u.userRepo.GetUserByID(update.Message.Chat.ID)
	msg := tgbotapi.NewMessage(u.enemy.ID, "")
	for update := range u.updates {
		switch update.CallbackQuery.Data {
		case "Fight":
			msg.Text = "Fight started"
			msg4u.Text = "Fight started"
			u.bot.Send(msg)
			u.bot.Send(msg4u)
		case "Back":
			msg.Text = "Retreat"
			u.bot.Send(msg)
		}
		break
	}
}

func (u *User) AttackCallBack(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "")
	userCheck, _ := u.userRepo.GetUserByID(chatID)
	for update := range u.updates {
		switch update.CallbackQuery.Data {
		case "strength":
			u.attackersTurn = Turn{userCheck.Strength, 0}
			msg.Text = "Attack with bow 🏹"
		case "intellect":
			u.attackersTurn = Turn{0, userCheck.Intellect}
			msg.Text = "Attack with rainbow 🏳️‍🌈"
		}
		break
	}
	u.bot.Send(msg)
	msg.Text = fmt.Sprintf("\n%+v", u.attackersTurn)
	u.bot.Send(msg)
}

func (u *User) DefenceCallBack(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "")
	userCheck, _ := u.userRepo.GetUserByID(chatID)
	for update := range u.updates {
		switch update.CallbackQuery.Data {
		case "strength":
			u.defendersTurn = Turn{userCheck.Defence, 0}
			msg.Text = "Use Shield"
		case "intellect":
			u.defendersTurn = Turn{0, userCheck.Defence}
			msg.Text = "Become invisible"
		}
		break
	}
	u.bot.Send(msg)
	msg.Text = fmt.Sprintf("\n%+v", u.defendersTurn)
	u.bot.Send(msg)
}

func (u *User) Fight(update tgbotapi.Update) {
	u.CStartFightKb(update)
	u.StartFight(update)
	var i = 0
	attacker := u.user
	defender := u.enemy
	msgA := tgbotapi.NewMessage(attacker.ID, "")
	msgD := tgbotapi.NewMessage(defender.ID, "")
	for true {
		var strPoint float64
		var intPoint float64
		u.KbAttack(attacker.ID)
		u.AttackCallBack(attacker.ID)
		u.KbDefence(defender.ID)
		u.DefenceCallBack(defender.ID)
		strPoint = u.attackersTurn.param1 - u.defendersTurn.param1
		intPoint = u.attackersTurn.param2 - u.defendersTurn.param2
		if strPoint > 0 {
			defender.Health -= strPoint * mult
			msgA.Text = fmt.Sprintf("your health - %v\n"+
				"enemy health - %v\n", attacker.Health, defender.Health)
			msgD.Text = fmt.Sprintf("your health - %v\n"+
				"enemy health - %v\n", defender.Health, attacker.Health)
		} else if intPoint > 0 {
			defender.Health -= intPoint * mult
			msgA.Text = fmt.Sprintf("your health - %v\n"+
				"enemy health - %v\n", attacker.Health, defender.Health)
			msgD.Text = fmt.Sprintf("your health - %v\n"+
				"enemy health - %v\n", defender.Health, attacker.Health)
		} else if u.attackersTurn.param1 > 0 && u.defendersTurn.param1 > u.attackersTurn.param1 {
			attacker.Health -= (u.defendersTurn.param1 - u.attackersTurn.param1) * mult
			msgA.Text = fmt.Sprintf("counterattack\n"+
				"your health - %v\n"+
				"enemy health - %v\n", attacker.Health, defender.Health)
			msgD.Text = fmt.Sprintf("counterattack\n"+
				"your health - %v\n"+
				"enemy health - %v\n", defender.Health, attacker.Health)
		} else if u.attackersTurn.param2 > 0 && u.defendersTurn.param2 > u.attackersTurn.param2 {
			attacker.Health -= (u.defendersTurn.param2 - u.attackersTurn.param2) * mult
			msgA.Text = fmt.Sprintf("counterattack\n"+
				"your health - %v\n"+
				"enemy health - %v\n", attacker.Health, defender.Health)
			msgD.Text = fmt.Sprintf("counterattack\n"+
				"your health - %v\n"+
				"enemy health - %v\n", defender.Health, attacker.Health)
		} else {
			msgA.Text = fmt.Sprintf("protection worked\n"+
				"your health - %v\n"+
				"enemy health - %v\n", defender.Health, attacker.Health)
			msgD.Text = fmt.Sprintf("protection worked\n"+
				"your health - %v\n"+
				"enemy health - %v\n", attacker.Health, defender.Health)
		}
		u.bot.Send(msgA)
		u.bot.Send(msgD)
		if attacker.Health < 1 || defender.Health < 1{
			switch attacker.Health > defender.Health {
			case true:
				msgA.Text = "You won"
				msgD.Text = "You lose"
			case false:
				msgD.Text = "You won"
				msgA.Text = "You lose"

			}
			u.bot.Send(msgA)
			u.bot.Send(msgD)
			break
		}
		i++
		attacker, defender = defender, attacker
	}
}
