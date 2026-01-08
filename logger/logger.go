package logger

import (
	"log"
	"os"
)

func GetLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return log.New(logFile, "[KEV-LSP] ", log.Ldate|log.Ltime|log.Lshortfile)
}
