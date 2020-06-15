package actions

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	"gihub.com/team3_qgame/database/repository"
	"gihub.com/team3_qgame/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	helpMsg = "\n/help - get all commands" +
		"\n/register - bot register new user" +
		"\n/rename - change user name" +
		"\n/delete - delete user" +
		"\n/me - shows your use data" +
		"\n/allusers - get every bot users" +
		"\n/changeteam - change or set your team" +
		"\n/rating - get my game rating" +
		"\n/startfight - lets start the fight "
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
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"Welcome tho the game! Chose /registration or use /help command for more information")
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
	userCheck, err := u.userRepo.GetUserByID(update.Message.Chat.ID)
	if err != nil {
		msg.Text = "Internal server error"
		log.Println("GetUserByID Err:", err)
		u.bot.Send(msg)
		return
	}
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

func (u *User) startFightKb(update tgbotapi.Update) {
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

func (u *User) kbAttack(chatID int64) {
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

func (u *User) kbDefence(chatID int64) {
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

func (u *User) startFight(update tgbotapi.Update) {
	var (
		err error

		msgUser = tgbotapi.NewMessage(update.Message.Chat.ID, "")
	)

	u.user, err = u.userRepo.GetUserByID(update.Message.Chat.ID)
	if err != nil {
		//msg.Text = "Internal server error"
		//log.Println("GetUserByID Err:", err)
		//u.bot.Send(msg)
		return
	}

	u.enemy, err = u.userRepo.GetRandomUser(u.user.ID)
	if err != nil {
		//msg.Text = "Internal server error"
		//log.Println("GetRandomUser Err:", err)
		//u.bot.Send(msg)
		return
	}

	msgEnemy := tgbotapi.NewMessage(u.enemy.ID, "")

	for update := range u.updates {
		switch update.CallbackQuery.Data {
		case "Fight":
			msgEnemy.Text = "Fight started"
			msgUser.Text = "Fight started"
			u.bot.Send(msgEnemy)
			u.bot.Send(msgUser)
		case "Back":
			msgEnemy.Text = "Retreat"
			u.bot.Send(msgEnemy)
		}
		break
	}
}

func (u *User) attackCallBack(chatID int64) {
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
}

func (u *User) Rating(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	user, err := u.userRepo.GetUserByID(update.Message.Chat.ID)
	if err != nil {
		msg.Text = "Internal server error"
		log.Println("GetUserByID Err:", err)
		u.bot.Send(msg)
		return
	}
	if user.ID == update.Message.Chat.ID {
		totalRating := user.Intellect + user.Defence + user.Strength
		msg.Text = "Your rating:" + fmt.Sprintf(
			"\nLevel %v"+
				"\n Defence %v"+
				"\n Health %v"+
				"\n Intellect %v"+
				"\n Strength %v"+
				"\n______________"+
				"\n Toatal: %f",
			user.Level, user.Defence, user.Health, user.Intellect, user.Strength, totalRating,
		)
		u.bot.Send(msg)
	} else {
		msg.Text = "You have no user yet"
		u.bot.Send(msg)
	}
}

func (u *User) defenceCallBack(chatID int64) {
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
}

func clear(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (u *User) Fight(update tgbotapi.Update) {
	u.startFightKb(update)
	u.startFight(update)

	attacker := u.user
	defender := u.enemy

	var damage float64

	var strPoint float64
	var intPoint float64

	// fight loop
	for true {

		msgAttaker := tgbotapi.NewMessage(attacker.ID, "")
		msgDeffender := tgbotapi.NewMessage(defender.ID, "")

		u.kbAttack(attacker.ID)
		u.attackCallBack(attacker.ID)
		u.kbDefence(defender.ID)
		u.defenceCallBack(defender.ID)

		// the player move calculation
		strPoint = u.attackersTurn.param1 - u.defendersTurn.param1
		intPoint = u.attackersTurn.param2 - u.defendersTurn.param2

		// player move result
		// successful attacker STRENGTH hit
		if strPoint > 0 {
			damage = strPoint * mult
			defender.Health -= damage
			msgAttaker.Text = fmt.Sprintf("You deal %v physical damage.\n"+
				"\n❤️ : %v\n"+
				"💙 : %v\n",
				damage, attacker.Health, defender.Health)
			msgDeffender.Text = fmt.Sprintf("You receive %v physical damage.\n"+
				"\n❤️ : %v\n"+
				"💙 : %v\n",
				damage, defender.Health, attacker.Health)
			// Successful attacker INTELLECT hit
		} else if intPoint > 0 {
			damage = intPoint * mult
			defender.Health -= damage
			msgAttaker.Text = fmt.Sprintf("You deal %v magic damage.\n"+
				"\n❤️ : %v\n"+
				"💙 : %v\n",
				damage, attacker.Health, defender.Health)
			msgDeffender.Text = fmt.Sprintf("You receive %v magic damage.\n"+
				"\n❤️ : %v\n"+
				"💙 : %v\n",
				damage, defender.Health, attacker.Health)
			//	If Defence is bigger than attacker strength, than attacker get a hit back.
		} else if strPoint < 0 {
			damage = -strPoint * mult
			attacker.Health -= damage
			msgAttaker.Text = fmt.Sprintf("Unsuccessful attack\n" +
				"Counterattack! You receive %v damage\n"+
				"\n❤️ : %v\n"+
				"💙 : %v\n",
				damage, attacker.Health, defender.Health)
			msgDeffender.Text = fmt.Sprintf("Defence successful\n" +
				"Counterattack! You deal %v damage\n"+
				"\n❤️ : %v\n"+
				"💙 : %v\n",
				damage, defender.Health, attacker.Health)
			//	If Defence is bigger than attacker intellect, than attacker get a hit back.
		} else if intPoint < 0 {
			damage = -intPoint * mult
			attacker.Health -= damage
			msgAttaker.Text = fmt.Sprintf("Unsuccessful attack\n" +
				"Counterattack! You receive %v damage\n"+
				"\n❤️ : %v\n"+
				"💙 : %v\n",
				damage, attacker.Health, defender.Health)
			msgDeffender.Text = fmt.Sprintf("Defence successful\n" +
				"Counterattack! You deal %v damage\n"+
				"\n❤️ : %v\n"+
				"💙 : %v\n",
				damage, defender.Health, attacker.Health)
			// If defence equal attack
		} else {
			msgAttaker.Text = fmt.Sprintf("Protection worked\n"+
				"\n❤️ : %v\n"+
				"💙 : %v\n",
				attacker.Health, defender.Health)
			msgDeffender.Text = fmt.Sprintf("Protection worked\n"+
				"\n❤️ : %v\n"+
				"💙 : %v\n",
				defender.Health, attacker.Health)
		}

		u.bot.Send(msgAttaker)
		u.bot.Send(msgDeffender)

		// Check who win
		if attacker.Health < 1 || defender.Health < 1 {
						switch attacker.Health > defender.Health {
			case true:
				msgAttaker.Text = "You won"
				msgDeffender.Text = "You lose"
			case false:
				msgDeffender.Text = "You won"
				msgAttaker.Text = "You lose"

			}
			u.bot.Send(msgAttaker)
			u.bot.Send(msgDeffender)
			//clear(u.user)
			//clear(u.enemy)
			break
		}

		// next player move
		attacker, defender = defender, attacker
	}
}

