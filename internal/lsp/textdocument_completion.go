package lsp

type CompletionRequest struct {
	Request
	Params CompletionParams `json:"params"`
}

type CompletionParams struct {
	TextDocumentPositionParams
}

type CompletionResponse struct {
	Response
	Result CompletionList `json:"result"`
}

type InsertTextFormat int

const (
	_ InsertTextFormat = iota
	PlainText
	FormatSnippet
)

type CompletionItem struct {
	Label               string             `json:"label"`
	Detail              string             `json:"detail"`
	FilterText          *string            `json:"filterText,omitempty"`
	Documentation       string             `json:"documentation"`
	Kind                CompletionItemKind `json:"kind"`
	InsertText          string             `json:"insertText"`
	TextEdit            *TextEdit          `json:"textEdit,omitempty"`
	InsertTextFormat    InsertTextFormat   `json:"insertTextFormat"`
	AdditionalTextEdits []TextEdit         `json:"additionalTextEdits"`
}

type CompletionItemKind int

const (
	_ CompletionItemKind = iota
	Text
	Method
	Function
	Constructor
	Field
	Variable
	Class
	Interface
	Module
	Property
	Unit
	Value
	Enum
	Keyword
	Snippet
	Color
	File
	Reference
	Folder
	EnumMember
	Constant
	Struct
	Event
	Operator
	TypeParameter
	Unknown
)

type CompletionList struct {
	IsIncomplete bool `json:"isIncomplete"`
	// TODO: ItemDefaults *ItemDefaults `json:"itemDefaults,omitempty"`
	Items []CompletionItem `json:"items"`
}
