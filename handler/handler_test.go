package handler

import (
	"os"
	"strings"
	"testing"

	"github.com/kev-lsp/logger"
)

func TestHandleInitialize(t *testing.T) {
	logFile := "test.log"
	defer os.Remove(logFile)
	logger := logger.GetLogger(logFile)

	content := []byte(`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"clientInfo":{"name":"kev-lsp","version":"0.0.1"}}}`)
	handleInitialize(logger, content)

	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Error reading log file: %s", err)
	}
	if !strings.Contains(string(content), "Initialized: kev-lsp") {
		t.Fatalf("\n\n===Expected===\n%s\n\n====Actual====\n%s\n\n",
			"Initialized: kev-lsp in log entry",
			string(content),
		)
	}
}
