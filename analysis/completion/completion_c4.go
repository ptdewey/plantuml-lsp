package completion

import (
	"os"
	"path/filepath"
	"plantuml_lsp/lsp"
	"plantuml_lsp/parse"
)

// TODO: add support for other lsp features (i.e. go to definition)
func getC4Items(c4dir string) ([]lsp.CompletionItem, []lsp.HoverResult, error) {
	var completionItems []lsp.CompletionItem
	var hoverResults []lsp.HoverResult

	err := filepath.WalkDir(c4dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ".puml" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			c4items := parse.ExtractC4Items(string(content))
			for _, item := range c4items {
				completionItems = append(completionItems, parse.C4ItemToCompletionItem(item))
				hoverResults = append(hoverResults, parse.C4ItemToHoverResult(item))
			}
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return completionItems, hoverResults, nil
}
