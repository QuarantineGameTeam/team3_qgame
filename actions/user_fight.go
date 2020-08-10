package actions

import (
	"errors"
	"fmt"
	"log"

	"github.com/team3_qgame/database/repository"
	"github.com/team3_qgame/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	helpFightMsg         = ""
	mult         float64 = 25

	fightStatus = "fight"
	mainStaus   = "main"

	succesAtt    = "Attack successful\nYou deal "
	unsuccsesAtt = "Unsuccessful attack\nYou deal "
	succesDef    = "Defence successful\nYou receive "
	unsuccsesDef = "Unsuccessful defence\nYou receive "
	counterAtt   = "Unsuccessful attack\nCounterattack! You receive "
	counterDef   = "Defence successful\nCounterattack! You deal "
)

type Fight struct {
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

func NewFight(userRepo *repository.UserRepository) *Fight {
	return &Fight{
		userRepo: userRepo,
	}
}

func (f *Fight) SetUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	f.bot = bot
	f.updates = updates
}

func (f *Fight) CHelp(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = helpMsg
	f.bot.Send(msg)
}

func (f *Fight) SwitchStatus(ChatID int64, status string) {
	userCheck, _ := f.userRepo.GetUserByID(ChatID)
	userCheck.Status = status
	_ = f.userRepo.UpdateUser(userCheck)
}

func (f *Fight) CheckStatus(ChatID int64, status string) error {
	userCheck, _ := f.userRepo.GetUserByID(ChatID)
	if userCheck.Status != status {
		err := errors.New("wrong status")
		return err
	}
	return nil
}

func (f *Fight) startFightKb(ChatID int64) {
	msg := tgbotapi.NewMessage(ChatID, "Are you sure")
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В БІЙ!!!", "Fight"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "Back"),
		),
	)

	msg.ReplyMarkup = &replyMarkup
	f.bot.Send(msg)
}

func (f *Fight) kbAttack(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Attack")
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Attack⚔️", "strength"),
			tgbotapi.NewInlineKeyboardButtonData("Attack💫", "intellect"),
		),
	)

	msg.ReplyMarkup = &replyMarkup
	f.bot.Send(msg)
}

func (f *Fight) kbDefence(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Defence")
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Defence🛡", "strength"),
			tgbotapi.NewInlineKeyboardButtonData("Defence🔮", "intellect"),
		),
	)

	msg.ReplyMarkup = &replyMarkup
	f.bot.Send(msg)
}

func (f *Fight) attackCallBack(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "")
	userCheck, _ := f.userRepo.GetUserByID(chatID)
	for update := range f.updates {
		if update.Message != nil {
			continue
		} else if update.CallbackQuery.Message.Chat.ID == chatID {
			switch update.CallbackQuery.Data {
			case "strength":
				f.attackersTurn = Turn{userCheck.Strength, 0}
				msg.Text = "Attack with bow 🏹"
			case "intellect":
				f.attackersTurn = Turn{0, userCheck.Intellect}
				msg.Text = "Attack with rainbow 🏳️‍🌈"
			}
			break
		} else {
			continue
		}
	}
	f.bot.Send(msg)
}

func (f *Fight) defenceCallBack(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "")
	userCheck, _ := f.userRepo.GetUserByID(chatID)
	for update := range f.updates {
		if update.Message != nil {
			continue
		} else if update.CallbackQuery.Message.Chat.ID == chatID {
			switch update.CallbackQuery.Data {
			case "strength":
				f.defendersTurn = Turn{userCheck.Defence, 0}
				msg.Text = "Use shield"
			case "intellect":
				f.defendersTurn = Turn{0, userCheck.Defence}
				msg.Text = "Use magic shield"
			}
			break
		} else {
			continue
		}
	}
	f.bot.Send(msg)
}

func (f *Fight) getEnemy() (bool, error) {
	var err error
	f.enemy, err = f.userRepo.GetRandomUser(f.user.ID)
	if err != nil {
		log.Println("GetRandomUser Err:", err)
		return false, err
	}
	msgEnemy := tgbotapi.NewMessage(f.enemy.ID, "")
	f.startFightKb(f.enemy.ID)
	for update := range f.updates {
		if update.Message != nil {
			continue
		} else if update.CallbackQuery.Message.Chat.ID == f.enemy.ID {
			switch update.CallbackQuery.Data {
			case "f_fight":
				findEnemy := true
				msgEnemy.Text = "Fight started"
				f.bot.Send(msgEnemy)
				return findEnemy, nil
			case "f_back":
				msgEnemy.Text = "Retreat"
				f.bot.Send(msgEnemy)
			}
			break
		} else {
			continue
		}
	}
	return false, nil
}

func (f *Fight) enemySearch(update tgbotapi.Update) {
	var (
		err     error
		msgUser = tgbotapi.NewMessage(update.Message.Chat.ID, "")
	)

	f.user, err = f.userRepo.GetUserByID(update.Message.Chat.ID)
	if err != nil {
		//msg.Text = "Internal server error"
		//log.Println("GetUserByID Err:", err)
		//f.bot.Send(msg)
		return
	}

	f.startFightKb(f.user.ID)

	for update := range f.updates {
		if update.Message != nil {
			continue
		} else if update.CallbackQuery.Message.Chat.ID == f.user.ID {
			switch update.CallbackQuery.Data {
			case "Fight":
				msgUser.Text = "Searching for the enemy ..."
				f.bot.Send(msgUser)
				findEnemy := false
				for findEnemy == false {
					findEnemy, _ = f.getEnemy()
				}
			case "Back":
				msgUser.Text = "Retreat"
				f.bot.Send(msgUser)
				return
			}
			break
		} else {
			continue
		}
	}

}

func (f *Fight) Fight(update tgbotapi.Update) {

	f.enemySearch(update)

	f.SwitchStatus(f.user.ID, fightStatus)
	f.SwitchStatus(f.enemy.ID, fightStatus)

	attacker := f.user
	defender := f.enemy

	var damage float64

	var strPoint float64
	var intPoint float64

	var msgAttaker tgbotapi.MessageConfig
	var msgDefender tgbotapi.MessageConfig

	// fight loop
	for true {

		msgAttaker = tgbotapi.NewMessage(attacker.ID, "")
		msgDefender = tgbotapi.NewMessage(defender.ID, "")

		f.kbAttack(attacker.ID)
		f.attackCallBack(attacker.ID)
		f.kbDefence(defender.ID)
		f.defenceCallBack(defender.ID)

		// the player move calculation
		strPoint = f.attackersTurn.param1 - f.defendersTurn.param1
		intPoint = f.attackersTurn.param2 - f.defendersTurn.param2

		// player move result
		// successful attacker STRENGTH hit
		if strPoint > 0 {
			damage = strPoint * mult
			defender.Health -= damage
			msgAttaker, msgDefender = fightMsg(attacker, defender, damage, succesAtt, unsuccsesDef)
			// Successful attacker INTELLECT hit
		} else if intPoint > 0 {
			damage = intPoint * mult
			defender.Health -= damage
			msgAttaker, msgDefender = fightMsg(attacker, defender, damage, succesAtt, unsuccsesDef)
			//	If Defence is bigger than attacker strength, than attacker get a hit back.
		} else if strPoint < 0 {
			damage = -strPoint * mult
			attacker.Health -= damage
			msgAttaker, msgDefender = fightMsg(attacker, defender, damage, counterAtt, counterDef)
			//	If Defence is bigger than attacker intellect, than attacker get a hit back.
		} else if intPoint < 0 {
			damage = -intPoint * mult
			attacker.Health -= damage
			msgAttaker, msgDefender = fightMsg(attacker, defender, damage, counterAtt, counterDef)
			// If defence equal attack
		} else {
			damage = 0
			msgAttaker, msgDefender = fightMsg(attacker, defender, damage, unsuccsesAtt, succesDef)
		}

		f.bot.Send(msgAttaker)
		f.bot.Send(msgDefender)

		// Check who win
		if attacker.Health < 1 || defender.Health < 1 {
			switch attacker.Health > defender.Health {
			case true:
				msgAttaker.Text = "You won"
				msgDefender.Text = "You lose"
			case false:
				msgDefender.Text = "You won"
				msgAttaker.Text = "You lose"

			}
			f.bot.Send(msgAttaker)
			f.bot.Send(msgDefender)

			break
		}

		// next player move
		attacker, defender = defender, attacker
	}

	f.SwitchStatus(f.user.ID, mainStaus)
	f.SwitchStatus(f.enemy.ID, mainStaus)
}

func fightMsg(
	attacker model.User,
	defender model.User,
	damage float64,
	msgAtt string,
	msgDef string) (
	msgAttacker tgbotapi.MessageConfig,
	msgDefender tgbotapi.MessageConfig,
) {

	msgAttacker = tgbotapi.NewMessage(attacker.ID, "")
	msgDefender = tgbotapi.NewMessage(defender.ID, "")

	msgAttacker.Text = fmt.Sprintf(msgAtt+"%v damage\n"+
		"\nYour  ❤️ : %v\n"+
		"Enemy 💙 : %v\n",
		damage, attacker.Health, defender.Health)
	msgDefender.Text = fmt.Sprintf(msgDef+"%v damage\n"+
		"\nYour  ❤️ : %v\n"+
		"Enemy 💙 : %v\n",
		damage, defender.Health, attacker.Health)
	return msgAttacker, msgDefender
}
