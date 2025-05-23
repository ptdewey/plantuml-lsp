package analysis

import (
	"github.com/ptdewey/plantuml-lsp/internal/lsp"
)

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	document := s.Documents[uri]

	currWord := getCurrentWord(document, position)
	location, exists := definitions[currWord]

	if !exists {
		return lsp.DefinitionResponse{
			Response: lsp.Response{
				RPC: "2.0",
				ID:  &id,
			},
			Result: lsp.Location{
				URI: uri,
				Range: lsp.Range{
					Start: position,
					End:   position,
				},
			},
		}
	}

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: location,
	}
}
