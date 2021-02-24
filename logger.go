package tower

import (
	"log"
	"os"
)

type LogLevel int

const (
	Silent LogLevel = iota + 1
	Error
	Warn
	Info
	Debug
)

type LogWrite interface {
	Printf(string, ...interface{})
}

type Logger interface {
	LogMode(lvl LogLevel) Logger
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
}

func NewLogger(w LogWrite, lvl LogLevel) Logger {
	return &logging{
		LogWrite: w,
		LogLevel: lvl,
	}
}

var (
	defaultLogging = NewLogger(log.New(os.Stdout, "[Tower] ", log.Ldate|log.Ltime|log.LUTC), Debug)
)

type logging struct {
	LogWrite
	LogLevel LogLevel
}

func (l *logging) LogMode(lvl LogLevel) Logger {
	l.LogLevel = lvl
	return l
}

func (l *logging) Debug(msg string, v ...interface{}) {
	if l.LogLevel >= Debug {
		l.Printf(msg, v...)
	}
}

func (l *logging) Info(msg string, v ...interface{}) {
	if l.LogLevel >= Info {
		l.Printf(msg, v...)
	}
}

func (l *logging) Warn(msg string, v ...interface{}) {
	if l.LogLevel >= Warn {
		l.Printf(msg, v...)
	}
}

func (l *logging) Error(msg string, v ...interface{}) {
	if l.LogLevel >= Error {
		l.Printf(msg, v...)
	}
}
