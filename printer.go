package logo

import (
	"fmt"
	"io"
)

type LogPrinter interface {
	SetOutput(writer io.Writer)
	Print(content []byte, color bool) error
}

type StdLogPrinter struct {
	writer io.Writer
	color  string
}

func NewStdLogPrinter(writer io.Writer, color string) *StdLogPrinter {
	return &StdLogPrinter{
		writer: writer,
		color:  color,
	}
}

func (s *StdLogPrinter) SetOutput(writer io.Writer) {
	s.writer = writer
}

func (s *StdLogPrinter) Print(message []byte, color bool) error {
	// io.MultiWriter()
	var err error
	if s.color == "" || !color {
		_, err = fmt.Fprint(s.writer, string(message))
	} else {
		_, err = fmt.Fprintf(s.writer, "\033[%s%s\033[0m", s.color, string(message))
	}
	return err
}
