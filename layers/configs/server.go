package configs

import "os"

type Server struct {
	Port string
}

func NewServer() Server {
	return Server{
		Port: os.Getenv("BE_PORT"),
	}
}
