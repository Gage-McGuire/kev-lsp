package lsp

type Request struct {
	RPCVersion string `json:"jsonrpc"`
	ID         int    `json:"id"`
	Method     string `json:"method"`

	// Params ...
	// Specified in the Request types
}

type Response struct {
	RPCVersion string `json:"jsonrpc"`
	ID         int    `json:"id"`

	// Result ...
	// Error ...
}

type Notification struct {
	RPCVersion string `json:"jsonrpc"`
	Method     string `json:"method"`
}
