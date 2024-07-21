package completion

import (
	"path/filepath"
	"plantuml_lsp/lsp"
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

	// TODO: call each getter function (or each requested)
	cis, hrs = getC4Items(filepath.Join(stdlibDir, "C4"))
	completionItems = append(completionItems, cis...)
	for k, v := range hrs {
		hoverResults[k] = v
	}

	return completionItems, hoverResults
}
