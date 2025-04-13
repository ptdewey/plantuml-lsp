package completion

import (
	"os"
	"path/filepath"

	"github.com/ptdewey/plantuml-lsp/internal/lsp"
)

// TODO: finish implementation
func GetFeatures(stdlibDir string) ([]lsp.CompletionItem, map[string]lsp.HoverResult) {
	var completionItems []lsp.CompletionItem
	hoverResults := make(map[string]lsp.HoverResult)

	cis, hrs := getCoreItems()
	completionItems = append(completionItems, cis...)
	for k, v := range hrs {
		hoverResults[k] = v
	}
	completionItems = append(completionItems, getCoreSnippets()...)

	c4path := ""
	if _, err := os.Stat(filepath.Join(stdlibDir, "C4")); err == nil {
		c4path = filepath.Join(stdlibDir, "C4")
	} else if _, err := os.Stat(filepath.Join(stdlibDir, "c4")); err == nil {
		c4path = filepath.Join(stdlibDir, "c4")
	}
	cis, hrs = getC4Items(c4path)
	completionItems = append(completionItems, cis...)
	for k, v := range hrs {
		hoverResults[k] = v
	}
	completionItems = append(completionItems, getC4Snippets()...)

	// TODO: Add other stdlib completions (available as toggleable features defined in server configuration)

	return completionItems, hoverResults
}
