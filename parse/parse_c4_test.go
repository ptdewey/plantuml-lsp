package parse_test

import (
	"fmt"
	"log"
	"os"
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
					Type:          "procedure",
					Documentation: "This is a procedure.",
					Parameters: []parse.Parameter{
						{Name: "param1", Optional: false},
						{Name: "param2", Optional: true, DefaultValue: `"default"`},
					},
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
					Type:          "procedure",
					Documentation: "This is the first procedure.\nThis is the second procedure.",
					Parameters: []parse.Parameter{
						{Name: "param1", Optional: false},
					},
				},
				{
					Name:          "SecondProc",
					Type:          "procedure",
					Documentation: "This is the first procedure.\nThis is the second procedure.",
					Parameters: []parse.Parameter{
						{Name: "param2", Optional: true, DefaultValue: "42"},
					},
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
					Type:          "procedure",
					Documentation: "",
					Parameters: []parse.Parameter{
						{Name: "param1", Optional: false},
						{Name: "param2", Optional: false},
					},
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
					Type:          "procedure",
					Documentation: "Docs for FirstProc.",
					Parameters: []parse.Parameter{
						{Name: "param1", Optional: false},
					},
				},
				{
					Name:          "SecondProc",
					Type:          "procedure",
					Documentation: "Docs for SecondProc.\nDocs for SecondProc continued.",
					Parameters: []parse.Parameter{
						{Name: "param2", Optional: true, DefaultValue: "42"},
					},
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

// Test function to read in a PlantUML file and print the output
func TestExtractC4ItemsFromFile(t *testing.T) {
	filePath := "/home/patrick/projects/plantuml-stuff/plantuml-stdlib/C4/C4_Container.puml"
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	c4Items := parse.ExtractC4Items(string(content))
	for _, item := range c4Items {
		fmt.Printf("Name: %s\n", item.Name)
		fmt.Printf("Type: %s\n", item.Type)
		fmt.Printf("Documentation: %s\n", item.Documentation)
		fmt.Println("Parameters:")
		for _, param := range item.Parameters {
			optionality := "Required"
			if param.Optional {
				optionality = "Optional"
			}
			defaultValue := ""
			if param.DefaultValue != `""` && param.DefaultValue != "" {
				defaultValue = fmt.Sprintf(" (default: %s)", param.DefaultValue)
			}
			fmt.Printf("  - %s: %s%s\n", param.Name, optionality, defaultValue)
		}
		fmt.Println()
	}
}
