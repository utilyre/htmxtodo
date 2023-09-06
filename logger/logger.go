package logger

import (
	"log/slog"
	"os"

	"github.com/utilyre/htmxtodo/config"
)

func New(c config.Config) *slog.Logger {
	var h slog.Handler
	switch c.Mode {
	case config.ModeDev:
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case config.ModeProd:
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	return slog.New(h)
}
