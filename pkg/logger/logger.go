package logger

type Interface interface {
	Debugf(message string, args ...interface{})
	Infof(message string, args ...interface{})
	Infow(message string, args ...interface{})
	Warnf(message string, args ...interface{})
	Warnw(message string, args ...interface{})
	Errorf(message string, args ...interface{})
	Errorw(message string, args ...interface{})
	Fatalf(message string, args ...interface{})
}
