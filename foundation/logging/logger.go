package logging

import (
	"log/slog"
	"os"
)

// New creates a new logger with the given attributes.
func New(attrs ...slog.Attr) *slog.Logger {
	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		},
	)

	return slog.New(handler.WithAttrs(attrs))
}
