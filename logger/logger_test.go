package logger

import (
	"os"
	"strings"
	"testing"
)

func TestGetLogger(t *testing.T) {
	logFile := "test.log"
	defer func() {
		_ = os.Remove(logFile)
	}()

	logger := GetLogger(logFile)
	if logger == nil {
		t.Fatalf("Expected logger to not be nil")
	}
	logger.Println("Test message")

	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Error reading log file: %s", err)
	}
	if !strings.Contains(string(content), "Test message") {
		t.Fatalf("\n\n===Expected===\n%s\n\n====Actual====\n%s\n\n",
			"Test message in log entry",
			string(content),
		)
	}
}
