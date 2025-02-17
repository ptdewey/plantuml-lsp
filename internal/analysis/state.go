package analysis

import (
	completion "github.com/ptdewey/plantuml-lsp/internal/features"
	"github.com/ptdewey/plantuml-lsp/internal/lsp"
)

// TODO: should these be moved to State?
// - possibly not since GetFeatures is being called asynchronously
var (
	completionItems []lsp.CompletionItem
	hoverResults    map[string]lsp.HoverResult
)

type State struct {
	// Map of file names to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string, execPath []string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text, execPath)
}

func (s *State) UpdateDocument(uri, text string, execPath []string) []lsp.Diagnostic {
	s.Documents[uri] = text

	// PERF: wait timer before running syntax checker -- possibly use channels?
	// - this slows down completion when typing (can also be ignored by calling UpdateDocument in a separate goroutine)
	return getDiagnosticsForFile(text, execPath)
}

// TODO: param options?
func (s *State) GetFeatures(stdlibPath string) {
	completionItems, hoverResults = completion.GetFeatures(stdlibPath)
}
