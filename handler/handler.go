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
		handleTextDocumentDidChange(logger, writer, state, content)
	case "textDocument/hover":
		handleTextDocumentHover(logger, writer, state, content)
	case "textDocument/definition":
		handleTextDocumentDefinition(logger, writer, state, content)
	case "textDocument/codeAction":
		handleTextDocumentCodeAction(logger, writer, state, content)
	case "textDocument/completion":
		handleTextDocumentCompletion(logger, writer, state, content)
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
	logger.Printf("[INFO] textDocument/didOpen: %s %d\n",
		notification.Params.TextDocument.URI,
		notification.Params.TextDocument.Version,
	)
	state.OpenDocument(
		notification.Params.TextDocument.URI,
		notification.Params.TextDocument.Text,
	)
}

func handleTextDocumentDidChange(logger *log.Logger, writer io.Writer, state analysis.State, content []byte) {
	var notification lsp.TextDocumentDidChangeNotification
	err := json.Unmarshal(content, &notification)
	if err != nil {
		logger.Printf("[ERROR] textDocument/didChange: %s\n", err)
	}
	logger.Printf("[INFO] textDocument/didChange: %s %d\n",
		notification.Params.TextDocument.URI,
		notification.Params.TextDocument.Version,
	)
	for _, contentChange := range notification.Params.ContentChanges {
		diagnostics := state.UpdateDocument(
			notification.Params.TextDocument.URI,
			contentChange.Text,
		)
		message := lsp.PublishDiagnosticsNotification{
			Notification: lsp.Notification{
				RPCVersion: rpc.RPCVersion,
				Method:     "textDocument/publishDiagnostics",
			},
			Params: lsp.PublishDiagnosticsParams{
				URI:         notification.Params.TextDocument.URI,
				Diagnostics: diagnostics,
			},
		}
		rpc.WriteResponse(writer, message)
	}
}

func handleTextDocumentHover(logger *log.Logger, writer io.Writer, state analysis.State, content []byte) {
	var request lsp.TextDocumentHoverRequest
	err := json.Unmarshal(content, &request)
	if err != nil {
		logger.Printf("[ERROR] textDocument/hover: %s\n", err)
	}
	logger.Printf("[INFO] textDocument/hover: %s %d:%d\n",
		request.Params.TextDocument.URI,
		request.Params.Position.Line,
		request.Params.Position.Character,
	)
	message := state.OnHover(
		request.ID,
		request.Params.TextDocument.URI,
		request.Params.Position,
	)
	rpc.WriteResponse(writer, message)
}

func handleTextDocumentDefinition(logger *log.Logger, writer io.Writer, state analysis.State, content []byte) {
	var request lsp.TextDocumentDefinitionRequest
	err := json.Unmarshal(content, &request)
	if err != nil {
		logger.Printf("[ERROR] textDocument/definition: %s\n", err)
	}
	logger.Printf("[INFO] textDocument/definition: %s %d:%d",
		request.Params.TextDocument.URI,
		request.Params.Position.Line,
		request.Params.Position.Character,
	)
	message := state.OnDefinition(
		request.ID,
		request.Params.TextDocument.URI,
		request.Params.Position,
	)
	rpc.WriteResponse(writer, message)
}

func handleTextDocumentCodeAction(logger *log.Logger, writer io.Writer, state analysis.State, content []byte) {
	var request lsp.TextDocumentCodeActionRequest
	err := json.Unmarshal(content, &request)
	if err != nil {
		logger.Printf("[ERROR] textDocument/codeAction: %s\n", err)
	}
	logger.Printf("[INFO] textDocument/codeAction: %s %d:%d",
		request.Params.TextDocument.URI,
		request.Params.Range.Start.Line,
		request.Params.Range.Start.Character,
	)
	message := state.OnCodeAction(request.ID, request.Params.TextDocument.URI)
	rpc.WriteResponse(writer, message)
}

func handleTextDocumentCompletion(logger *log.Logger, writer io.Writer, state analysis.State, content []byte) {
	var request lsp.TextDocumentCompletionRequest
	err := json.Unmarshal(content, &request)
	if err != nil {
		logger.Printf("[ERROR] textDocument/completion: %s\n", err)
	}
	logger.Printf("[INFO] textDocument/completion: %s %d:%d\n",
		request.Params.TextDocument.URI,
		request.Params.Position.Line,
		request.Params.Position.Character,
	)
	message := state.OnCompletion(request.ID)
	rpc.WriteResponse(writer, message)
}
