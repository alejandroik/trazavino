package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
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

	return &ZapLogger{logger: logger.Sugar()}
}

func (l *ZapLogger) Debugf(message string, args ...interface{}) {
	l.logger.Debugf(message, args...)
}

func (l *ZapLogger) Infof(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

func (l *ZapLogger) Infow(message string, args ...interface{}) {
	l.logger.Infow(message, args...)
}

func (l *ZapLogger) Warnf(message string, args ...interface{}) {
	l.logger.Warnf(message, args...)
}

func (l *ZapLogger) Warnw(message string, args ...interface{}) {
	l.logger.Warnw(message, args...)
}

func (l *ZapLogger) Errorf(message string, args ...interface{}) {
	l.logger.Errorf(message, args...)
}

func (l *ZapLogger) Errorw(message string, args ...interface{}) {
	l.logger.Errorw(message, args...)
}

func (l *ZapLogger) Fatalf(message string, args ...interface{}) {
	l.logger.Fatalf(message, args...)
}
