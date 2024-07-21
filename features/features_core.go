package completion

import (
	"fmt"
	"plantuml_lsp/language"
	"plantuml_lsp/lsp"
)

/**
 * CompletionItemKind:
 * Text = 1; Method = 2; Function = 3; Constructor = 4; Field = 5; Variable = 6; Class = 7; Interface = 8;
 * Module = 9; Property = 10; Unit = 11; Value = 12; Enum = 13; Keyword = 14; Snippet = 15; Color = 16; File = 17;
 * Reference = 18; Folder = 19; EnumMember = 20; Constant = 21; Struct = 22; Event = 23; Operator = 24; TypeParameter = 25;
 */

// TODO: param options to not include certain definitions (i.e. skinparams?)
func getCoreItems() ([]lsp.CompletionItem, []lsp.HoverResult) {
	var completionItems []lsp.CompletionItem
	var hoverResults []lsp.HoverResult

	// TODO: improve documentation strings
	for _, t := range language.Types {
		doc := fmt.Sprintf("core:\n```puml\n%s\n```", t)
		completionItems = append(completionItems, lsp.CompletionItem{
			Label:         t,
			Detail:        "",
			Documentation: doc,
			Kind:          7, // FIX: the kinds here may be more complicated than the others, define some as enum, interface, operator, etc.
		})
		hoverResults = append(hoverResults, lsp.HoverResult{
			Contents: doc,
		})
	}
	for _, k := range language.Keywords {
		doc := fmt.Sprintf("core:\n```puml\n%s\n```", k)
		completionItems = append(completionItems, lsp.CompletionItem{
			Label:         k,
			Detail:        "",
			Documentation: doc,
			Kind:          14,
		})
		hoverResults = append(hoverResults, lsp.HoverResult{
			Contents: doc,
		})
	}
	for _, p := range language.Preprocessors {
		doc := fmt.Sprintf("core:\n```puml\n%s\n```", p)
		completionItems = append(completionItems, lsp.CompletionItem{
			Label:         p,
			Detail:        "",
			Documentation: doc,
			Kind:          14,
		})
		hoverResults = append(hoverResults, lsp.HoverResult{
			Contents: doc,
		})
	}
	for _, a := range language.Arrows {
		doc := fmt.Sprintf("core:\n```puml\n%s\n```", a)
		completionItems = append(completionItems, lsp.CompletionItem{
			Label:         a,
			Detail:        "",
			Documentation: doc,
			Kind:          24,
		})
		hoverResults = append(hoverResults, lsp.HoverResult{
			Contents: doc,
		})
	}
	// TODO: colors and skinparams (probably make optional)

	return completionItems, hoverResults
}
