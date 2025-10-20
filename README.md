# qrterm

A simple CLI tool that generates QR codes from URLs and displays them directly in your terminal. Perfect for presentations where you want to share links with your audience!

## Features

- Generate QR codes from any URL
- Display QR codes directly in the terminal using Unicode block characters
- Clean, bordered output perfect for presentations
- Zero runtime dependencies (single binary)
- Fast and lightweight

## Installation

### From Source

```bash
go install github.com/infamousjoeg/qrterm@latest
```

### Build Locally

```bash
git clone https://github.com/infamousjoeg/qrterm.git
cd qrterm
go build
```

## Usage

Basic usage:

```bash
qrterm https://example.com
```

With custom size:

```bash
qrterm -size 512 https://example.com
```

### Options

- `-size int` - QR code size in pixels (default: 256). Note: The terminal display is adjusted automatically for readability.

## Example

```bash
$ qrterm https://github.com/infamousjoeg/qrterm

████████████████████████████████████████████████

        ██████████████  ████  ██████████████
        ██          ██  ████  ██          ██
        ██  ██████  ██  ████  ██  ██████  ██
        ██  ██████  ██  ████  ██  ██████  ██
        ██  ██████  ██  ████  ██  ██████  ██
        ██          ██  ████  ██          ██
        ██████████████  ████  ██████████████

████████████████████████████████████████████████
```

The generated QR code can be scanned with any QR code reader on smartphones or other devices.

## Use Cases

- **Presentations**: Display links to slides, repositories, or documentation
- **Workshops**: Share resource links with participants
- **Demos**: Quick sharing of app URLs or documentation
- **Remote Sessions**: Share links when screen sharing

## How It Works

qrterm uses the [go-qrcode](https://github.com/skip2/go-qrcode) library to generate QR codes and renders them using Unicode block characters (`██`) for optimal visibility in terminal environments.

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
