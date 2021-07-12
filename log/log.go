package log

import (
	"sync"

	"github.com/seerx/logo"
	"github.com/seerx/logo/logs"
)

var (
	instance logo.Logger
	once     sync.Once
)

func GetDefaultLogger() logo.Logger {
	once.Do(func() {
		instance = logs.NewLogger()
	})
	return instance
}

func WithError(err error) *logo.ProxyLogger {
	return GetDefaultLogger().WithError(err)
}

func WithData(data interface{}) *logo.ProxyLogger {
	return GetDefaultLogger().WithData(data)
}

func Debug(v ...interface{}) {
	GetDefaultLogger().Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	GetDefaultLogger().Debugf(format, v...)
}

func Info(v ...interface{}) {
	GetDefaultLogger().Info(v...)
}

func Infof(format string, v ...interface{}) {
	GetDefaultLogger().Infof(format, v...)
}

func Warn(v ...interface{}) {
	GetDefaultLogger().Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	GetDefaultLogger().Warnf(format, v...)
}

func Error(v ...interface{}) {
	GetDefaultLogger().Error(v...)
}

func Errorf(format string, v ...interface{}) {
	GetDefaultLogger().Errorf(format, v...)
}

func SetLevel(level logo.LogLevel) {
	GetDefaultLogger().SetLevel(level)
}

func SetLogErrorCallStacks(show bool) {
	GetDefaultLogger().SetLogErrorCallStacks(show)
}

func SetFormatter(fmt logo.Formatter) {
	GetDefaultLogger().SetFormatter(fmt)
}

func SetColorLog(color bool) {
	GetDefaultLogger().SetColorLog(color)
}

func SetLogFileLine(log bool) {
	GetDefaultLogger().SetLogFileLine(log)
}
