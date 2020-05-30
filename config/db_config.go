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
	Port     int
	User     string
	Password string
	DBName   string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "teem3bot"
	dbname   = "team3bot"
)

func (c *DBConfig) InitPgConfig() {
	flag.StringVar(&c.Host, "pg_host", host, "database discovery url")
	flag.IntVar(&c.Port, "pg_port", port, "database port")
	flag.StringVar(&c.User, "pg_user", user, "database user name")
	flag.StringVar(&c.Password, "pg_user", password, "database user password")
	flag.StringVar(&c.DBName, "pg_user", dbname, "name of database")
	flag.Parse()
	log.Printf("app starts whith database configs:\n host=%s ,\n port=%d ,\n user_name=%s ,\n db_name=%s ;\n",
		c.Host, c.Port, c.User, c.DBName)
}
