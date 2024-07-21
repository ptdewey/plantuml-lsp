package completion

import (
	"fmt"
	"plantuml_lsp/language"
	"plantuml_lsp/lsp"
)

// TODO: param options to not include certain definitions (i.e. skinparams?)
func getCoreItems() ([]lsp.CompletionItem, map[string]lsp.HoverResult) {
	var completionItems []lsp.CompletionItem
	hoverResults := make(map[string]lsp.HoverResult)

	definitionToCompletionItem(language.Types, &completionItems, hoverResults)
	definitionToCompletionItem(language.Keywords, &completionItems, hoverResults)
	definitionToCompletionItem(language.Preprocessors, &completionItems, hoverResults)
	definitionToCompletionItem(language.Arrows, &completionItems, hoverResults)

	// TODO: colors and skinparams (probably make optional)

	return completionItems, hoverResults
}

func definitionToCompletionItem(defs language.LangDefs, completionItems *[]lsp.CompletionItem, hoverResults map[string]lsp.HoverResult) {
	for _, def := range defs.Defs {
		// doc := fmt.Sprintf("core:\n```puml\n%s\n```", def)
		doc := fmt.Sprintf("core:\n```rust\n%s\n```", def) // NOTE: using rust syntax highlighting since puml doesn't have a treesitter parser
		*completionItems = append(*completionItems, lsp.CompletionItem{
			Label:         def,
			Detail:        "", // TODO: Do something with this?
			Documentation: doc,
			Kind:          defs.Kind,
		})
		hoverResults[def] = lsp.HoverResult{
			Contents: doc,
		}
	}
}
