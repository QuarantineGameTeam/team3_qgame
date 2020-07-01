package database

import (
	"database/sql"
	"fmt"
	"github.com/team3_qgame/config/database"
	"log"

	_ "github.com/lib/pq"
)

/*
	config repository designed to store data related to presets,
	telling the part of the software that is closed from the user,
	how to proceed in the case specified by the rules.
*/

type DBConnection struct {
	dbConnection *sql.DB
	config       *database.DBConfig
}

func NewDBConnection(config *database.DBConfig) *DBConnection {
	return &DBConnection{
		config: config,
	}
}

func (d *DBConnection) GetConnection() (*sql.DB, error) {
	err := d.connect()
	if err != nil {
		return nil, err
	}

	return d.dbConnection, nil
}

func (d *DBConnection) connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.config.Host, d.config.Port, d.config.User, d.config.Password, d.config.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	d.dbConnection = db

	log.Println("Successfully connected!")
	return nil
}
