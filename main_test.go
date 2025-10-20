package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestGenerateAndPrintQR(t *testing.T) {
	tests := []struct {
		name    string
		content string
		size    int
		wantErr bool
	}{
		{
			name:    "valid URL",
			content: "https://example.com",
			size:    256,
			wantErr: false,
		},
		{
			name:    "short text",
			content: "test",
			size:    256,
			wantErr: false,
		},
		{
			name:    "empty string should error",
			content: "",
			size:    256,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout to prevent terminal output during tests
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := generateAndPrintQR(tt.content, tt.size)

			// Restore stdout
			w.Close()
			os.Stdout = old
			io.ReadAll(r) // Drain the pipe

			if (err != nil) != tt.wantErr {
				t.Errorf("generateAndPrintQR() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrintQRToTerminal(t *testing.T) {
	// Create a simple 3x3 bitmap for testing
	bitmap := [][]bool{
		{true, false, true},
		{false, true, false},
		{true, false, true},
	}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Should not panic
	printQRToTerminal(bitmap)

	// Restore stdout
	w.Close()
	os.Stdout = old
	output, _ := io.ReadAll(r)

	// Verify we got some output
	if len(output) == 0 {
		t.Error("printQRToTerminal() produced no output")
	}

	// Verify output contains block characters
	outputStr := string(output)
	if !bytes.Contains(output, []byte(fullBlock)) {
		t.Errorf("printQRToTerminal() output missing full blocks, got: %s", outputStr)
	}
}

func TestPrintBorder(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test with a width of 10
	printBorder(10)

	// Restore stdout
	w.Close()
	os.Stdout = old
	output, _ := io.ReadAll(r)

	// Should output 10 full blocks plus a newline
	expectedLength := len(fullBlock)*10 + 1 // +1 for newline
	if len(output) != expectedLength {
		t.Errorf("printBorder(10) output length = %d, want %d", len(output), expectedLength)
	}
}
