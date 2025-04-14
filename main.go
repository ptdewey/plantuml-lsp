package main

import (
	"bufio"
	"flag"
	"fmt"
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

	stdlibPath := flag.String("stdlib-path", "", "PlantUML stdlib path")
	execCmd := flag.String("exec-path", "", "PlantUML executable command")
	jarPath := flag.String("jar-path", "", "PlantUML jar path")
	flag.Parse()

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

	handler.SendLogMessage(writer, "Started plantuml-lsp", lsp.Debug)

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
