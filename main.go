package main

import (
	"github.com/utilyre/htmxtodo/layers/config"
	"github.com/utilyre/htmxtodo/layers/database"
	"github.com/utilyre/htmxtodo/layers/handlers"
	"github.com/utilyre/htmxtodo/layers/logger"
	"github.com/utilyre/htmxtodo/layers/router"
	"github.com/utilyre/htmxtodo/layers/storages"
	"github.com/utilyre/htmxtodo/layers/templates"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.New,
			logger.New,
			database.New,
			templates.New,
			router.New,

			storages.NewTodosStorage,
		),
		fx.Invoke(
			handlers.Public,
			handlers.Todos,
		),
	).Run()
}
