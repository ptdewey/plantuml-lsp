package completion

import (
	"log"
	"path/filepath"
	"plantuml_lsp/lsp"
)

// TODO: finish implementation
func GetFeatures(stdlibDir string) ([]lsp.CompletionItem, []lsp.HoverResult, error) {
	var completionItems []lsp.CompletionItem
	var hoverResults []lsp.HoverResult

	// TODO: call each getter function (or each requested)
	c4cis, c4hrs, err := getC4Items(filepath.Join(stdlibDir, "C4"))
	if err != nil {
		log.Println(err)
	}
	completionItems = append(completionItems, c4cis...)
	hoverResults = append(hoverResults, c4hrs...)

	ccis, chrs := getCoreItems()
	completionItems = append(completionItems, ccis...)
	hoverResults = append(hoverResults, chrs...)

	return completionItems, hoverResults, nil
}
