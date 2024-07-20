package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// TODO: params
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id"`

	// TODO: result, error
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
