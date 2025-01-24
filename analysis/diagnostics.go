package analysis

import (
	"bytes"
	"os/exec"
	"plantuml_lsp/lsp"
	"strconv"
	"strings"
)

func getDiagnosticsForFile(text string, plantUmlExecPath []string) []lsp.Diagnostic {
	if len(plantUmlExecPath) == 0 {
		return []lsp.Diagnostic{}
	}
	plantUmlDiagnostics := getPlantUmlDiagnostics(text, plantUmlExecPath)
	splitText := strings.Split(text, "\n")
	return parsePlantUmlDiagnostics(plantUmlDiagnostics, splitText)
}

func getPlantUmlDiagnostics(text string, plantUmlExecPath []string) string {
	plantUmlExecPath = append(plantUmlExecPath, "-syntax")
	plantumlCmd := exec.Command(plantUmlExecPath[0], plantUmlExecPath[1:]...)
	plantumlCmd.Stdin = bytes.NewReader([]byte(text))
	output, _ := plantumlCmd.CombinedOutput()
	return string(output)
}

func parsePlantUmlDiagnostics(plantUmlDiagnostics string, text []string) []lsp.Diagnostic {
	diagnosticsStrings := strings.Split(plantUmlDiagnostics, "\n")
	if len(diagnosticsStrings) < 3 || diagnosticsStrings[0] != "ERROR" {
		return []lsp.Diagnostic{}
	}

	lineNumber, err := strconv.Atoi(diagnosticsStrings[1])
	if err != nil {
		//TODO: propagate diagnostics parse failure to logger(s)
		return []lsp.Diagnostic{}
	}

	lineLength := len(text[lineNumber])

	parsedDiagnostics := []lsp.Diagnostic{}
	parsedDiagnostics = append(parsedDiagnostics, lsp.Diagnostic{
		Range:    lineRange(lineNumber, 0, lineLength),
		Severity: 1,
		Source:   "plantuml-lsp",
		Message:  diagnosticsStrings[2],
	})
	return parsedDiagnostics
}
