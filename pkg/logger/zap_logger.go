package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func New() *Logger {
	cfg := zap.Config{
		Encoding:    "console",
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
		Level: zap.NewAtomicLevelAt(zapcore.InfoLevel),
	}

	logger, _ := cfg.Build()

	return &Logger{logger: logger.Sugar()}
}

func (l *Logger) Debugf(message string, args ...interface{}) {
	l.logger.Debugf(message, args...)
}

func (l *Logger) Infof(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

func (l *Logger) Infow(message string, args ...interface{}) {
	l.logger.Infow(message, args...)
}

func (l *Logger) Warnf(message string, args ...interface{}) {
	l.logger.Warnf(message, args...)
}

func (l *Logger) Warnw(message string, args ...interface{}) {
	l.logger.Warnw(message, args...)
}

func (l *Logger) Errorf(message string, args ...interface{}) {
	l.logger.Errorf(message, args...)
}

func (l *Logger) Errorw(message string, args ...interface{}) {
	l.logger.Errorw(message, args...)
}

func (l *Logger) Fatalf(message string, args ...interface{}) {
	l.logger.Fatalf(message, args...)
}
