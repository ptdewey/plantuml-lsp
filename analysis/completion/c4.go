package stdlib

import (
	"os"
	"path/filepath"
	"plantuml_lsp/lsp"
	"plantuml_lsp/parse"
)

func GenC4Completions(c4dir string) error {
	var items []parse.C4Item

	err := filepath.WalkDir(c4dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ".puml" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			items = append(items, parse.ExtractC4Items(string(content))...)
		}

		return nil
	})
	if err != nil {
		return err
	}

	// TODO: convert to completionitem slice
	lsp.CompletionItem

	return nil
}
