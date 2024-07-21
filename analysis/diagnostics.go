package analysis

import "plantuml_lsp/lsp"

// TODO: actual diagnostics
func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	_ = text
	// for row, line := range strings.Split(text, "\n") {
	// 	if strings.Contains(line, "text") {
	// 		idx := strings.Index(line, "text")
	// 		diagnostics = append(diagnostics, lsp.Diagnostic{
	// 			Range:    LineRange(row, idx, idx+len("text")),
	// 			Severity: 2,
	// 			Source:   "plantuml-lsp",
	// 			Message:  "message",
	// 		})
	// 	}
	// }

	return diagnostics
}
