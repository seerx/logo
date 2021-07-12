package logo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type Formatter func(*Entry) ([]byte, error)

func TextFormatter(e *Entry) ([]byte, error) {
	// data := ""
	var err error
	buf := &bytes.Buffer{}
	if e.File != "" {
		// 等级 时间 文件 行号 信息
		_, err = buf.WriteString(fmt.Sprintf("[%5s] %s %s:%d\n%s", GetLevelString(e.Level), e.Time.Format(time.RFC3339), e.File, e.Line, e.Message))
		if err != nil {
			return nil, err
		}
	} else {
		// 等级 时间 信息
		_, err = buf.WriteString(fmt.Sprintf("[%5s] %s %s", GetLevelString(e.Level), e.Time.Format(time.RFC3339), e.Message))
		if err != nil {
			return nil, err
		}
	}

	// str.Get()
	// strings.Builder
	if e.Data != nil {
		jsonData, err := json.MarshalIndent(e.Data, "", "  ")
		if err != nil {
			return nil, err
		}
		_, err = buf.WriteString("\nDATA: ")
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(jsonData)
		if err != nil {
			return nil, err
		}
	}

	// errString := ""
	if e.Err != nil {
		// errString = e.Err.Error()
		_, err = buf.WriteString(fmt.Sprintf("\nERROR: %s", e.Err.Error()))
		if err != nil {
			return nil, err
		}
	}
	err = buf.WriteByte('\n')
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
	// return []byte(fmt.Sprintf("[%5s] %s %s\n%s\n%s", GetLevelString(e.Level), e.Time.Format(time.RFC3339), e.Message, data, errString)), nil
}

func JSONFormatter(e *Entry) ([]byte, error) {
	e.TimeStr = e.Time.Format(time.RFC3339)
	e.LevelStr = GetLevelString(e.Level)
	if e.Err != nil {
		e.Error = e.Err.Error()
	}
	// format = "[%s] %s %s"
	// data := ""
	// if e.Data != nil {
	return json.MarshalIndent(e, "", "  ")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	// data = string(jsonData)
	// }
	// return []byte(fmt.Sprintf("[%5s] %s %s\n%s", e.Level, e.Time, e.Message, data)), nil
}
