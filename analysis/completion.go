package analysis

import "plantuml_lsp/lsp"

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: completionItems,
	}

	return response
}
