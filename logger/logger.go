package logger

import (
	"fmt"
	"log"
)

// Logger provides a logging interface.
type Logger interface {
	// Debug prints a debug log.
	Debug(format string, v ...interface{})
	// Info prints an info log.
	Info(format string, v ...interface{})
	// Warn prints a warn log.
	Warn(format string, v ...interface{})
	// Error prints an error log.
	Error(format string, v ...interface{})
	// SetLevel sets the verbose logging level.
	SetLevel(lv Level)
}

// Level indicates the verbose logging level.
type Level int

const (
	Ldebug Level = 0
	Linfo  Level = 1
	Lwarn  Level = 2
	Lerror Level = 3
)

var (
	instance Logger = &loggerImpl{}
)

// Get returns the Logger.
func Get() Logger { return instance }

type loggerImpl struct {
	verbose Level
}

var levelToLabel = map[Level]string{
	Ldebug: "D",
	Linfo:  "I",
	Lwarn:  "W",
	Lerror: "E",
}

func (s *loggerImpl) SetLevel(lv Level) { s.verbose = lv }

func (s *loggerImpl) Debug(format string, v ...interface{}) {
	s.outputWithLevel(Ldebug, format, v...)
}

func (s *loggerImpl) Info(format string, v ...interface{}) {
	s.outputWithLevel(Linfo, format, v...)
}

func (s *loggerImpl) Warn(format string, v ...interface{}) {
	s.outputWithLevel(Lwarn, format, v...)
}

func (s *loggerImpl) Error(format string, v ...interface{}) {
	s.outputWithLevel(Lerror, format, v...)
}

func (s *loggerImpl) outputWithLevel(lv Level, format string, v ...interface{}) {
	if lv >= s.verbose {
		x := fmt.Sprintf(format, v...)
		s.output("%s %s", levelToLabel[lv], x)
	}
}

func (*loggerImpl) output(format string, v ...interface{}) { log.Printf(format, v...) }
