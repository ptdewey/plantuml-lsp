package logger

import (
	"io"

	"github.com/ptdewey/plantuml-lsp/internal/lsp"
	"github.com/ptdewey/plantuml-lsp/internal/rpc"
)

// REFACTOR: init function that takes in writer and spawns logger instance, remove writer param from `SendLogMessage`
func SendLogMessage(writer io.Writer, message string, level int) {
	logMessage := lsp.LogMessage{
		Notification: lsp.Notification{
			RPC:    "2.0",
			Method: "window/logMessage",
		},
		Params: lsp.LogMessageParams{
			Type:    level,
			Message: message,
		},
	}

	WriteResponse(writer, logMessage)
}

// REFACTOR: use logger instance's writer
func WriteResponse(writer io.Writer, msg any) {
	reply, err := rpc.EncodeMessage(msg)
	if err != nil {
		SendLogMessage(writer, "Error encoding response: "+err.Error(), lsp.Error)
		return
	}

	if _, err := writer.Write([]byte(reply)); err != nil {
		SendLogMessage(writer, "Error writing reponse: "+err.Error(), lsp.Error)
		return
	}
}
