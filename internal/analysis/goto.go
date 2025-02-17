package analysis

import "github.com/ptdewey/plantuml-lsp/internal/lsp"

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	document := s.Documents[uri]

	currWord := getCurrentWord(document, position)
	_, exists := hoverResults[currWord] // TODO: go to definition

	if !exists {
		return lsp.DefinitionResponse{
			Response: lsp.Response{
				RPC: "2.0",
				ID:  &id,
			},
		}
	}

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}
