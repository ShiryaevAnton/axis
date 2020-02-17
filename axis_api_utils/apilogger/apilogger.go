package apilogger

import (
	"fmt"

	"go.uber.org/zap"
)

//Logger ...
type Logger interface {
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})
	Warning(args ...interface{})
	Warningln(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})
	V(l int) bool
	Test(string)
}

type logger struct {
	log *zap.Logger
}

//NewLogger ...
func NewLogger(log *zap.Logger) Logger {
	return &logger{log: log}
}

//TODO

func (l *logger) Test(msg string) {
	l.log.Info(msg)
}

func (l *logger) Info(args ...interface{}) {
	msg := fmt.Sprintf("%v", args)
	l.log.Info(msg)
}

func (l *logger) Infoln(args ...interface{}) {

}

func (l *logger) Infof(format string, args ...interface{}) {

}

func (l *logger) Warning(args ...interface{}) {

}

func (l *logger) Warningln(args ...interface{}) {}

func (l *logger) Warningf(format string, args ...interface{}) {

}

func (l *logger) Error(args ...interface{}) {

}

func (l *logger) Errorln(args ...interface{}) {

}

func (l *logger) Errorf(format string, args ...interface{}) {

}

func (l *logger) Fatal(args ...interface{}) {
	msg := fmt.Sprintf("%v:%v", args[0], args[1])
	l.log.Fatal(msg)
}

func (l *logger) Fatalln(args ...interface{}) {

}

func (l *logger) Fatalf(format string, args ...interface{}) {

}

func (l *logger) V(level int) bool {
	return false
}
