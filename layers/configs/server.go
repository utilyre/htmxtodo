package configs

import "os"

type Server struct {
	Host string
	Port string
}

func NewServer() Server {
	return Server{
		Host: os.Getenv("BE_HOST"),
		Port: os.Getenv("BE_PORT"),
	}
}
