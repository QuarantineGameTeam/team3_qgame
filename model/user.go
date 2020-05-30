package model

import "github.com/google/uuid"

/*
	In this file, the model that we will often use will be described. Such models as User, World, Weapons, etc.
	These models will be used when writing, reading in the database, and also when receiving data from the client,
	in our case, from the bot.
*/

// User struct contains user data
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Team	 	string    `json:"team"`
	Health	 	float32   `json:"health"`
	Strength	float32   `json:"strength"`
	Protection  float32   `json:"protection"`
	Intellect	float32   `json:"intellect"`
	Level	 	float32   `json:"levle"`
}

