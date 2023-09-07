package configs

import "os"

type Mode int

const (
	ModeDev Mode = iota + 1
	ModeProd
)

func NewMode() Mode {
	switch os.Getenv("MODE") {
	case "dev":
		return ModeDev
	case "prod":
		return ModeProd
	default:
		return ModeDev
	}
}
