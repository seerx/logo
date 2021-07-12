package logo

import "github.com/seerx/logo/tools"

type Logger interface {
	WithData(data interface{}) *ProxyLogger
	WithError(err error) *ProxyLogger
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	SetLevel(level LogLevel)
	GetLevel() LogLevel

	SetLogFileLine(log bool)
	IsLogFileLine() bool
	SetLogErrorCallStacks(log bool)
	GetCallStack() *tools.CallStack

	SetFormatter(fmt Formatter)
	GetFormatter(level LogLevel) Formatter

	GetPrinter(level LogLevel) LogPrinter

	SetColorLog(bool)
	IsColorLog() bool
}

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

var levels = []string{
	"debug", "info", "warn", "error",
}

func GetLevelString(level LogLevel) string {
	return levels[level]
}

// var levelColors = []string{
// 	"36m", "32m", "33m", "31m",
// }

// func GetLevelColor(level LogLevel) string {
// 	return levelColors[level]
// }
