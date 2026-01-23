package lsp

type TextDocumentCodeActionRequest struct {
	Request
	Params TextDocumentCodeActionParams `json:"params"`
}

type TextDocumentCodeActionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range        Range                  `json:"range"`
	Context      CodeActionContext      `json:"codeActionContext"`
}

type TextDocumentCodeActionResponse struct {
	Response
	Result []TextDocumentCodeAction `json:"result"`
}

type CodeActionContext struct {
	// TODO: Add the actual context here
}

type TextDocumentCodeAction struct {
	Title   string         `json:"title"`
	Edit    *WorkspaceEdit `json:"edit,omitempty"`
	Command *Command       `json:"command,omitempty"`
}

type Command struct {
	Title     string `json:"title"`
	Command   string `json:"command"`
	Arguments []any  `json:"arguments,omitempty"`
}
