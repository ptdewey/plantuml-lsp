package completion

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ptdewey/plantuml-lsp/internal/lsp"
	"github.com/ptdewey/plantuml-lsp/internal/parse/stdlib"
)

// TODO: add support for other lsp features (i.e. go to definition)
func getC4Items(c4dir string) ([]lsp.CompletionItem, map[string]lsp.HoverResult, map[string]lsp.Location) {
	var completionItems []lsp.CompletionItem
	hoverResults := make(map[string]lsp.HoverResult)
	definitions := make(map[string]lsp.Location)

	// TODO: move dir walking to parse/stdlib package
	// this function should only set create the feature related items
	err := filepath.WalkDir(c4dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ".puml" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			c4items := stdlib.ExtractC4Items(string(content))
			for _, item := range c4items {
				// add value if it is not a duplicate
				if _, exists := hoverResults[item.Name]; !exists {
					item.SourceFile = path

					completionItems = append(completionItems, item.C4ItemToCompletionItem())
					hoverResults[item.Name] = item.C4ItemToHoverResult()
					definitions[item.Name] = item.C4ItemToLocation()
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return completionItems, hoverResults, definitions
}

// TODO: param to decide link vs path vs builtin include path
func getC4Snippets() []lsp.CompletionItem {
	var completionItems = []lsp.CompletionItem{}

	// TODO:
	// - possibly pull theme definitions from stdlib local location? (depending on passed opts)
	// - pull includes from stdlib local location as well

	// FIX: Replace previously typed text on completion trigger.
	// - Current behavior fills in from cursor position to end of line, doesn't remove any previously typed text
	// - Use "https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#insertReplaceEdit"

	// theme snippets
	themePath := "https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/themes"
	themes := []string{
		"C4_blue", "C4_brown", "C4_green", "C4_sandstone", "C4_superhero", "C4_united", "C4_violet",
	}
	for _, t := range themes {
		completionItems = append(completionItems, lsp.CompletionItem{
			Label:            t,
			Detail:           "Theme",
			Documentation:    "Invoke theme `" + t + "`",
			Kind:             lsp.Snippet,
			InsertText:       "!theme " + t + " from " + themePath,
			InsertTextFormat: lsp.FormatSnippet,
		})
	}

	// include snippets
	includePath := "https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/"
	includes := []string{
		"C4_Component", "C4_Container", "C4_Context", "C4_Deployment", "C4_Dynamic", "C4_Sequence",
	}
	for _, i := range includes {
		completionItems = append(completionItems, lsp.CompletionItem{
			Label:            i,
			Detail:           "Theme",
			Documentation:    "Include `C4/" + i + "`",
			Kind:             lsp.Snippet,
			InsertText:       "!include " + includePath + i + ".puml", // TODO: this would have to change to allow local/builtin include
			InsertTextFormat: lsp.FormatSnippet,
		})
	}

	return completionItems
}
