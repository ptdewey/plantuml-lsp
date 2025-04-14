package parse_test

import (
	// "fmt"
	// "log"
	// "os"
	"reflect"
	"testing"

	"github.com/ptdewey/plantuml-lsp/internal/lsp"
	"github.com/ptdewey/plantuml-lsp/internal/parse"
)

func TestExtractC4Items(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []parse.C4Item
	}{
		{
			name: "single procedure with documentation",
			text: `
' This is a procedure.
' ##################
!procedure ExampleProc(param1, param2 = "default")
`,
			want: []parse.C4Item{
				{
					Name:          "ExampleProc",
					Type:          "This is a procedure.",
					Documentation: "```rust\nExampleProc(param1, param2 = \"default\")\n```\nParameters: `param1` (required), `param2` (optional, default: `\"default\"`)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
					Kind:          3,
					LineNumber:    3,
					SourceFile:    "",
				},
			},
		},
		{
			name: "optional parameter with empty default value",
			text: `
' This is a procedure with empty default.
' ##################
!procedure ExampleProcWithEmptyDefault(param1, param2 = "")
`,
			want: []parse.C4Item{
				{
					Name:          "ExampleProcWithEmptyDefault",
					Type:          "This is a procedure with empty default.",
					Documentation: "```rust\nExampleProcWithEmptyDefault(param1, param2 = \"\")\n```\nParameters: `param1` (required), `param2` (optional)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
					Kind:          3,
					LineNumber:    3,
					SourceFile:    "",
				},
			},
		},
		{
			name: "multiple procedures with shared documentation",
			text: `
' This is the first procedure.
' This is the second procedure.
' ###################
!procedure FirstProc(param1)
!procedure SecondProc(param2 = 42)
`,
			want: []parse.C4Item{
				{
					Name:          "FirstProc",
					Type:          "This is the first procedure.\nThis is the second procedure.",
					Documentation: "```rust\nFirstProc(param1)\n```\nParameters: `param1` (required)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
					Kind:          3,
					LineNumber:    4,
					SourceFile:    "",
				},
				{
					Name:          "SecondProc",
					Type:          "This is the first procedure.\nThis is the second procedure.",
					Documentation: "```rust\nSecondProc(param2 = 42)\n```\nParameters: `param2` (optional, default: `42`)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
					Kind:          3,
					LineNumber:    5,
					SourceFile:    "",
				},
			},
		},
		{
			name: "no documentation",
			text: `
!procedure NoDocsProc(param1, param2)
`,
			want: []parse.C4Item{
				{
					Name:          "NoDocsProc",
					Type:          "",
					Documentation: "```rust\nNoDocsProc(param1, param2)\n```\nParameters: `param1` (required), `param2` (required)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
					Kind:          3,
					LineNumber:    1,
					SourceFile:    "",
				},
			},
		},
		{
			name: "mixed comments and procedures",
			text: `
' Docs for FirstProc.
' ## Block 1
!procedure FirstProc(param1)
' Docs for SecondProc.
' Docs for SecondProc continued.
' ## Block 2
!procedure SecondProc(param2 = 42)
`,
			want: []parse.C4Item{
				{
					Name:          "FirstProc",
					Type:          "Docs for FirstProc.",
					Documentation: "```rust\nFirstProc(param1)\n```\nParameters: `param1` (required)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
					Kind:          3,
					LineNumber:    3,
					SourceFile:    "",
				},
				{
					Name:          "SecondProc",
					Type:          "Docs for SecondProc.\nDocs for SecondProc continued.",
					Documentation: "```rust\nSecondProc(param2 = 42)\n```\nParameters: `param2` (optional, default: `42`)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
					Kind:          3,
					LineNumber:    7,
					SourceFile:    "",
				},
			},
		},
		{
			name: "mixed comments, procedures, and non-comments",
			text: `
' Docs for non-comment line.
hide footbox
' This is a procedure.
' ##################
!procedure ExampleProc(param1, param2 = "default")
`,
			want: []parse.C4Item{
				{
					Name:          "ExampleProc",
					Type:          "This is a procedure.",
					Documentation: "```rust\nExampleProc(param1, param2 = \"default\")\n```\nParameters: `param1` (required), `param2` (optional, default: `\"default\"`)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
					Kind:          3,
					LineNumber:    5,
					SourceFile:    "",
				},
			},
		},
		{
			name: "no parameters",
			text: `
' This is a procedure.
' ##################
!procedure ExampleProc()
`,
			want: []parse.C4Item{
				{
					Name:          "ExampleProc",
					Type:          "This is a procedure.",
					Documentation: "```rust\nExampleProc()\n```",
					Kind:          3,
					LineNumber:    3,
					SourceFile:    "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse.ExtractC4Items(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractC4Items() = got:\n %v\n want:\n %v", got, tt.want)
			}
		})
	}
}

func TestC4ItemToCompletionItem(t *testing.T) {
	item := parse.C4Item{
		Name:          "ExampleProc",
		Type:          "This is a procedure.",
		Documentation: "```rust\nExampleProc(param1, param2 = \"default\")\n```\nParameters: `param1` (required), `param2` (optional, default: `\"default\"`)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
		Kind:          3,
		LineNumber:    3,
		SourceFile:    "",
	}
	want := lsp.CompletionItem{
		Label:            "ExampleProc",
		Detail:           "This is a procedure.",
		Documentation:    "```rust\nExampleProc(param1, param2 = \"default\")\n```\nParameters: `param1` (required), `param2` (optional, default: `\"default\"`)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
		Kind:             3,
		InsertText:       "ExampleProc($0)",
		InsertTextFormat: 2,
	}
	got := item.C4ItemToCompletionItem()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected result from C4ItemToCompletionItem()\n\tgot:  %v\n\twant: %v", got, want)
	}
}

func TestC4ItemToHoverResult(t *testing.T) {
	item := parse.C4Item{
		Name:          "ExampleProc",
		Type:          "This is a procedure.",
		Documentation: "```rust\nExampleProc(param1, param2 = \"default\")\n```\nParameters: `param1` (required), `param2` (optional, default: `\"default\"`)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
		Kind:          3,
		LineNumber:    5,
		SourceFile:    "",
	}
	want := lsp.HoverResult{
		Contents: "```rust\nExampleProc(param1, param2 = \"default\")\n```\nParameters: `param1` (required), `param2` (optional, default: `\"default\"`)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
	}
	got := item.C4ItemToHoverResult()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected result from C4ItemToHoverResult()\n\tgot:  %v\n\twant: %v", got, want)
	}
}
