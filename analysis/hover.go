package analysis

import (
	"plantuml_lsp/lsp"
	"strings"
)

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Documents[uri]

	currWord := getCurrentWord(document, position)
	hoverResult, exists := hoverResults[currWord]

	if !exists {
		return lsp.HoverResponse{
			Response: lsp.Response{
				RPC: "2.0",
				ID:  &id,
			},
		}
	}

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: hoverResult.Contents,
		},
	}
}

func getCurrentWord(content string, position lsp.Position) string {
	lines := strings.Split(content, "\n")
	if position.Line >= len(lines) {
		return ""
	}
	line := lines[position.Line]

	start := position.Character
	for start > 0 {
		start--
	}

	end := position.Character
	for end < len(line) {
		end++
	}

	return line[start:end]
}
