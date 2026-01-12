package main

import (
	"bufio"
	"os"

	"github.com/joho/godotenv"
	"github.com/kev-lsp/analysis"
	"github.com/kev-lsp/handler"
	"github.com/kev-lsp/logger"
	"github.com/kev-lsp/rpc"
)

func main() {
	godotenv.Load()
	logFile := os.Getenv("LOG_FILE")
	logger := logger.GetLogger(logFile)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		message := scanner.Bytes()
		method, content, err := rpc.Decode(message)
		if err != nil {
			logger.Printf("[ERROR] RPC Decode: %s\n", err)
			continue
		}
		handler.HandleMessage(logger, writer, state, method, content)
	}
}
