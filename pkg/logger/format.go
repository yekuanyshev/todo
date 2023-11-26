package logger

import (
	"fmt"
	"slices"
)

type Format string

const (
	FormatJSON = "json"
	FormatText = "text"
)

func parseFormat(format string) (Format, error) {
	formats := []string{
		FormatJSON,
		FormatText,
	}

	if !slices.Contains(formats, format) {
		return "", fmt.Errorf("invalid format: %s", format)
	}

	return Format(format), nil
}
