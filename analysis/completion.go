package analysis

import "plantuml_lsp/lsp"

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	// TODO: always showing result as "Text" and not specified type
	// see CompletionItemKind in docs https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemKind
	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: completionItems,
	}

	return response
}
