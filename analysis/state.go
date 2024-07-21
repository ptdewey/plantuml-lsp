package analysis

import (
	"plantuml_lsp/features"
	"plantuml_lsp/lsp"
)

var (
	completionItems []lsp.CompletionItem
	hoverResults    []lsp.HoverResult
)

type State struct {
	// Map of file names to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

// TODO: param options?
func (s *State) GetFeatures() error {
	// TODO: don't hardcode directory path
	c, h, err := completion.GetFeatures("/home/patrick/projects/plantuml-stuff/plantuml-stdlib")
	// TODO: possibly switch to log.Fatalf instead of return
	if err != nil {
		return err
	}

	completionItems = c
	hoverResults = h
	return nil
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
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
