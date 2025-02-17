package handler_test

import (
	"io"
	"testing"

	"github.com/ptdewey/plantuml-lsp/internal/analysis"
	"github.com/ptdewey/plantuml-lsp/internal/handler"
)

func TestHandleMessage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		writer     io.Writer
		state      analysis.State
		method     string
		contents   []byte
		stdlibPath string
		execPath   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler.HandleMessage(tt.writer, tt.state, tt.method, tt.contents, tt.stdlibPath, tt.execPath)
		})
	}
}
