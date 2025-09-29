# go-ascii-art

Transform images into expressive ASCII art using Go. This repository provides a
small library for image-to-ASCII conversion as well as two command line tools:

- `asciiart`: convert any PNG or JPEG image into ASCII characters, optionally
  with ANSI colors.
- `genimage`: generate a small gradient sample image stored at
  `testdata/test.png` for quick experimentation.

## Features

- Automatic image resizing to a specified character width while preserving the
  aspect ratio.
- Grayscale ASCII conversion that maps brightness levels to a curated character
  set ranging from dark (`@`) to light (` `).
- Optional 24-bit color output that prints colored block characters for a vivid
  terminal preview.
- Utility helpers for detecting file existence and generating sample imagery.
- Thorough unit tests for the charset mapping, scaling logic, converters, and
  file utilities.

## Getting started

### Prerequisites

- Go 1.21 or later (see `go.mod` for the exact version used by the project).

### Installation

Clone the repository and download dependencies:

```bash
git clone https://github.com/hamed0406/go-ascii-art.git
cd go-ascii-art
go mod download
```

To install the CLI tools into your `GOBIN` (or `GOPATH/bin`):

```bash
go install ./cmd/asciiart
# Optionally install the helper image generator
# go install ./cmd/genimage
```

## Usage

### Convert an image to ASCII

Provide the path to any PNG or JPEG file. The tool automatically resizes the
image to the requested width (default `80` characters) while maintaining the
aspect ratio.

```bash
asciiart -width 100 path/to/image.png
```

To create a colorful representation instead of grayscale characters, enable the
`-color` flag. The output uses ANSI escape codes, so ensure your terminal
supports true color.

```bash
asciiart -color path/to/image.jpg
```

If the provided file path does not exist, the CLI falls back to searching within
`testdata/` for convenience when experimenting with bundled samples.

### Generate a sample image

Use the helper command to create a simple 10x10 gradient image at
`testdata/test.png`:

```bash
go run ./cmd/genimage
# or, after installation
genimage
```

This sample image is handy for validating conversions without sourcing your own
image.

## Development

### Project layout

```
cmd/            # Command line entry points
  asciiart/     # Main CLI for converting images
  genimage/     # Helper tool for generating a sample test image
internal/
  ascii/        # Conversion library: resizing, grayscale, color rendering
  utils/        # Shared helpers (file detection, test image creation)
testdata/       # Sample images used in tests and demos
```

### Run tests

Execute the full suite to verify functionality:

```bash
go test ./...
```

### Contributing

Contributions are welcome! Please open an issue or pull request with
improvements, bug fixes, or new features.

## License

This project is licensed under the MIT License. See [`LICENSE`](LICENSE) for
details.
