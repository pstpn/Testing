package logger

import (
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Interface interface {
	Debugf(message string, args ...interface{})
	Infof(message string, args ...interface{})
	Warnf(message string, args ...interface{})
	Errorf(message string, args ...interface{})
	Fatalf(message string, args ...interface{})
}

type Logger struct {
	logger *zerolog.Logger
}

func New(level string, w io.Writer) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger := zerolog.New(w).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		logger: &logger,
	}
}

func (l *Logger) Debugf(message string, args ...interface{}) {
	l.logger.Debug().Msgf(message, args...)
}

func (l *Logger) Infof(message string, args ...interface{}) {
	l.logger.Info().Msgf(message, args...)
}

func (l *Logger) Warnf(message string, args ...interface{}) {
	l.logger.Warn().Msgf(message, args...)
}

func (l *Logger) Errorf(message string, args ...interface{}) {
	l.logger.Error().Msgf(message, args...)
}

func (l *Logger) Fatalf(message string, args ...interface{}) {
	l.logger.Fatal().Msgf(message, args...)
	os.Exit(1)
}
