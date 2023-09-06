package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Mode int

const (
	ModeDev Mode = iota + 1
	ModeProd
)

type Config struct {
	Mode Mode

	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string

	BEHost string
	BEPort string
}

func New() (Config, error) {
	var mode Mode
	switch os.Getenv("MODE") {
	case "":
		mode = ModeDev
	case "dev":
		mode = ModeDev
	case "prod":
		mode = ModeProd
	default:
		return Config{}, errors.New("invalid mode")
	}

	if mode == ModeDev {
		if err := godotenv.Load(); err != nil {
			return Config{}, err
		}
	}

	if err := assertValidEnv([]string{
		"DB_USER",
		"DB_PASS",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"BE_HOST",
		"BE_PORT",
	}); err != nil {
		return Config{}, err
	}

	return Config{
		Mode: mode,

		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),

		BEHost: os.Getenv("BE_HOST"),
		BEPort: os.Getenv("BE_PORT"),
	}, nil
}

func assertValidEnv(keys []string) error {
	for _, key := range keys {
		if _, ok := os.LookupEnv(key); !ok {
			return fmt.Errorf("variable '%s' not found in the environment", key)
		}
	}

	return nil
}
