package logger

import "go.uber.org/zap"

type Interface interface {
	Debugf(message string, args ...interface{})
	Infof(message string, args ...interface{})
	Warnf(message string, args ...interface{})
	Errorf(message string, args ...interface{})
	Fatalf(message string, args ...interface{})
}

// Logger -.
type Logger struct {
	logger *zap.Logger
}

var _ Interface = (*Logger)(nil)

func New() *Logger {
	logger, _ := zap.NewProduction()

	return &Logger{logger: logger}
}

func (l *Logger) Debugf(message string, args ...interface{}) {
	l.logger.Sugar().Debugf(message, args...)
}

// Info -.
func (l *Logger) Infof(message string, args ...interface{}) {
	l.logger.Sugar().Infof(message, args...)
}

// Warn -.
func (l *Logger) Warnf(message string, args ...interface{}) {
	l.logger.Sugar().Warnf(message, args...)
}

// Error -.
func (l *Logger) Errorf(message string, args ...interface{}) {
	l.logger.Sugar().Errorf(message, args...)
}

// Fatal -.
func (l *Logger) Fatalf(message string, args ...interface{}) {
	l.logger.Sugar().Fatalf(message, args...)
}
