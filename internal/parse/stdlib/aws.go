package stdlib

import "github.com/ptdewey/plantuml-lsp/internal/lsp"

type AWSItem struct {
	Name          string
	Type          string
	Documentation string
	Kind          lsp.CompletionItemKind
	LineNumber    int
	SourceFile    string
}

// TODO:

func ExtractAWSItems(awsDir string) []AWSItem {
	// var out []AWSItem
	// var currType string
	// var typeBuf []string

	// TODO: read from puml file (import it upon completion)
	// - use regex on define
	// - likely include parent dir (posibly multiple levels) in name due to potential overlaps
	//
	// !define MESSAGE(alias) PUML_ENTITY(artifact,#D9A842,message,alias,message)
	//
	// !definelong MESSAGE(alias,label,e_type="artifact",e_color="#D9A842",e_stereo="message",e_sprite="message")
	// PUML_ENTITY(e_type,e_color,e_sprite,label,alias,e_stereo)
	// !enddefinelong

	return nil
}
