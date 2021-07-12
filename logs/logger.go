package logs

import (
	"io"
	"os"
	"sync"

	"github.com/seerx/logo"
	"github.com/seerx/logo/tools"
)

type logger struct {
	logFileLine bool
	fmt         logo.Formatter
	logLevel    logo.LogLevel
	loggers     []logo.LogPrinter
	mutex       sync.Mutex
	entryPool   sync.Pool
	stack       *tools.CallStack
	colorOutput bool
}

const callDepth = 3

func NewLogger() *logger {
	loggers := []logo.LogPrinter{
		logo.NewStdLogPrinter(io.Discard, "36m"),
		logo.NewStdLogPrinter(os.Stdout, "32m"),
		logo.NewStdLogPrinter(os.Stdout, "33m"),
		logo.NewStdLogPrinter(os.Stdout, "31m"),
	}
	return &logger{
		fmt:         logo.TextFormatter,
		logLevel:    logo.LevelInfo,
		loggers:     loggers,
		colorOutput: true,
	}
}

func (l *logger) NewProxyLogger() *logo.ProxyLogger {
	entry, ok := l.entryPool.Get().(*logo.ProxyLogger)
	if ok {
		return entry
	}
	return &logo.ProxyLogger{
		Logger: l,
	}
}

func (l *logger) ReleaseProxyLogger(plog *logo.ProxyLogger) {
	plog.CleanProxyDataAndError()
	l.entryPool.Put(plog)
}

func (l *logger) WithData(data interface{}) *logo.ProxyLogger {
	e := l.NewProxyLogger()
	return e.WithData(data)
}

func (l *logger) WithError(err error) *logo.ProxyLogger {
	e := l.NewProxyLogger()
	return e.WithError(err)
}

func (l *logger) Debug(v ...interface{}) {
	l.NewProxyLogger().ProxyLog(logo.LevelDebug, callDepth, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.NewProxyLogger().ProxyLogf(logo.LevelDebug, callDepth, format, v...)
}

func (l *logger) Info(v ...interface{}) {
	l.NewProxyLogger().ProxyLog(logo.LevelInfo, callDepth, v...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.NewProxyLogger().ProxyLogf(logo.LevelInfo, callDepth, format, v...)
}

func (l *logger) Warn(v ...interface{}) {
	l.NewProxyLogger().ProxyLog(logo.LevelWarn, callDepth, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.NewProxyLogger().ProxyLogf(logo.LevelWarn, callDepth, format, v...)
}

func (l *logger) Error(v ...interface{}) {
	l.NewProxyLogger().ProxyLog(logo.LevelError, callDepth, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.NewProxyLogger().ProxyLogf(logo.LevelError, callDepth, format, v...)
}

func (l *logger) SetLevel(level logo.LogLevel) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logLevel = level
	for n, logger := range l.loggers {
		if n >= int(level) {
			logger.SetOutput(os.Stdout)
		} else {
			logger.SetOutput(io.Discard)
		}
	}

}
func (l *logger) GetLevel() logo.LogLevel { return l.logLevel }

func (l *logger) SetLogFileLine(log bool) {
	l.logFileLine = log
}

func (l *logger) IsLogFileLine() bool {
	return l.logFileLine
}

func (l *logger) SetLogErrorCallStacks(log bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if log {
		l.stack = tools.NewCallStack()
	} else {
		l.stack = nil
	}
}

func (l *logger) GetCallStack() *tools.CallStack { return l.stack }

func (l *logger) SetFormatter(fmt logo.Formatter) { l.fmt = fmt }

func (l *logger) GetFormatter(level logo.LogLevel) logo.Formatter {
	return l.fmt
}

func (l *logger) GetPrinter(level logo.LogLevel) logo.LogPrinter {
	return l.loggers[level]
}

func (l *logger) SetColorLog(color bool) {
	l.colorOutput = color
}

func (l *logger) IsColorLog() bool {
	return l.colorOutput
}
