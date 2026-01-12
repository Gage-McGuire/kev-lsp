package handler

import (
	"encoding/json"
	"io"
	"log"

	"github.com/kev-lsp/analysis"
	"github.com/kev-lsp/lsp"
	"github.com/kev-lsp/rpc"
)

func HandleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, content []byte) {
	logger.Printf("[INFO] Handler Received: %s\n", method)
	switch method {
	case "initialize":
		handleInitialize(logger, writer, content)
	case "textDocument/didOpen":
		handleTextDocumentDidOpen(logger, state, content)
	case "textDocument/didChange":
		handleTextDocumentDidChange(logger, state, content)
	case "textDocument/hover":
		handleTextDocumentHover(logger, writer, content)
	}
}

func handleInitialize(logger *log.Logger, writer io.Writer, content []byte) {
	var request lsp.InitializeRequest
	err := json.Unmarshal(content, &request)
	if err != nil {
		logger.Printf("[ERROR] Initialize: %s\n\n", err)
	}
	logger.Printf("[INFO] Initialized: %s %s\n\n",
		request.Params.ClientInfo.Name,
		request.Params.ClientInfo.Version,
	)
	message := lsp.NewInitializeResponse(request.ID)
	rpc.WriteResponse(writer, message)
}

func handleTextDocumentDidOpen(logger *log.Logger, state analysis.State, content []byte) {
	var notification lsp.TextDocumentDidOpenNotification
	err := json.Unmarshal(content, &notification)
	if err != nil {
		logger.Printf("[ERROR] textDocument/didOpen: %s\n", err)
	}
	logger.Printf("[OPEN] textDocument/didOpen: %s %d\n",
		notification.Params.TextDocument.URI,
		notification.Params.TextDocument.Version,
	)
	state.OpenDocument(
		notification.Params.TextDocument.URI,
		notification.Params.TextDocument.Text,
	)
}

func handleTextDocumentDidChange(logger *log.Logger, state analysis.State, content []byte) {
	var notification lsp.TextDocumentDidChangeNotification
	err := json.Unmarshal(content, &notification)
	if err != nil {
		logger.Printf("[ERROR] textDocument/didChange: %s\n", err)
	}
	logger.Printf("[CHANGE] textDocument/didChange: %s %d\n",
		notification.Params.TextDocument.URI,
		notification.Params.TextDocument.Version,
	)
	for _, contentChange := range notification.Params.ContentChanges {
		state.UpdateDocument(
			notification.Params.TextDocument.URI,
			contentChange.Text,
		)
	}
}

func handleTextDocumentHover(logger *log.Logger, writer io.Writer, content []byte) {
	var request lsp.TextDocumentHoverRequest
	err := json.Unmarshal(content, &request)
	if err != nil {
		logger.Printf("[ERROR] textDocument/hover: %s\n", err)
	}
	logger.Printf("[HOVER] textDocument/hover: %s %d:%d\n",
		request.Params.TextDocument.URI,
		request.Params.Position.Line,
		request.Params.Position.Character,
	)
	message := lsp.TextDocumentHoverResponse{
		Response: lsp.Response{
			RPCVersion: "2.0",
			ID:         request.ID,
		},
		Result: lsp.TextDocumentHoverResult{
			Contents: "Hover content",
		},
	}
	rpc.WriteResponse(writer, message)
}
