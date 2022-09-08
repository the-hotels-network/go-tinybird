package tinybird

import (
	"github.com/sirupsen/logrus"
)

type LogrusAdapter struct{}

func (l LogrusAdapter) Debugf(fmt string, args ...interface{}) {
	logrus.Debugf(fmt, args...)
}

func NewLogger(a Adapter) Logger {
	return Logger{adapter: a}
}

type Adapter interface {
	Debugf(string, ...interface{})
}

func (l *Logger) Debugf(fmt string, args ...interface{}) {
	l.adapter.Debugf(fmt, args...)
}

type Logger struct {
	adapter Adapter
}
