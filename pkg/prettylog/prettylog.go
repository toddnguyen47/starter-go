package prettylog

import (
	"time"

	"github.com/phuslu/log"
)

const (
	TimeFormatISO8601 = "2006-01-02T15:04:05Z"
)

func GetConsoleLogger() log.Logger {
	logger := log.Logger{
		Level:      log.InfoLevel,
		TimeField:  "ts",
		TimeFormat: time.RFC3339,
		Writer: &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
		},
	}

	return logger
}
