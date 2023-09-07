package configs

import "os"

type Database struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func NewDatabase() Database {
	return Database{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}
}
