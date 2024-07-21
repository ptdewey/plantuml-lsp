package analysis

import (
	"plantuml_lsp/features"
	"plantuml_lsp/lsp"
)

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

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

// TODO: param options?
// - pass in std lib directory path
func (s *State) GetFeatures() {
	completionItems, hoverResults = completion.GetFeatures("/home/patrick/projects/plantuml-stuff/plantuml-stdlib")
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}
