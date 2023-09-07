package router

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/utilyre/htmxtodo/layers/config"
	"go.uber.org/fx"
)

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) Validate(s any) error {
	if err := v.validate.Struct(s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return nil
}

func New(lc fx.Lifecycle, cfg config.Config) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	e.Validator = &Validator{validate: validator.New()}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				e.Logger.Fatal(e.Start(":" + cfg.BEPort))
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})

	return e
}
