package completion

import (
	"fmt"
	"plantuml_lsp/lsp"
)

// TODO: finish implementation
func GetFeatures(stdlibDir string) ([]lsp.CompletionItem, []lsp.HoverResult, error) {
	var completionItems []lsp.CompletionItem
	var hoverResults []lsp.HoverResult

	// TODO: call each getter function (or each requested)
	c4cis, c4hrs, err := GetC4Items(fmt.Sprintf("%s/%s", stdlibDir, "C4")) // TODO: switch to filepath.Join instead of sprintf
	if err != nil {
		return nil, nil, err
	}
	completionItems = append(completionItems, c4cis...)
	hoverResults = append(hoverResults, c4hrs...)

	return completionItems, hoverResults, nil
}
