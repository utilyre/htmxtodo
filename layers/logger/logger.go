package logger

import (
	"log/slog"
	"os"

	"github.com/utilyre/htmxtodo/layers/configs"
)

func New(mode configs.Mode) *slog.Logger {
	var h slog.Handler
	switch mode {
	case configs.ModeDev:
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case configs.ModeProd:
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	return slog.New(h)
}
