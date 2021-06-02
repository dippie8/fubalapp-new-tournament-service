package file

import (
	"log"
	"os"
	"time"
)

type Logger struct {}

func NewLogger() *Logger{
	return &Logger{}
}

func (m *Logger) Log(message string) {
	datetime := time.Now().String()
	f, err := os.OpenFile("/var/log/fubalapp/new-tournament",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, datetime, log.LstdFlags)
	logger.Println(message)

}