package console

import (
	"fmt"
	"time"
)

type Logger struct {}

func NewLogger() *Logger{
	return &Logger{}
}

func (m *Logger) Log(message string) {
	datetime := time.Now().String()
	fmt.Println(datetime + ": " + message)
}