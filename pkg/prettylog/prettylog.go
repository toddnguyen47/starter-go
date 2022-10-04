package prettylog

import (
	"github.com/phuslu/log"
)

const (
	TimeFormatISO8601 = "2006-01-02T15:04:05Z"
)

// GetConsoleLogger - Ref: https://github.com/phuslu/log
func GetConsoleLogger() log.Logger {
	logger := log.Logger{
		Level:      log.InfoLevel,
		TimeField:  "ts",
		TimeFormat: TimeFormatISO8601,
		Writer: &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
		},
	}

	return logger
}
