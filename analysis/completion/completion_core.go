package completion

import (
	"plantuml_lsp/lsp"
)

/**
 * CompletionItemKind:
 * Text = 1; Method = 2; Function = 3; Constructor = 4; Field = 5; Variable = 6; Class = 7; Interface = 8;
 * Module = 9; Property = 10; Unit = 11; Value = 12; Enum = 13; Keyword = 14; Snippet = 15; Color = 16; File = 17;
 * Reference = 18; Folder = 19; EnumMember = 20; Constant = 21; Struct = 22; Event = 23; Operator = 24; TypeParameter = 25;
 */

// TODO: figure out a way to strip definitions out of the core code (or its documentation)
// see this repo/dir for info https://github.com/plantuml/plantuml.js/blob/main/plantuml-docs/docs_from_alphadoc/developers.adoc
func getCoreItems() ([]lsp.CompletionItem, []lsp.HoverResult, error) {
	var completionItems []lsp.CompletionItem
	var hoverResults []lsp.HoverResult

	// manually defined definitions (move to their own type to also be used with hover)
	completionItems = append(completionItems, []lsp.CompletionItem{
		{
			Label:         "@startuml",
			Detail:        "Start PlantUML Diagram.",
			Documentation: "```puml\n@startuml\n```\nStart PlantUML Diagram.",
			Kind:          14,
		},
		{
			Label:         "@enduml",
			Detail:        "End PlantUML Diagram.",
			Documentation: "```puml\n@enduml\n```\nEnd PlantUML Diagram.",
			Kind:          14,
		},
		{
			Label:         "rectangle",
			Detail:        "PlantUML rectangle component",
			Documentation: "```puml\nrectangle {name} as {tag}\n```\n", // TODO: better documentation
			Kind:          7,
		},
	}...)

	return completionItems, hoverResults, nil
}
