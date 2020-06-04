package config

/*
	config repository designed to store data related to presets,
	telling the part of the software that is closed from the user,
	how to proceed in the case specified by the rules.
*/

import (
	"flag"
	"log"
)

type DBConfig struct {
	Host     string
	DBPort   int
	User     string
	Password string
	DBName   string
}

const (
	host     = "localhost"
	dbport   = 5432
	user     = "postgres"
	password = "team3bot"
	dbname   = "team3bot"
)

func (c *DBConfig) InitPgConfig() {
	flag.StringVar(&c.Host, "pg_host", host, "database discovery url")
	flag.IntVar(&c.DBPort, "pg_port", dbport, "database port")
	flag.StringVar(&c.User, "pg_user", user, "database user name")
	flag.StringVar(&c.Password, "pg_password", password, "database user password")
	flag.StringVar(&c.DBName, "pg_dbname", dbname, "name of database")
	log.Printf("app starts whith database configs:\n host=%s,\n port=%d,\n user_name=%s,\n db_name=%s;\n",
		c.Host, c.DBPort, c.User, c.DBName)
}
