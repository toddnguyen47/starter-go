package prettylog

import (
	"github.com/phuslu/log"
)

const (
	TimeFormatISO8601 = "2006-01-02T15:04:05Z"
)

// GetConsoleAndFileLogger - Ref: https://github.com/phuslu/log
func GetConsoleAndFileLogger() log.Logger {
	logger := log.Logger{
		Level:      log.InfoLevel,
		TimeField:  "ts",
		TimeFormat: TimeFormatISO8601,
		Writer: &log.MultiEntryWriter{
			&log.ConsoleWriter{
				ColorOutput:    true,
				QuoteString:    true,
				EndWithMessage: true,
			},
			&log.FileWriter{
				Filename:     "logs/main.log",
				FileMode:     0600,
				MaxSize:      100 * 1024 * 1024,
				MaxBackups:   7,
				EnsureFolder: true,
				LocalTime:    true,
			},
		},
	}

	return logger
}

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
