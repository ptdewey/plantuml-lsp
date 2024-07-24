package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"plantuml_lsp/analysis"
	"plantuml_lsp/handler"
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
			handler.SendLogMessage(writer, "Error decoding message: "+err.Error(), lsp.Error)
			continue
		}

		handler.HandleMessage(writer, state, method, contents, *stdlibPath)
	}
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
