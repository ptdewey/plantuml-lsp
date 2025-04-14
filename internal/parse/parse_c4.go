package parse

import (
	"regexp"
	"strings"

	"github.com/ptdewey/plantuml-lsp/internal/lsp"
)

type C4Item struct {
	Name          string
	Type          string
	Documentation string

	// CompletionItemKind:
	// Text = 1; Method = 2; Function = 3; Constructor = 4; Field = 5; Variable = 6; Class = 7; Interface = 8;
	// Module = 9; Property = 10; Unit = 11; Value = 12; Enum = 13; Keyword = 14; Snippet = 15; Color = 16; File = 17;
	// Reference = 18; Folder = 19; EnumMember = 20; Constant = 21; Struct = 22; Event = 23; Operator = 24; TypeParameter = 25;
	Kind int

	LineNumber int
	SourceFile string
}

// text should be full text of a puml file containing c4 model definitions
func ExtractC4Items(text string) []C4Item {
	var out []C4Item
	var currType string
	var typeBuf []string

	procRe := regexp.MustCompile(`^\s*!(unquoted\s+)?procedure\s+(\w+)\((.*)\)`)

	for i, line := range strings.Split(text, "\n") {
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
			if procMatch := procRe.FindStringSubmatch(line); procMatch != nil {
				out = append(out, C4Item{
					Name:          procMatch[2],
					Type:          currType,
					Documentation: formatDocs(procMatch[2], procMatch[3]),
					Kind:          3, // NOTE: choose between method (2) and function(3)
					LineNumber:    i,
				})
			}
			// TODO: plantuml functions and definitions and other things
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

	// TODO: update documentation with core vs stdlib/lib to make things more clear

	// def := fmt.Sprintf("```puml\n%s(%s)\n```", name, params)
	def := "```rust\n" + name + "(" + params + ")\n```" // NOTE: use this for somewhat working syntax highlights since plantuml doesn't have a parser for most editors

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
				out = append(out, "`"+name+"` (optional)")
			} else {
				out = append(out, "`"+name+"` (optional, default: `"+defaultValue+"`)")
			}
		} else {
			out = append(out, "`"+param+"` (required)")
		}
	}
	return def + "\nParameters: " + strings.Join(out, ", ") +
		"\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)"
}

func (i C4Item) C4ItemToCompletionItem() lsp.CompletionItem {
	return lsp.CompletionItem{
		Label:            i.Name,
		Detail:           i.Type,
		Documentation:    i.Documentation,
		Kind:             i.Kind,
		InsertText:       i.Name + "($0)", // TODO: possibly add params as extra snippet locations
		InsertTextFormat: 2,
	}
}

func (i C4Item) C4ItemToHoverResult() lsp.HoverResult {
	return lsp.HoverResult{
		Contents: i.Documentation,
	}
}

func (i C4Item) C4ItemToLocation() lsp.Location {
	return lsp.Location{
		URI: "file://" + i.SourceFile,
		Range: lsp.Range{
			Start: lsp.Position{
				Line:      i.LineNumber,
				Character: 0,
			},
			End: lsp.Position{
				Line:      i.LineNumber,
				Character: 0,
			},
		},
	}
}
