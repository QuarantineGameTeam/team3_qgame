package database

import (
	"database/sql"
	"fmt"
	"log"

	"gihub.com/team3_qgame/config"
	_ "github.com/lib/pq"
)

/*
	This is where the client’s database connection code is implemented.
	Typically, the server side is implemented such that a new thread processes requests for the connection.
	A connection pool is implemented here to minimize outlet opening.
	Usually, you are not indifferent through the connection (if everyone is connected as the same user)
	that you received the database result set. You do not want to consume resources, so you want to be pleasant,
	and when you are done, you close the connection. I believe that every server ends the connection today if there
	is no activity for some time (timeout), that is, working with the database.
*/

type DBConnection struct {
	dbConnection *sql.DB
	config       *config.DBConfig
}

func NewDBConnection(config *config.DBConfig) *DBConnection {
	config.InitPgConfig()
	//fmt.Println(config.InitPgConfig)
	return &DBConnection{
		config: config,
	}

}

func (d *DBConnection) GetConnection() *sql.DB {

	return d.dbConnection

}

func (d *DBConnection) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.config.Host, d.config.Port, d.config.User, d.config.Password, d.config.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {

		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		
		return err
	}

	d.dbConnection = db

	log.Println("Successfully connected!")
	return nil
}
