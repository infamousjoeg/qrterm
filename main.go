// Package main provides a CLI tool for generating and displaying QR codes in the terminal.
// It takes a URL as input and renders a scannable QR code using Unicode block characters.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

const (
	// Unicode block characters for better visual quality
	fullBlock  = "██"
	emptyBlock = "  "
)

func main() {
	// Define flags
	size := flag.Int("size", 256, "QR code size (will be adjusted for terminal display)")
	flag.Parse()

	// Get URL from arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: qrterm [OPTIONS] <URL>")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExample:")
		fmt.Fprintln(os.Stderr, "  qrterm https://example.com")
		os.Exit(1)
	}

	url := args[0]

	// Generate QR code
	if err := generateAndPrintQR(url, *size); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating QR code: %v\n", err)
		os.Exit(1)
	}
}

func generateAndPrintQR(content string, size int) error {
	// Generate QR code with medium recovery level
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return fmt.Errorf("failed to create QR code: %w", err)
	}

	// Get the bitmap representation
	bitmap := qr.Bitmap()

	// Print the QR code to terminal
	printQRToTerminal(bitmap)

	return nil
}

func printQRToTerminal(bitmap [][]bool) {
	// Add top padding
	fmt.Println()

	// Add border
	borderWidth := len(bitmap[0]) + 4
	printBorder(borderWidth)

	// Print each row with side padding
	for _, row := range bitmap {
		fmt.Print(emptyBlock) // Left padding
		for _, pixel := range row {
			if pixel {
				fmt.Print(fullBlock) // Black pixel
			} else {
				fmt.Print(emptyBlock) // White pixel
			}
		}
		fmt.Print(emptyBlock) // Right padding
		fmt.Println()
	}

	// Add bottom border
	printBorder(borderWidth)

	// Add bottom padding
	fmt.Println()
}

func printBorder(width int) {
	for i := 0; i < width; i++ {
		fmt.Print(fullBlock)
	}
	fmt.Println()
}
