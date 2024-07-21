package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"plantuml_lsp/analysis"
	"plantuml_lsp/lsp"
	"plantuml_lsp/rpc"
)

func main() {
	// TODO: pass in plantuml_lsp.rc file to use for config stuff
	// - include log level https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#window_logMessage

	logPath := flag.String("log-path", "", "LSP log path")
	stdlibPath := flag.String("stdlib-path", "", "PlantUML stdlib path")
	flag.Parse()

	logger := getLogger(*logPath)
	logger.Println("Started plantuml-lsp")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			sendLogMessage(writer, "Error decoding message: "+err.Error(), lsp.Error)
			continue
		}

		handleMessage(writer, state, method, contents, *stdlibPath)
	}
}

// FIX: logging to lsp log not working
func handleMessage(writer io.Writer, state analysis.State, method string, contents []byte, stdlibPath string) {
	sendLogMessage(writer, "Received msg with method: "+method, lsp.Debug)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			sendLogMessage(writer, "Could not parse initialize request: "+err.Error(), lsp.Error)
			return
		}

		sendLogMessage(writer, "Connected to: "+request.Params.ClientInfo.Name+" "+request.Params.ClientInfo.Version, lsp.Info)

		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg)

		go func() {
			sendLogMessage(writer, "Fetching language features...", lsp.Debug)
			state.GetFeatures(stdlibPath)
			sendLogMessage(writer, "Language features initialized.", lsp.Debug)
		}()

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			sendLogMessage(writer, "textDocument/didOpen: "+err.Error(), lsp.Error)
			return
		}

		sendLogMessage(writer, "Opened: "+request.Params.TextDocument.URI, lsp.Debug)
		diagnostics := state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
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
			sendLogMessage(writer, "textDocument/didChange: "+err.Error(), lsp.Error)
			return
		}

		sendLogMessage(writer, "Changed: "+request.Params.TextDocument.URI, lsp.Debug)
		for _, change := range request.Params.ContentChanges {
			diagnostics := state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
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

	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			sendLogMessage(writer, "textDocument/hover: "+err.Error(), lsp.Warning)
			return
		}

		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)

	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			sendLogMessage(writer, "textDocument/definition: "+err.Error(), lsp.Warning)
			return
		}

		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)

	case "textDocument/codeAction":
		var request lsp.CodeActionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			sendLogMessage(writer, "textDocument/codeAction: "+err.Error(), lsp.Warning)
			return
		}

		response := state.TextDocumentCodeAction(request.ID, request.Params.TextDocument.URI)
		writeResponse(writer, response)

	case "textDocument/completion":
		var request lsp.CompletionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			sendLogMessage(writer, "textDocument/codeAction: "+err.Error(), lsp.Warning)
			return
		}

		response := state.TextDocumentCompletion(request.ID, request.Params.TextDocument.URI)
		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply, err := rpc.EncodeMessage(msg)
	if err != nil {
		sendLogMessage(writer, "Error encoding response: "+err.Error(), lsp.Error)
		return
	}

	writer.Write([]byte(reply))
}

func sendLogMessage(writer io.Writer, message string, level int) {
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

func getLogger(filename string) *log.Logger {
	if filename == "" {
		return log.New(os.Stdout, "[plantuml-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
	}

	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[plantuml-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
