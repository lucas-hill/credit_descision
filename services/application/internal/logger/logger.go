package logger

import (
	"log/slog"
	"os"
)

func New(level string, environment string) *slog.Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{Level: logLevel}

	if environment == "local" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		// JSON logs for dev/prod — structured for CloudWatch/Datadog
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
