package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/phsym/console-slog"
)

var (
	loggerContextKey = struct{}{}

	noopLogger = slog.New(slog.NewTextHandler(io.Discard, nil))
)

// NewContext returns a new context with the logger in it.
func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

// FromContext gets the current logger from the context. If there is none
// it will return a noop logger.
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerContextKey).(*slog.Logger); ok {
		return logger
	}

	return noopLogger
}

// NewConsoleLogger returns a logger for the console.
func NewConsoleLogger(level slog.Level) *slog.Logger {
	_, disableColor := os.LookupEnv("NO_COLOR")
	return slog.New(console.NewHandler(
		os.Stderr,
		&console.HandlerOptions{
			Level:      level,
			NoColor:    disableColor,
			TimeFormat: time.TimeOnly,
		},
	))
}
