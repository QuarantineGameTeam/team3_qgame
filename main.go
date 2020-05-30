package main

import (
	"fmt"
	"gihub.com/team3_qgame/database/repository"
	"gihub.com/team3_qgame/model"
	"github.com/google/uuid"
	"log"

	"gihub.com/team3_qgame/config"
	"gihub.com/team3_qgame/database"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

/*
	In the main file we collect all function that needed for running our app.
*/

var newDBConfig *config.DBConfig

func main() {

	NewDBConnection := database.NewDBConnection(newDBConfig)
	err := NewDBConnection.Connect()
	if err != nil {
		log.Fatal("Connection to DB failed")
	}

	// Нижче код буде в майбутньому перенесений в папку з логікою гри
	// зараз він тут для нагядності
	conn := NewDBConnection.GetConnection()

	userRepo := repository.NewUserRepository(conn)


	/*GamerNoOne := model.User{
		ID:   uuid.New(),
		Name: "Alessandro",
	}

	GamerNoTwo := model.User{
		ID: uuid.New(),
		Name: "Jessica",
	}

	// Записуємо дані користувача "Alessandro" в базу данних
	_ = userRepo.NewUser(GamerNoOne)
	// Записуємо дані користувача "Jessica" в базу данних
	_ = userRepo.NewUser(GamerNoOne)

	allUsers, _ := userRepo.GetAllUsers()
	fmt.Println("All users id DB", allUsers)

	// користувач змінює свої данні
	GamerNoTwo.Name = "Rebeka"
	_ = userRepo.UpdateUser(GamerNoTwo)

	// отримати дані користувача за (UUID) унікальним ідентифікатором
	gamerNoOne, _ := userRepo.GetUserByID(GamerNoOne.ID)
	fmt.Println("This is user no one", gamerNoOne)
	gamerNoTwo, _ := userRepo.GetUserByID(GamerNoTwo.ID)
	fmt.Println("This is user no two", gamerNoTwo)

	// видалення користувача з бази данних
	//_ = userRepo.DeleteUserByID(GamerNoOne.ID)

	// Дивимось що у нас залишилось в базі після видалення
	fmt.Println("All users id DB", allUsers)*/
	bot, err := tgbotapi.NewBotAPI("Token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "register":
				msg.Text = "Enter your name"
				bot.Send(msg)
				for update := range updates {
					if update.Message == nil {
						continue
					} else {
						userName := update.Message.Text

						newUser := model.User{
							ID:   uuid.New(),
							Name: userName,
						}
						_ = userRepo.NewUser(newUser)
						break
					}
				}

			case "allusers":
				allUsers, _ := userRepo.GetAllUsers()
				fmt.Println("All users id DB", allUsers)
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}

		}

	}
}
