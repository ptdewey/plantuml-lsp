package parse

import (
	"fmt"
	"plantuml_lsp/lsp"
	"regexp"
	"strings"
)

type C4Item struct {
	Name          string
	Type          string
	Documentation string
}

// text should be full text of a puml file containing c4 model definitions
func ExtractC4Items(text string) []C4Item {
	var out []C4Item
	var currType string
	var typeBuf []string

	procRe := regexp.MustCompile(`^\s*!(unquoted\s+)?procedure\s+(\w+)\((.*)\)`)

	for _, line := range strings.Split(text, "\n") {
		// end of documentation block
		if strings.HasPrefix(line, "' ##") {
			currType = strings.Join(typeBuf, "\n")
			typeBuf = []string{}
			continue
		}

		// handle doc comments
		if strings.HasPrefix(line, "'") {
			typeBuf = append(typeBuf, strings.TrimSpace(strings.TrimPrefix(line, "'")))
		} else if strings.HasPrefix(line, "!") {
			// handle procedure definitions
			// TODO: store locations for goto definition later
			if procMatch := procRe.FindStringSubmatch(line); procMatch != nil {
				out = append(out, C4Item{
					Name:          procMatch[2],
					Type:          currType,
					Documentation: formatDocs(procMatch[2], procMatch[3]),
				})
			}
			// TODO: plantuml functions
		}

		// reset docs buffer when non-comment line is found
		// TODO: use newly found comments for next single item only
		// - this will mainly be for functions once they are included
		if !strings.HasPrefix(line, "'") && len(typeBuf) > 0 {
			typeBuf = []string{}
		}
	}

	return out
}

func formatDocs(name string, params string) string {
	params = strings.TrimSpace(params)
	// def := fmt.Sprintf("```python\n%s(%s)\n```", name, params) // NOTE: use this for somewhat working syntax highlights since plantuml doesn't have a parser for most editors
	def := fmt.Sprintf("```puml\n%s(%s)\n```", name, params)

	if params == "" {
		return def
	}

	var out = []string{}

	for _, param := range strings.Split(params, ",") {
		param = strings.TrimSpace(param)
		if strings.Contains(param, "=") {
			parts := strings.SplitN(param, "=", 2)
			name := strings.TrimSpace(parts[0])
			defaultValue := strings.TrimSpace(parts[1])
			if defaultValue == `""` {
				out = append(out, fmt.Sprintf("%s (optional)", name))
			} else {
				out = append(out, fmt.Sprintf("%s (optional, default: %s)", name, defaultValue))
			}
		} else {
			out = append(out, fmt.Sprintf("%s (required)", param))
		}
	}
	return fmt.Sprintf("%s\n\nParameters: %s", def, strings.Join(out, ", "))
}

func C4ItemToCompletionItem(i C4Item) lsp.CompletionItem {
	return lsp.CompletionItem{
		Label:         i.Name,
		Detail:        i.Type,
		Documentation: i.Documentation,
	}
}

func C4ItemToHoverResult(i C4Item) lsp.HoverResult {
	return lsp.HoverResult{
		Contents: i.Documentation,
	}
}
