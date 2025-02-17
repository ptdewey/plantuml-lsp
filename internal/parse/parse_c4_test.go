package parse_test

import (
	// "fmt"
	// "log"
	// "os"
	"plantuml_lsp/parse"
	"reflect"
	"testing"
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
				},
				{
					Name:          "SecondProc",
					Type:          "This is the first procedure.\nThis is the second procedure.",
					Documentation: "```rust\nSecondProc(param2 = 42)\n```\nParameters: `param2` (optional, default: `42`)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
					Kind:          3,
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
				},
				{
					Name:          "SecondProc",
					Type:          "Docs for SecondProc.\nDocs for SecondProc continued.",
					Documentation: "```rust\nSecondProc(param2 = 42)\n```\nParameters: `param2` (optional, default: `42`)\n\n[`stdlib/C4`](https://github.com/plantuml/plantuml-stdlib/tree/master/C4)",
					Kind:          3,
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
