package mslog

import (
	"errors"
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	config *Config
	logger *logrus.Logger
	sentry *sentry.Hub
}

func NewLogger(config *Config, lr *logrus.Logger, sentry *sentry.Hub) *Logger {
	logger := &Logger{
		config: config,
		logger: lr,
		sentry: sentry,
	}

	return logger
}

func (logger *Logger) Info(args ...interface{}) {
	logger.logger.Info(args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.logger.Debug(args...)
}

func (logger *Logger) Trace(args ...interface{}) {
	logger.logger.Trace(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.SentryWarn(args...)
	logger.logger.Warn(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.SentryException(errors.New(fmt.Sprint(args...)))
	logger.logger.Error(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.SentryException(errors.New(fmt.Sprint(args...)))
	logger.logger.Fatal(args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.SentryException(errors.New(fmt.Sprint(args...)))
	logger.logger.Panic(args...)
}

func (logger *Logger) Print(v ...interface{}) {
	logger.logger.Print(v...)
}

func (logger *Logger) SentryException(exception error) {
	if logger.sentry != nil {
		logger.sentry.CaptureException(exception)
	}
}

func (logger *Logger) SentryWarn(args ...interface{}) {
	if logger.sentry != nil {
		logger.sentry.CaptureMessage(fmt.Sprint(args...))
	}
}

func (logger *Logger) Logrus() *logrus.Logger {
	return logger.logger
}
