package logger

import (
	"log"
	"os"
)

type Logger struct {
	Service string
	Funtion string
	LogType string
	Msg     string
}

func (l *Logger) LogToFile(filename string) {
	logFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, l.LogType+": ", log.LstdFlags)
	logger.SetOutput(logFile)
	logger.Println(l.Msg)
}
