package database

/*
	config repository designed to store data related to presets,
	telling the part of the software that is closed from the user,
	how to proceed in the case specified by the rules.
*/

import (
	"flag"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

const (
	host     = "localhost"
	dbPort   = 5432
	user     = "postgres"
	password = "03042013"
	dbname   = "team3bot"
)

func (c *DBConfig) InitPgConfig() {
	flag.StringVar(&c.Host, "pg_host", host, "database discovery url")
	flag.IntVar(&c.Port, "pg_port", dbPort, "database port")
	flag.StringVar(&c.User, "pg_user", user, "database user name")
	flag.StringVar(&c.Password, "pg_password", password, "database user password")
	flag.StringVar(&c.DBName, "pg_dbname", dbname, "name of database")
}
