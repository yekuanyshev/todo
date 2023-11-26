package logger

import (
	"io"
	"log/slog"
)

func New(out io.Writer, level, format string) (*slog.Logger, error) {
	logLevel, err := parseLevel(level)
	if err != nil {
		return nil, err
	}

	logFormat, err := parseFormat(format)
	if err != nil {
		return nil, err
	}

	handlerOpts := &slog.HandlerOptions{
		Level: logLevel,
	}

	var handler slog.Handler

	switch logFormat {
	case FormatJSON:
		handler = slog.NewJSONHandler(out, handlerOpts)
	case FormatText:
		handler = slog.NewTextHandler(out, handlerOpts)
	}

	return slog.New(handler), nil
}
