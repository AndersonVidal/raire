package logging

import (
	"log/slog"
	"os"
)

// New returns a JSON logger. Logs are structured and must never contain PII;
// the access log intentionally omits query strings and request bodies.
func New(env string) *slog.Logger {
	level := slog.LevelInfo
	if env == "dev" {
		level = slog.LevelDebug
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
}
