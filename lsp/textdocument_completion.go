package lsp

const (
	// InsertTextFormat
	PlainText    = 1 // primary text is inserted as plain text string
	FormatSnippe = 2 // primary text being inserted is treated as snippet
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
	Label               string     `json:"label"`
	Detail              string     `json:"detail"`
	Documentation       string     `json:"documentation"`
	Kind                int        `json:"kind"` // TODO: potentially change this to completionitemkind (convenience)
	InsertText          string     `json:"insertText"`
	InsertTextFormat    int        `json:"insertTextFormat"`
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits"`
}

// CompletionItemKind type
// /**
//   - The kind of a completion entry.
//     */
//	export namespace CompletionItemKind {
//		export const Text = 1;
//		export const Method = 2;
//		export const Function = 3;
//		export const Constructor = 4;
//		export const Field = 5;
//		export const Variable = 6;
//		export const Class = 7;
//		export const Interface = 8;
//		export const Module = 9;
//		export const Property = 10;
//		export const Unit = 11;
//		export const Value = 12;
//		export const Enum = 13;
//		export const Keyword = 14;
//		export const Snippet = 15;
//		export const Color = 16;
//		export const File = 17;
//		export const Reference = 18;
//		export const Folder = 19;
//		export const EnumMember = 20;
//		export const Constant = 21;
//		export const Struct = 22;
//		export const Event = 23;
//		export const Operator = 24;
//		export const TypeParameter = 25;
//	}
