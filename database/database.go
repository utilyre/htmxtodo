package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/utilyre/htmxtodo/config"
	"go.uber.org/fx"
)

func New(lc fx.Lifecycle, c config.Config, l *slog.Logger) *sqlx.DB {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName,
		),
	)
	if err != nil {
		l.Error("failed to open database connection", "err", err)
		os.Exit(1)
	}

	l.Info("successfully opened database connection")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := db.PingContext(ctx); err != nil {
				return err
			}

			l.Info("successfully pinged database")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := db.Close(); err != nil {
				return err
			}

			l.Info("successfully closed database connection")
			return nil
		},
	})

	return db
}
