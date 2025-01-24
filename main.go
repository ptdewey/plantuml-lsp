package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"plantuml_lsp/analysis"
	"plantuml_lsp/handler"
	"plantuml_lsp/lsp"
	"plantuml_lsp/rpc"
	"strings"
)

func main() {
	// TODO: pass in plantuml_lsp.rc file to use for config stuff
	// - include log level https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#window_logMessage

	logPath := flag.String("log-path", "", "LSP log path")
	stdlibPath := flag.String("stdlib-path", "", "PlantUML stdlib path")
	execCmd := flag.String("exec-path", "", "PlantUML executable command")
	flag.Parse()

	logger := getLogger(*logPath)
	logger.Println("Started plantuml-lsp")

	var execCmdSplit []string
	if len(*execCmd) > 0 {
		execCmdSplit = strings.Split(*execCmd, " ")
		if len(execCmdSplit) > 1 {
			n := len(execCmdSplit) - 1
			execPath := strings.TrimSuffix(execCmdSplit[n], "'")
			if _, err := os.Stat(execPath); err != nil {
				panic(fmt.Sprintf("Error checking executable path: '%s', Error: %v", execPath, err))
			}
			execCmdSplit[0] = strings.TrimPrefix(execCmdSplit[0], "'")
			execCmdSplit[n] = execPath
		} else {
			execPath := execCmdSplit[0]
			if strings.HasPrefix(execPath, "/") {
				if _, err := os.Stat(execPath); err != nil {
					panic(fmt.Sprintf("Error checking executable path: '%s', Error: %v", execPath, err))
				}
			}
		}
	}

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

		handler.HandleMessage(writer, state, method, contents, *stdlibPath, execCmdSplit)
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
