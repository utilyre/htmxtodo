package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/utilyre/htmxtodo/layers/config"
	"go.uber.org/fx"
)

func New(lc fx.Lifecycle, cfg config.Config, logger *slog.Logger) *sqlx.DB {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
		),
	)
	if err != nil {
		logger.Error("failed to open database connection", "err", err)
		os.Exit(1)
	}

	logger.Info("successfully opened database connection")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := db.PingContext(ctx); err != nil {
				return err
			}

			logger.Info("successfully pinged database")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := db.Close(); err != nil {
				return err
			}

			logger.Info("successfully closed database connection")
			return nil
		},
	})

	return db
}
