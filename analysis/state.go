package analysis

import (
	"fmt"
	"strings"

	"github.com/Gage-McGuire/kev-lsp/lsp"
	"github.com/Gage-McGuire/kev-lsp/rpc"
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

func getDiagnostics(text string) []lsp.Diagnostic {
	// TODO: Add the actual diagnostics look up here
	diagnostics := []lsp.Diagnostic{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "let")
		if idx >= 0 {
			diagnostic := lsp.Diagnostic{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      row,
						Character: idx,
					},
					End: lsp.Position{
						Line:      row,
						Character: idx + 3,
					},
				},
				Severity: lsp.SeverityError,
				Source:   "KEV LSP",
				Message:  "let is not allowed, use var instead",
			}
			diagnostics = append(diagnostics, diagnostic)
		}
	}
	return diagnostics
}

func (s *State) OpenDocument(documentURI string, text string) []lsp.Diagnostic {
	s.Documents[documentURI] = text
	return getDiagnostics(text)
}

func (s *State) UpdateDocument(documentURI string, text string) []lsp.Diagnostic {
	s.Documents[documentURI] = text
	return getDiagnostics(text)
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

func (s *State) OnCodeAction(id int, documentURI string) lsp.TextDocumentCodeActionResponse {
	documentContents := s.Documents[documentURI]
	// TODO: Add support with diagnostics
	actions := []lsp.TextDocumentCodeAction{}
	for row, line := range strings.Split(documentContents, "\n") {
		idx := strings.Index(line, "let")
		if idx >= 0 {
			change := map[string][]lsp.TextEdit{}
			change[documentURI] = []lsp.TextEdit{
				{
					Range: lsp.Range{
						Start: lsp.Position{
							Line:      row,
							Character: idx,
						},
						End: lsp.Position{
							Line:      row,
							Character: idx + 3,
						},
					},
					NewText: "var",
				},
			}
			actions = append(actions, lsp.TextDocumentCodeAction{
				Title: "var",
				Edit: &lsp.WorkspaceEdit{
					Changes: change,
				},
			})
		}
	}

	message := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPCVersion: rpc.RPCVersion,
			ID:         id,
		},
		Result: actions,
	}
	return message
}

func (s *State) OnCompletion(id int) lsp.TextDocumentCompletionResponse {
	// TODO: Add the actual completion look up here
	items := []lsp.CompletionItem{
		{
			Label:         "var",
			Detail:        "Variable",
			Documentation: "A variable is a named value that can be used to store data.",
		},
		{
			Label:         "func",
			Detail:        "Function",
			Documentation: "A function is a named block of code that can be used to perform a task.",
		},
	}
	message := lsp.TextDocumentCompletionResponse{
		Response: lsp.Response{
			RPCVersion: rpc.RPCVersion,
			ID:         id,
		},
		Result: items,
	}
	return message
}
