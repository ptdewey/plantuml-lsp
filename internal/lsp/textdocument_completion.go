package lsp

const (
	// InsertTextFormat
	PlainText     = 1 // primary text is inserted as plain text string
	FormatSnippet = 2 // primary text being inserted is treated as snippet
)

type CompletionRequest struct {
	Request
	Params CompletionParams `json:"params"`
}

type CompletionParams struct {
	TextDocumentPositionParams
}

type CompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

type CompletionItem struct {
	Label               string             `json:"label"`
	Detail              string             `json:"detail"`
	Documentation       string             `json:"documentation"`
	Kind                CompletionItemKind `json:"kind"`
	InsertText          string             `json:"insertText"`
	TextEdit            *TextEdit          `json:"textEdit"`
	InsertTextFormat    int                `json:"insertTextFormat"`
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
