package handler

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/kev-lsp/logger"
)

func TestHandleInitialize(t *testing.T) {
	logFile := "test.log"
	defer os.Remove(logFile)
	testLogger := logger.GetLogger(logFile)

	var writer bytes.Buffer

	content := []byte(`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"clientInfo":{"name":"kev-lsp","version":"0.0.1"}}}`)
	handleInitialize(testLogger, &writer, content)

	// Check the response written to the io.Writer
	response := writer.String()
	if !strings.Contains(response, "Content-Length:") {
		t.Fatalf("Response missing Content-Length header:\n%s", response)
	}
	if !strings.Contains(response, `"jsonrpc":"2.0"`) {
		t.Fatalf("Response missing jsonrpc version:\n%s", response)
	}
	if !strings.Contains(response, `"serverInfo"`) {
		t.Fatalf("Response missing serverInfo:\n%s", response)
	}

	// Check the log file
	logContent, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Error reading log file: %s", err)
	}
	if !strings.Contains(string(logContent), "Initialized: kev-lsp") {
		t.Fatalf("Initialized: kev-lsp not found in log file")
	}
}
