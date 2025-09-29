package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hamed0406/go-ascii-art/internal/ascii"
)

// main is the entrypoint for the CLI application.
// It parses flags, validates input, and calls the ASCII converter.
func main() {
	// Define CLI flags
	width := flag.Uint("width", 80, "output width in characters")
	color := flag.Bool("color", false, "enable ANSI color output (true/false)")

	flag.Parse()

	// Require at least one non-flag argument (image path)
	if flag.NArg() < 1 {
		fmt.Println("Usage: asciiart [options] <image-path>")
		flag.PrintDefaults()
		return
	}

	imagePath := flag.Arg(0)

	// Call the converter with parsed options
	result, err := ascii.ConvertImageToASCII(imagePath, *width, *color)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print final ASCII result
	fmt.Println(result)
}
