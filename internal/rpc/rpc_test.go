package rpc_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/ptdewey/plantuml-lsp/internal/rpc"
)

type EncodingExample struct {
	Testing bool
}

// FailingEncodingExample is designed to fail JSON encoding.
// It's intended to use in test cases where it's necessary to simulate a JSON encoding error.
type FailingEncodingExample struct{}

func (e FailingEncodingExample) MarshalJSON() ([]byte, error) {
	return nil, errors.New("intentional error")
}

func TestEncodeMessage(t *testing.T) {
	tests := []struct {
		name string
		msg  any
		want string
		err  error
	}{
		{
			name: "message",
			msg:  EncodingExample{Testing: true},
			want: "Content-Length: 16\r\n\r\n{\"Testing\":true}",
			err:  nil,
		},
		{
			name: "message with marshaling error",
			msg:  FailingEncodingExample{},
			want: "",
			err:  errors.New("json: error calling MarshalJSON for type rpc_test.FailingEncodingExample: intentional error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rpc.EncodeMessage(tt.msg)
			if !equalErrors(t, err, tt.err) {
				t.Fatalf("Unexpected error:\n\tgot:  %v\n\twant: %v", err, tt.err)
			}

			if got != tt.want {
				t.Errorf("Invalid message encoding:\n\tgot:  %s\n\twant: %s", got, tt.want)
			}
		})
	}
}

func TestDecodeMessage(t *testing.T) {
	tests := []struct {
		name    string
		msg     []byte
		method  string
		content []byte
		err     error
	}{
		{
			name:    "message",
			msg:     []byte("Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"),
			method:  "hi",
			content: []byte("{\"Method\":\"hi\"}"),
			err:     nil,
		},
		{
			name:    "message without separator",
			msg:     []byte("Content-Length: 15 {\"Method\":\"hi\"}"),
			method:  "",
			content: nil,
			err:     errors.New("Separator not found"),
		},
		{
			name:    "message with invalid separator",
			msg:     []byte("Content-Length: 15\r\n{\"Method\":\"hi\"}"),
			method:  "",
			content: nil,
			err:     errors.New("Separator not found"),
		},
		{
			name:    "message with invalid content length",
			msg:     []byte("Content-Length: --\r\n\r\n{\"Method\":\"hi\"}"),
			method:  "",
			content: nil,
			err:     errors.New("strconv.Atoi: parsing \"--\": invalid syntax"),
		},
		{
			name:    "message with invalid JSON",
			msg:     []byte("Content-Length: 15\r\n\r\n{\"Method\":\"hi\"x_x"),
			method:  "",
			content: nil,
			err:     errors.New("invalid character 'x' after object key:value pair"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method, content, err := rpc.DecodeMessage(tt.msg)
			if !equalErrors(t, err, tt.err) {
				t.Fatalf("Unexpected error:\n\tgot:  %v\n\twant: %v", err, tt.err)
			}

			if method != tt.method {
				t.Errorf("Invalid method decoding:\n\tgot:  %s\n\twant: %s", method, tt.method)
			}

			if !bytes.Equal(content, tt.content) {
				t.Errorf("Invalid content decoding:\n\tgot:  %s\n\twant: %s", content, tt.content)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		atEOF   bool
		length  int
		content []byte
		err     error
	}{
		{
			name:    "message",
			data:    []byte("Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"),
			atEOF:   false,
			length:  37,
			content: []byte("Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"),
			err:     nil,
		},
		{
			name:    "message without separator",
			data:    []byte("Content-Length: 15 {\"Method\":\"hi\"}"),
			atEOF:   false,
			length:  0,
			content: nil,
			err:     errors.New("Separator not found"),
		},
		{
			name:    "message with invalid separator",
			data:    []byte("Content-Length: 15\r\n{\"Method\":\"hi\"}"),
			atEOF:   false,
			length:  0,
			content: nil,
			err:     errors.New("Separator not found"),
		},
		{
			name:    "message with invalid content length",
			data:    []byte("Content-Length: --\r\n\r\n{\"Method\":\"hi\"}"),
			atEOF:   false,
			length:  0,
			content: nil,
			err:     errors.New("strconv.Atoi: parsing \"--\": invalid syntax"),
		},
		{
			name:    "message with incorrect content length",
			data:    []byte("Content-Length: 100\r\n\r\n{\"Method\":\"hi\"}"),
			atEOF:   false,
			length:  0,
			content: nil,
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			length, content, err := rpc.Split(tt.data, tt.atEOF)
			if !equalErrors(t, err, tt.err) {
				t.Fatalf("Unexpected error:\n\tgot:  %v\n\twant: %v", err, tt.err)
			}

			if length != tt.length {
				t.Errorf("Invalid length:\n\tgot:  %d\n\twant: %d", length, tt.length)
			}

			if !bytes.Equal(content, tt.content) {
				t.Errorf("Invalid content:\n\tgot:  %s\n\twant: %s", content, tt.content)
			}
		})
	}
}

func equalErrors(t testing.TB, a, b error) bool {
	t.Helper()

	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return a.Error() == b.Error()
}
