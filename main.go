package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kev-lsp/rpc"
)

func main() {
	logger := getLogger("kev-lsp.log")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		message := scanner.Text()
		handleMessage(logger, message)
	}
}

func handleMessage(logger *log.Logger, message any) {
	logger.Println(message)
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return log.New(logFile, "[KEV-LSP] ", log.Ldate|log.Ltime|log.Lshortfile)
}
