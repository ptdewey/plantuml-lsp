package analysis

import (
	completion "github.com/ptdewey/plantuml-lsp/internal/features"
	"github.com/ptdewey/plantuml-lsp/internal/lsp"
	"github.com/ptdewey/plantuml-lsp/internal/utils/debounce"
)

// TODO: should these be moved to State?
// - possibly not since GetFeatures is being called asynchronously
var (
	completionItems []lsp.CompletionItem
	hoverResults    map[string]lsp.HoverResult
	definitions     map[string]lsp.Location
)

type State struct {
	// Map of file names to contents
	Documents map[string]string

	// Debounce timers for different features (primarily for diagnostics)
	Timers map[string]*debounce.Debouncer
}

func NewState() State {
	return State{
		Documents: map[string]string{},
		Timers:    map[string]*debounce.Debouncer{},
	}
}

func (s *State) OpenDocument(uri, text string, execPath []string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text, execPath)
}

func (s *State) UpdateDocument(uri, text string, execPath []string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text, execPath)
}

// TODO: param options?
func (s *State) GetFeatures(stdlibPath string) {
	completionItems, hoverResults, definitions = completion.GetFeatures(stdlibPath)
}
