package templates

import (
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
)

func New(logger *slog.Logger) *template.Template {
	tmpl, err := template.ParseGlob(filepath.Join("views", "*.html"))
	if err != nil {
		logger.Error("failed to parse html templates", "err", err)
		os.Exit(1)
	}

	return tmpl
}
