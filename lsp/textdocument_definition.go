package lsp

type TextDocumentDefinitionRequest struct {
	Request
	Params TextDocumentDefinitionParams `json:"params"`
}

type TextDocumentDefinitionParams struct {
	TextDocumentPositionParams
}

type TextDocumentDefinitionResponse struct {
	Response
	Result Location `json:"result"`
}
