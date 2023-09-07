package logger

import (
	"log/slog"
	"os"

	"github.com/utilyre/htmxtodo/layers/config"
)

func New(cfg config.Config) *slog.Logger {
	var h slog.Handler
	switch cfg.Mode {
	case config.ModeDev:
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case config.ModeProd:
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	return slog.New(h)
}
