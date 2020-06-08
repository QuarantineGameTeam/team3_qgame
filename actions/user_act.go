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

/*type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageID string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
}*/
/*type EditMessageReplyMarkup struct {
	ChatID      int64                 `json:"chat_id"`
	MessageID   int64                 `json:"message_id"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
func (u *User) ProcessClanSelection(update tgbotapi.Update) {
	EditMessageReplyMarkup(update.Message.Chat.ID, update.Message.MessageID, nil)

	SaveUserClan(update)

	user, err := GetUserFromDB(update.From.ID) 
	if err != nil {
		log.Println("Could not get user", err)
	}

	SendMessage(update.Message.Chat.ID, "Welcome to " + user.team + " team)", nil)

	SendStartBattleMessage(update)
}*/

func (u *User) method3() {}
