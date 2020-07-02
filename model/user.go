package model

import "database/sql"

/*
	In this file, the model that we will often use will be described. Such models as User, World, Weapons, etc.
	These models will be used when writing, reading in the database, and also when receiving data from the client,
	in our case, from the bot.
*/

// User struct contains user data
type User struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Team      sql.NullString `json:"team"`
	Status    bool           `json:"role"`
	Health    float64        `json:"health"`
	Strength  float64        `json:"strength"`
	Defence   float64        `json:"defence"`
	Intellect float64        `json:"intellect"`
	Level     float64        `json:"level"`
	Currency  int            `json:"currency"`
	Inventory []int          `json:"inventory"`
}
