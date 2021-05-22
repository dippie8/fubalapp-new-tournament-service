package console

import (
	"fmt"
	"time"
)

type Logger struct {}

func NewLogger() *Logger{
	return &Logger{}
}

func (m *Logger) Log(message string) error {
	datetime := time.Now().String()
	fmt.Println(datetime + ": " + message)
	return nil
}