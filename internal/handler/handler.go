package handler

import (
	"encoding/json"
	"io"
	"time"

	"github.com/ptdewey/plantuml-lsp/internal/analysis"
	"github.com/ptdewey/plantuml-lsp/internal/lsp"
	"github.com/ptdewey/plantuml-lsp/internal/rpc"
	"github.com/ptdewey/plantuml-lsp/internal/utils/debounce"
)

var debounceDelay time.Duration = 500 * time.Millisecond

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

		uri := request.Params.TextDocument.URI

		SendLogMessage(writer, "Changed: "+uri, lsp.Debug)

		d, exists := state.Timers["textDocument/didChange"]
		if !exists {
			d = debounce.New(debounceDelay,
				func(args ...any) {
					uri := args[0].(string)
					changes := args[1].([]lsp.TextDocumentContentChangeEvent)
					for _, change := range changes {
						diagnostics := state.UpdateDocument(uri, change.Text, execPath)
						writeResponse(writer, lsp.PublishDiagnosticsNotification{
							Notification: lsp.Notification{
								RPC:    "2.0",
								Method: "textDocument/publishDiagnostics",
							},
							Params: lsp.PublishDiagnosticsParams{
								URI:         uri,
								Diagnostics: diagnostics,
							},
						})
					}
				})
			state.Timers["textDocument/didChange"] = d
		}

		// Debounce diagnostics to avoid spawning many plantuml processes at once
		d.Set(uri, request.Params.ContentChanges)
		d.Debounced()

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

	if _, err := writer.Write([]byte(reply)); err != nil {
		SendLogMessage(writer, "Error writing reponse: "+err.Error(), lsp.Error)
		return
	}
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
