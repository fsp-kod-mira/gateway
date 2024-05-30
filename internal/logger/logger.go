package logger

import (
	"gateway/internal/config"
	"log/slog"
	"os"
)

func New(config *config.Config) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: getLevel(config.LogLevel)}))
}

func getLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
