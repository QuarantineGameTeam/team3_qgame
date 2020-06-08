package actions

import (
	//"log"
	"gihub.com/team3_qgame/database/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type User struct {
	userRepo *repository.UserRepository
	bot      *tgbotapi.BotAPI
	updates  tgbotapi.UpdatesChannel
}

var InformationKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("about me", "5"),
		tgbotapi.NewInlineKeyboardButtonData("about game", "1"),
		tgbotapi.NewInlineKeyboardButtonData("Rating", "33"),
	),
)

func NewUser(userRepo *repository.UserRepository) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) SetUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	u.bot = bot
	u.updates = updates
}

func (u *User) MSStart(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome tho the game! Chose registration")
	_, _ = u.bot.Send(msg)
}
func (u *User) MSRegistration(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	userCheck, _ := u.userRepo.GetUserByID(update.Message.Chat.ID)
	if userCheck.ID != update.Message.Chat.ID {
		msg.Text = "Enter your name"
		u.bot.Send(msg)
		for update := range u.updates {
			if update.Message == nil {
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
func (u *User) MSIformation(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "What information is needed?")
	msg.ReplyMarkup = InformationKeyboard
	_, _ = u.bot.Send(msg)
}

func (u *User) StartClanSelection(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Start clan selection for user ID")
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Clan 1", "CLAN_SELECT_1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Clan 2", "CLAN_SELECT_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Clan 3", "CLAN_SELECT_3"),
		),
	)
	msg.ReplyMarkup = replyMarkup
	msg.Text = "Please select a clan"
		u.bot.Send(msg)
}

/*func ProcessClanSelection(callbackQuery *telegram.CallbackQuery) {
	EditMessageReplyMarkup(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, nil)

	SaveUserClan(callbackQuery)

	user, err := GetUserFromDB(callbackQuery.From.ID) 
	if err != nil {
		log.Println("Could not get user", err)
	}

	SendMessage(callbackQuery.Message.Chat.ID, "Welcome to " + user.Clan + " clan)", nil)

	SendStartBattleMessage(callbackQuery)
}*/

func (u *User) method3() {}
