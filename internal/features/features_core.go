package completion

import (
	"github.com/ptdewey/plantuml-lsp/internal/language"
	"github.com/ptdewey/plantuml-lsp/internal/lsp"
)

// TODO: param options to not include certain definitions (i.e. skinparams?)
func getCoreItems() ([]lsp.CompletionItem, map[string]lsp.HoverResult) {
	var completionItems []lsp.CompletionItem
	hoverResults := make(map[string]lsp.HoverResult)

	definitionToCompletionItem(language.Types, &completionItems, hoverResults)
	definitionToCompletionItem(language.Keywords, &completionItems, hoverResults)
	definitionToCompletionItem(language.Preprocessors, &completionItems, hoverResults)
	definitionToCompletionItem(language.Arrows, &completionItems, hoverResults)
	definitionToCompletionItem(language.Colors, &completionItems, hoverResults)
	definitionToCompletionItem(language.Skinparameters, &completionItems, hoverResults)

	// TODO: colors and skinparams (probably make optional)

	return completionItems, hoverResults
}

func definitionToCompletionItem(defs language.LangDefs, completionItems *[]lsp.CompletionItem, hoverResults map[string]lsp.HoverResult) {
	for _, def := range defs.Defs {
		doc := "```puml\n" + def + "\n```\n[`plantuml/core`](https://github.com/plantuml/plantuml)"
		// TODO: currently, @, !, - completions are not being shown initially
		*completionItems = append(*completionItems, lsp.CompletionItem{
			Label:            def,
			Detail:           defs.Type,
			Documentation:    doc,
			Kind:             defs.Kind,
			InsertText:       def,
			InsertTextFormat: 1,
		})
		hoverResults[def] = lsp.HoverResult{
			Contents: doc,
		}
	}
}

func getCoreSnippets() []lsp.CompletionItem {
	var completionItems = []lsp.CompletionItem{
		{
			Label:            "startuml",
			Detail:           "Insert start/end UML Tags",
			Documentation:    "",
			Kind:             15,
			InsertText:       "@startuml\n\n$0", // TODO: figure out how to make enduml insert at end of file
			InsertTextFormat: 2,
			AdditionalTextEdits: []lsp.TextEdit{
				{
					Range: lsp.Range{
						Start: lsp.Position{
							Line:      999999, // using a long location to avoid needing file length at setup time
							Character: 0,
						},
						End: lsp.Position{
							Line:      999999,
							Character: 0,
						},
					},
					NewText: "\n@enduml",
				},
			},
		},
	}

	return completionItems
}
