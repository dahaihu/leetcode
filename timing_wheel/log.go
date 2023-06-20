package timing_wheel

import "log"

var l Log

func init() {
	l = &DefaultLog{}
}

// SetLog sets the log for timer package
func SetLog(log Log) {
	l = log
}

// Log is the log interface for timer package
// user can impletment this interface to use their own log
type Log interface {
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
}

type DefaultLog struct{}

func (l *DefaultLog) Debugf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *DefaultLog) Debug(args ...interface{}) {
	log.Println(args...)
}

func (l *DefaultLog) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *DefaultLog) Info(args ...interface{}) {
	log.Println(args...)
}

func (l *DefaultLog) Warnf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *DefaultLog) Warn(args ...interface{}) {
	log.Println(args...)
}

func (l *DefaultLog) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *DefaultLog) Error(args ...interface{}) {
	log.Println(args...)
}

func (l *DefaultLog) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func (l *DefaultLog) Fatal(args ...interface{}) {
	log.Fatal(args...)
}
