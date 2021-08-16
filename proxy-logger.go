package logo

import (
	"fmt"
	"runtime"
	"time"
)

type Entry struct {
	File           string      `json:"file"`
	Line           int         `json:"line"`
	Time           time.Time   `json:"-"`     // 转换
	TimeStr        string      `json:"time"`  //
	Level          LogLevel    `json:"-"`     // 转换
	LevelStr       string      `json:"level"` //
	Message        string      `json:"message"`
	Data           interface{} `json:"data"`
	Err            error       `json:"-"` // 转换
	hideCallStacks bool        `json:"-"` // 不需要输出
	Error          string      `json:"error"`
}

type ProxyLogger struct {
	Logger
	entry Entry
}

const callDepth = 2

func (e *ProxyLogger) WithData(data interface{}) *ProxyLogger {
	e.entry.Data = data
	return e
}
func (e *ProxyLogger) WithError(err error, hideCallStacks ...bool) *ProxyLogger {
	e.entry.Err = err
	if len(hideCallStacks) > 0 {
		e.entry.hideCallStacks = hideCallStacks[0]
	} else {
		e.entry.hideCallStacks = false
	}
	return e
}

func (e *ProxyLogger) CleanProxyDataAndError() *ProxyLogger {
	e.entry.Data = nil
	e.entry.Err = nil
	return e
}

// func (e *ProxyLogger) SetProxyLevelAndTime(level LogLevel, tim time.Time) *ProxyLogger {
// 	e.entry.Level = level
// 	e.entry.Time = tim
// 	return e
// }
// func (e *ProxyLogger) SetProxyMessagef(format string, v ...interface{}) *ProxyLogger {
// 	e.entry.Message = fmt.Sprintf(format, v...)
// 	return e
// }
// func (e *ProxyLogger) SetProxyMessage(v ...interface{}) *ProxyLogger {
// 	e.entry.Message = fmt.Sprint(v...)
// 	return e
// }
// func (e *ProxyLogger) GetEntry() *Entry {
// 	return &e.entry
// }

// func (e *ProxyLogger) Logout(out LogPrinter, fmt Formatter) {
// 	// plog.SetProxyLevelAndTime(level, time.Now())
// 	msg, err := fmt(&e.entry)
// 	if err != nil {
// 		e.WithError(err).Error("log format error")
// 		return
// 	}
// 	out.Print(msg)
// }

func (e *ProxyLogger) ProxyLog(level LogLevel, callDepth int, v ...interface{}) {
	e.entry.Level = level
	e.entry.Time = time.Now()
	e.entry.Message = fmt.Sprint(v...)

	if e.IsLogFileLine() {
		var ok bool
		_, e.entry.File, e.entry.Line, ok = runtime.Caller(callDepth)
		if !ok {
			e.entry.File = "???"
			e.entry.Line = 0
		}
	}

	if e.entry.Err != nil && !e.entry.hideCallStacks {
		stack := e.Logger.GetCallStack()
		if stack != nil {
			e.entry.Err = stack.WrapErrorSkip(e.entry.Err, 2)
		}
	}

	fmt := e.Logger.GetFormatter(level)
	printer := e.Logger.GetPrinter(level)
	msg, err := fmt(&e.entry)
	if err != nil {
		e.WithError(err, false).Error("log format error")
		return
	}
	printer.Print(msg, e.IsColorLog())
}

func (e *ProxyLogger) ProxyLogf(level LogLevel, callDepth int, format string, v ...interface{}) {
	e.entry.Level = level
	e.entry.Time = time.Now()
	e.entry.Message = fmt.Sprintf(format, v...)

	if e.IsLogFileLine() {
		var ok bool
		_, e.entry.File, e.entry.Line, ok = runtime.Caller(callDepth)
		if !ok {
			e.entry.File = "???"
			e.entry.Line = 0
		}
	}

	if e.entry.Err != nil && !e.entry.hideCallStacks {
		stack := e.Logger.GetCallStack()
		if stack != nil {
			e.entry.Err = stack.WrapErrorSkip(e.entry.Err, 2)
		}
	}

	fmt := e.Logger.GetFormatter(level)
	printer := e.Logger.GetPrinter(level)
	msg, err := fmt(&e.entry)
	if err != nil {
		e.WithError(err, false).Error("log format error")
		return
	}
	printer.Print(msg, e.IsColorLog())
}

func (e *ProxyLogger) Debug(v ...interface{}) {
	e.ProxyLog(LevelDebug, callDepth, v...)
}

func (e *ProxyLogger) Debugf(format string, v ...interface{}) {
	e.ProxyLogf(LevelDebug, callDepth, format, v...)
}

func (e *ProxyLogger) Info(v ...interface{}) {
	e.ProxyLog(LevelInfo, callDepth, v...)
}

func (e *ProxyLogger) Infof(format string, v ...interface{}) {
	e.ProxyLogf(LevelInfo, callDepth, format, v...)
}

func (e *ProxyLogger) Warn(v ...interface{}) {
	e.ProxyLog(LevelWarn, callDepth, v...)
}

func (e *ProxyLogger) Warnf(format string, v ...interface{}) {
	e.ProxyLogf(LevelWarn, callDepth, format, v...)
}

func (e *ProxyLogger) Error(v ...interface{}) {
	e.ProxyLog(LevelError, callDepth, v...)
}

func (e *ProxyLogger) Errorf(format string, v ...interface{}) {
	e.ProxyLogf(LevelError, callDepth, format, v...)
}
