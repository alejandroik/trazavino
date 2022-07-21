package logger

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// Logger -.
type Logger struct {
	//logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	//l.msg("debug", message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	//l.log(message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	//l.log(message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	//if l.logger.GetLevel() == zerolog.DebugLevel {
	//	l.Debug(message, args...)
	//}
	//
	//l.msg("error", message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	//l.msg("fatal", message, args...)
	//
	//os.Exit(1)
}
