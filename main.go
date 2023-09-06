package main

import (
	"github.com/utilyre/htmxtodo/config"
	"github.com/utilyre/htmxtodo/database"
	"github.com/utilyre/htmxtodo/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			logger.New,
			config.New,
			database.New,
		),
		fx.Invoke(),
	).Run()
}
