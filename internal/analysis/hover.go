package analysis

import (
	"strings"

	"github.com/ptdewey/plantuml-lsp/internal/lsp"
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
	for start > 0 && isValidWordByte(line[start-1]) {
		start--
	}

	end := position.Character
	for end < len(line) && isValidWordByte(line[end]) {
		end++
	}

	return line[start:end]
}

func isValidWordByte(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '_' || b == '@' || b == '!'
}
