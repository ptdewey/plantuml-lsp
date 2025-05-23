package completion

import (
	"github.com/ptdewey/plantuml-lsp/internal/language"
	"github.com/ptdewey/plantuml-lsp/internal/lsp"
)

func getCoreItems() ([]lsp.CompletionItem, map[string]lsp.HoverResult) {
	var completionItems []lsp.CompletionItem
	hoverResults := make(map[string]lsp.HoverResult)

	definitionToCompletionItem(language.Types, &completionItems, hoverResults)
	definitionToCompletionItem(language.Keywords, &completionItems, hoverResults)
	definitionToCompletionItem(language.Preprocessors, &completionItems, hoverResults)
	definitionToCompletionItem(language.Arrows, &completionItems, hoverResults)
	definitionToCompletionItem(language.Colors, &completionItems, hoverResults)
	definitionToCompletionItem(language.Skinparameters, &completionItems, hoverResults)

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
			InsertTextFormat: lsp.PlainText,
		})
		hoverResults[def] = lsp.HoverResult{
			Contents: doc,
		}
	}
}

func getCoreSnippets() []lsp.CompletionItem {
	var completionItems = []lsp.CompletionItem{
		{
			Label:         "startuml",
			Detail:        "Insert start/end UML Tags",
			Documentation: "",
			Kind:          lsp.Snippet,
			// REFACTOR: determine optimal snippet behavior for startuml
			// - behaves differently in vscode with InsertText
			InsertText: "@startuml\n",
			// TextEdit: &lsp.TextEdit{
			// 	NewText: "@startuml\n",
			// 	Range: lsp.Range{
			// 		Start: lsp.Position{
			// 			Line:      0,
			// 			Character: 0,
			// 		},
			// 		End: lsp.Position{
			// 			Line:      0,
			// 			Character: 7,
			// 		},
			// 	},
			// },
			InsertTextFormat: lsp.PlainText,
			AdditionalTextEdits: []lsp.TextEdit{
				{
					NewText: "\n@enduml",
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
				},
			},
		},
	}

	return completionItems
}
