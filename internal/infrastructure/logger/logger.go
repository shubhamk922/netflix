package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var loggerInstance *logger
var once sync.Once

type logger struct {
	instance *zap.Logger
}

func InitLogger() *logger {
	once.Do(
		func() {
			NewLogger()
		},
	)
	return loggerInstance
}

func NewLogger() *logger {
	if loggerInstance == nil {
		l, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("Failed to instantiate logger isnstance %+v", err)
			return nil
		}
		loggerInstance = &logger{instance: l}
	}
	return loggerInstance
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func Close() {

	if loggerInstance.instance != nil {
		loggerInstance.instance.Sync()
	}
}

func (l *logger) Infof(msg string, args ...any) {
	if l.instance != nil {
		l.instance.Sugar().Infof(msg, args...)
	}
}

func (l *logger) Errorf(msg string, args ...any) {
	if l.instance != nil {
		l.instance.Sugar().Errorf(msg, args...)
	}
}

func (l *logger) Debugf(msg string, args ...any) {
	if l.instance != nil {
		l.instance.Sugar().Debugf(msg, args...)
	}
}

func (l *logger) Warnf(msg string, args ...any) {
	if l.instance != nil {
		l.instance.Sugar().Warnf(msg, args...)
	}
}
