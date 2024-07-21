package lsp

const (
	Error   = 1
	Warning = 2
	Info    = 3
	Log     = 4
	Debug   = 5
)

type LogMessage struct {
	Notification
	Params LogMessageParams `json:"params"`
}

type LogMessageParams struct {
	Type    int    `json:"type"`
	Message string `json:"message"`
}
