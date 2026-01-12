package analysis

import (
	"fmt"

	"github.com/kev-lsp/lsp"
	"github.com/kev-lsp/rpc"
)

type State struct {
	// documentURI -> text
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: make(map[string]string),
	}
}

func (s *State) OpenDocument(documentURI string, text string) {
	s.Documents[documentURI] = text
}

func (s *State) UpdateDocument(documentURI string, text string) {
	s.Documents[documentURI] = text
}

func (s *State) OnHover(id int, documentURI string, position lsp.Position) lsp.TextDocumentHoverResponse {
	documentContents := s.Documents[documentURI]
	message := lsp.TextDocumentHoverResponse{
		Response: lsp.Response{
			RPCVersion: rpc.RPCVersion,
			ID:         id,
		},
		Result: lsp.TextDocumentHoverResult{
			Contents: fmt.Sprintf("File: %s, Characters: %d", documentURI, len(documentContents)),
		},
	}
	return message
}

func (s *State) OnDefinition(id int, documentURI string, position lsp.Position) lsp.TextDocumentDefinitionResponse {
	// TODO: Add the actual definition look up here
	// Right now we're just returning the line above the cursor
	message := lsp.TextDocumentDefinitionResponse{
		Response: lsp.Response{
			RPCVersion: rpc.RPCVersion,
			ID:         id,
		},
		Result: lsp.Location{
			URI: documentURI,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
	return message
}
