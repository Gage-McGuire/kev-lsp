package handler

import (
	"encoding/json"
	"log"
	"os"

	"github.com/kev-lsp/lsp"
	"github.com/kev-lsp/rpc"
)

func HandleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("Received Message: %s\n\n", method)
	switch method {
	case "initialize":
		handleInitialize(logger, content)
	case "textDocument/didOpen":
		handleTextDocumentDidOpen(logger, content)
	}
}

func handleInitialize(logger *log.Logger, content []byte) {
	var request lsp.InitializeRequest
	err := json.Unmarshal(content, &request)
	if err != nil {
		logger.Printf("Error unmarshalling message: %s\n\n", err)
	}
	logger.Printf("Initialized: %s %s\n\n",
		request.Params.ClientInfo.Name,
		request.Params.ClientInfo.Version,
	)
	message := lsp.NewInitializeResponse(request.ID)
	response := rpc.Encode(message)
	writer := os.Stdout
	writer.Write([]byte(response))

	logger.Printf("Sent Response: %s\n\n", response)
}

func handleTextDocumentDidOpen(logger *log.Logger, content []byte) {
	var notification lsp.TextDocumentDidOpenNotification
	err := json.Unmarshal(content, &notification)
	if err != nil {
		logger.Printf("Error unmarshalling message: %s\n\n", err)
	}
	logger.Printf("Text Document Opened: %s %s\n\n",
		notification.Params.TextDocument.URI,
		notification.Params.TextDocument.Text,
	)
}
