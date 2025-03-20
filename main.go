package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ptdewey/plantuml-lsp/internal/analysis"
	"github.com/ptdewey/plantuml-lsp/internal/handler"
	"github.com/ptdewey/plantuml-lsp/internal/lsp"
	"github.com/ptdewey/plantuml-lsp/internal/rpc"
)

func main() {
	// TODO: pass in plantuml_lsp.rc file to use for config stuff
	// - include log level https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#window_logMessage

	useStdio := flag.Bool("stdio", false, "")
	logPath := flag.String("log-path", "", "LSP log path")
	stdlibPath := flag.String("stdlib-path", "", "PlantUML stdlib path")
	execCmd := flag.String("exec-path", "", "PlantUML executable command")
	jarPath := flag.String("jar-path", "", "PlantUML jar path")
	flag.Parse()

	logger := getLogger(*useStdio, *logPath)
	logger.Println("Started plantuml-lsp")

	var plantumlCmd []string
	if len(*execCmd) > 0 {
		if strings.HasPrefix(*execCmd, "/") {
			if _, err := os.Stat(*execCmd); err != nil {
				panic(fmt.Sprintf("Error checking executable path: '%s', Error: %v", *execCmd, err))
			}
		}
		plantumlCmd = []string{*execCmd}
	} else if len(*jarPath) > 0 {
		if _, err := os.Stat(*jarPath); err != nil {
			panic(fmt.Sprintf("Error checking executable path: '%s', Error: %v", *jarPath, err))
		}
		plantumlCmd = []string{"java", "-jar", *jarPath}
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

		handler.HandleMessage(writer, state, method, contents, *stdlibPath, plantumlCmd)
	}
}

func getLogger(useStdio bool, filename string) *log.Logger {
	if useStdio {
		return log.New(os.Stderr, "[plantuml-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
	}

	if filename == "" {
		return log.New(os.Stdout, "[plantuml-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
	}

	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[plantuml-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
