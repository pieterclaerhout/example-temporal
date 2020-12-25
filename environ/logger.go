package environ

import (
	"github.com/pieterclaerhout/go-log"
)

type Logger struct {
}

func NewLogger() Logger {
	log.PrintColors = true
	log.PrintTimestamp = true
	return Logger{}
}

func (l Logger) Debug(msg string, keyvals ...interface{}) {
	log.Debug(l.args(msg, keyvals)...)
}

func (l Logger) Info(msg string, keyvals ...interface{}) {
	log.Info(l.args(msg, keyvals)...)
}

func (l Logger) Warn(msg string, keyvals ...interface{}) {
	log.Warn(l.args(msg, keyvals)...)
}

func (l Logger) Error(msg string, keyvals ...interface{}) {
	log.Error(l.args(msg, keyvals)...)
}

func (l Logger) args(msg string, keyvals ...interface{}) []interface{} {
	args := []interface{}{msg}
	args = append(args, keyvals...)
	return args
}
