package config

import "os"

type Mode int

const (
	ModeDev Mode = iota + 1
	ModeProd
)

type Config struct {
	Mode   Mode
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
	BEPort string
}

func New() Config {
	var mode Mode
	switch os.Getenv("MODE") {
	case "dev":
		mode = ModeDev
	case "prod":
		mode = ModeProd
	default:
		mode = ModeDev
	}

	return Config{
		Mode:   mode,
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		BEPort: os.Getenv("BE_PORT"),
	}
}
