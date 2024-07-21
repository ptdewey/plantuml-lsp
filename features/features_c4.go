package completion

import (
	"log"
	"os"
	"path/filepath"
	"plantuml_lsp/lsp"
	"plantuml_lsp/parse"
)

// TODO: add support for other lsp features (i.e. go to definition)
func getC4Items(c4dir string) ([]lsp.CompletionItem, map[string]lsp.HoverResult) {
	var completionItems []lsp.CompletionItem
	hoverResults := make(map[string]lsp.HoverResult)

	err := filepath.WalkDir(c4dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// TODO: don't add duplicate definitions (defined in multiple files)
		// - probably use a map to check
		if !d.IsDir() && filepath.Ext(d.Name()) == ".puml" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			c4items := parse.ExtractC4Items(string(content))
			for _, item := range c4items {
				// add value if it is not a duplicate
				if _, exists := hoverResults[item.Name]; !exists {
					completionItems = append(completionItems, item.C4ItemToCompletionItem())
					hoverResults[item.Name] = item.C4ItemToHoverResult()
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return completionItems, hoverResults
}
