package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() *zapLogger {
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

	return &zapLogger{logger: logger.Sugar()}
}

func (l *zapLogger) Debugf(message string, args ...interface{}) {
	l.logger.Debugf(message, args...)
}

func (l *zapLogger) Infof(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

func (l *zapLogger) Infow(message string, args ...interface{}) {
	l.logger.Infow(message, args...)
}

func (l *zapLogger) Warnf(message string, args ...interface{}) {
	l.logger.Warnf(message, args...)
}

func (l *zapLogger) Warnw(message string, args ...interface{}) {
	l.logger.Warnw(message, args...)
}

func (l *zapLogger) Errorf(message string, args ...interface{}) {
	l.logger.Errorf(message, args...)
}

func (l *zapLogger) Errorw(message string, args ...interface{}) {
	l.logger.Errorw(message, args...)
}

func (l *zapLogger) Fatalf(message string, args ...interface{}) {
	l.logger.Fatalf(message, args...)
}
