package main

import (
	"fmt"
	"gihub.com/team3_qgame/database/repository"
	"gihub.com/team3_qgame/model"
	"github.com/google/uuid"
	"log"

	"gihub.com/team3_qgame/config"
	"gihub.com/team3_qgame/database"
)

/*
	In the main file we collect all function that needed for running our app.
*/

var newDBConfig config.DBConfig

func main() {

	NewDBConnection := database.NewDBConnection(&newDBConfig)
	err := NewDBConnection.Connect()
	if err != nil {
		log.Fatal("Connection to DB failed")
	}

	// Нижче код буде в майбутньому перенесений в папку з логікою гри
	// зараз він тут для нагядності
	conn := NewDBConnection.GetConnection()

	userRepo := repository.NewUserRepository(conn)

	//
	GamerNoOne := model.User{
		ID:   uuid.New(),
		Name: "Alessandro",
	}

	GamerNoTwo := model.User{
		ID:   uuid.New(),
		Name: "Jessica",
	}

	// Записуємо дані користувача "Alessandro" в базу данних
	_ = userRepo.NewUser(GamerNoOne)
	// Записуємо дані користувача "Jessica" в базу данних
	_ = userRepo.NewUser(GamerNoTwo)

	allUsers, _ := userRepo.GetAllUsers()
	fmt.Println("All users id DB", allUsers)

	// користувач змінює свої данні
	GamerNoTwo.Name = "Rebeka"
	_ = userRepo.UpdateUser(GamerNoTwo)

	// отримати дані користувача за (UUID) унікальним ідентифікатором
	gamerNoTwo, _ := userRepo.GetUserByID(GamerNoTwo.ID)
	fmt.Println("This is user no two", gamerNoTwo)

	// видалення користувача з бази данних
	_ = userRepo.DeleteUserByID(GamerNoTwo.ID)

	// Дивимось що у нас залишилось в базі після видалення
	fmt.Println("All users id DB", allUsers)
}
