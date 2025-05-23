package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ptdewey/plantuml-lsp/internal/analysis"
	"github.com/ptdewey/plantuml-lsp/internal/handler"
	"github.com/ptdewey/plantuml-lsp/internal/logger"
	"github.com/ptdewey/plantuml-lsp/internal/lsp"
	"github.com/ptdewey/plantuml-lsp/internal/rpc"
	"github.com/ptdewey/plantuml-lsp/internal/utils"
)

var (
	useStdio   bool
	stdlibPath string
	execCmd    string
	jarPath    string
)

func main() {
	flag.BoolVar(&useStdio, "stdio", true, "Deprecated.")
	flag.StringVar(&stdlibPath, "stdlib-path", "", "PlantUML stdlib path")
	flag.StringVar(&execCmd, "exec-path", "", "PlantUML executable command")
	flag.StringVar(&jarPath, "jar-path", "", "PlantUML jar path")
	flag.Parse()

	var err error
	stdlibPath, err = utils.SanitizePath(stdlibPath)
	if err != nil {
		panic(err)
	}

	var plantumlCmd []string
	if len(execCmd) > 0 {
		cmd := execCmd
		if strings.HasPrefix(execCmd, string(filepath.Separator)) || strings.HasPrefix(execCmd, "~") {
			cmd, err = utils.SanitizePath(execCmd)
			if err != nil {
				panic(err)
			}

			if _, err := os.Stat(cmd); err != nil {
				panic(fmt.Sprintf("Error checking executable path: '%s', Error: %v", cmd, err))
			}
		}
		plantumlCmd = []string{cmd}
	} else if len(jarPath) > 0 {
		jar, err := utils.SanitizePath(jarPath)
		if err != nil {
			panic(err)
		}

		if _, err := os.Stat(jar); err != nil {
			panic(fmt.Sprintf("Error checking executable path: '%s', Error: %v", jar, err))
		}
		plantumlCmd = []string{"java", "-jar", jar}
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	logger.SendLogMessage(writer, "Started plantuml-lsp", lsp.Debug)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.SendLogMessage(writer, "Error decoding message: "+err.Error(), lsp.Error)
			continue
		}

		handler.HandleMessage(writer, state, method, contents, stdlibPath, plantumlCmd)
	}
}
