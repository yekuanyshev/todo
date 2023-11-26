package logger

import (
	"fmt"
	"log/slog"
)

type Level string

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARN"
	LevelError Level = "ERROR"
)

func parseLevel(level string) (slog.Level, error) {
	switch Level(level) {
	case LevelDebug:
		return slog.LevelDebug, nil
	case LevelInfo:
		return slog.LevelInfo, nil
	case LevelWarn:
		return slog.LevelWarn, nil
	case LevelError:
		return slog.LevelError, nil
	}

	return 0, fmt.Errorf("invalid log level: %s", level)
}
