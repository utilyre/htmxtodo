package main

import (
	"github.com/utilyre/htmxtodo/layers/configs"
	"github.com/utilyre/htmxtodo/layers/database"
	"github.com/utilyre/htmxtodo/layers/logger"
	"github.com/utilyre/htmxtodo/layers/templates"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			configs.NewMode,
			configs.NewDatabase,
			configs.NewServer,
			logger.New,
			database.New,
			templates.New,
		),
		fx.Invoke(),
	).Run()
}
