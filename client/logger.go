package client

import (
	"syscall/js"
)

type Logger struct {
	logger js.Value
}

func NewLogger() *Logger {
	return &Logger{
		logger: js.Global().Get("console"),
	}
}

func (l *Logger) Print(msg string) {
	l.logger.Call("log", msg)
}
