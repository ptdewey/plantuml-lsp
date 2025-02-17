package handler

import (
	"encoding/json"
	"io"

	"github.com/ptdewey/plantuml-lsp/internal/analysis"
	"github.com/ptdewey/plantuml-lsp/internal/lsp"
	"github.com/ptdewey/plantuml-lsp/internal/rpc"
)

// FIX: logging to lsp log not working
func HandleMessage(writer io.Writer, state analysis.State, method string, contents []byte, stdlibPath string, execPath []string) {
	SendLogMessage(writer, "Received msg with method: "+method, lsp.Debug)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			SendLogMessage(writer, "Could not parse initialize request: "+err.Error(), lsp.Error)
			return
		}

		SendLogMessage(writer, "Connected to: "+request.Params.ClientInfo.Name+" "+request.Params.ClientInfo.Version, lsp.Info)

		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg)

		go func() {
			SendLogMessage(writer, "Fetching language features...", lsp.Debug)
			state.GetFeatures(stdlibPath)
			SendLogMessage(writer, "Language features initialized.", lsp.Debug)
		}()

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			SendLogMessage(writer, "textDocument/didOpen: "+err.Error(), lsp.Error)
			return
		}

		SendLogMessage(writer, "Opened: "+request.Params.TextDocument.URI, lsp.Debug)
		diagnostics := state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text, execPath)
		writeResponse(writer, lsp.PublishDiagnosticsNotification{
			Notification: lsp.Notification{
				RPC:    "2.0",
				Method: "textDocument/publishDiagnostics",
			},
			Params: lsp.PublishDiagnosticsParams{
				URI:         request.Params.TextDocument.URI,
				Diagnostics: diagnostics,
			},
		})

	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			SendLogMessage(writer, "textDocument/didChange: "+err.Error(), lsp.Error)
			return
		}

		SendLogMessage(writer, "Changed: "+request.Params.TextDocument.URI, lsp.Debug)

		// Goroutine here may be band-aid fix for diagnostics performance
		go func() {
			for _, change := range request.Params.ContentChanges {
				diagnostics := state.UpdateDocument(request.Params.TextDocument.URI, change.Text, execPath)
				writeResponse(writer, lsp.PublishDiagnosticsNotification{
					Notification: lsp.Notification{
						RPC:    "2.0",
						Method: "textDocument/publishDiagnostics",
					},
					Params: lsp.PublishDiagnosticsParams{
						URI:         request.Params.TextDocument.URI,
						Diagnostics: diagnostics,
					},
				})
			}
		}()

	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			SendLogMessage(writer, "textDocument/hover: "+err.Error(), lsp.Warning)
			return
		}

		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)

	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			SendLogMessage(writer, "textDocument/definition: "+err.Error(), lsp.Warning)
			return
		}

		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)

	case "textDocument/codeAction":
		var request lsp.CodeActionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			SendLogMessage(writer, "textDocument/codeAction: "+err.Error(), lsp.Warning)
			return
		}

		response := state.TextDocumentCodeAction(request.ID, request.Params.TextDocument.URI)
		writeResponse(writer, response)

	case "textDocument/completion":
		var request lsp.CompletionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			SendLogMessage(writer, "textDocument/codeAction: "+err.Error(), lsp.Warning)
			return
		}

		response := state.TextDocumentCompletion(request.ID, request.Params.TextDocument.URI)
		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply, err := rpc.EncodeMessage(msg)
	if err != nil {
		SendLogMessage(writer, "Error encoding response: "+err.Error(), lsp.Error)
		return
	}

	writer.Write([]byte(reply))
}

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

	writeResponse(writer, logMessage)
}
