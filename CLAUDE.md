# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

qrterm is a lightweight Go CLI tool that generates QR codes from URLs and displays them in the terminal using Unicode block characters. The entire application is contained in a single `main.go` file with three core functions.

## Architecture

This is a single-file Go application with a simple architecture:

- **main()**: Handles CLI argument parsing using the `flag` package and delegates to QR generation
- **generateAndPrintQR()**: Uses the `github.com/skip2/go-qrcode` library to generate QR code bitmaps
- **printQRToTerminal()**: Renders the bitmap to terminal using Unicode characters (`██` for black, `  ` for white) with decorative borders and padding

The QR code is generated with `qrcode.Medium` error correction level and rendered as a boolean bitmap. The size flag is accepted but the actual display size is determined by the QR code library's bitmap dimensions.

## Building and Running

Build the binary:
```bash
go build
```

Run directly with go:
```bash
go run main.go <URL>
```

Run the built binary:
```bash
./qrterm <URL>
```

With custom size:
```bash
./qrterm -size 512 <URL>
```

## Dependencies

The project uses Go modules (`go.mod`) with a single external dependency:
- `github.com/skip2/go-qrcode` - QR code generation library

To update dependencies:
```bash
go mod tidy
```

## Installation

Users can install via:
```bash
go install github.com/infamousjoeg/qrterm@latest
```

## Key Implementation Details

- Uses Unicode block characters (`fullBlock = "██"`, `emptyBlock = "  "`) for rendering
- Adds 2-character padding on all sides of the QR code for visual clarity
- Prints decorative borders above and below the QR code
- Error handling follows standard Go patterns with descriptive error wrapping
