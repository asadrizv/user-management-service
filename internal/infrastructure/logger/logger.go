package logger

import "log"

type Logger interface {
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type DefaultLogger struct{}

func (l DefaultLogger) Infof(format string, v ...interface{}) {
	log.Printf("[INFO] "+format, v...)
}

func (l DefaultLogger) Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR] "+format, v...)
}
