package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hamed0406/go-ascii-art/internal/ascii"
	"github.com/hamed0406/go-ascii-art/internal/utils"
)

// resolveImagePath tries to find the image either in the provided path
// or, if not found, inside the "testdata" folder.
func resolveImagePath(imageName string) string {
	if utils.FileExists(imageName) {
		return imageName
	}
	testdataPath := filepath.Join("testdata", imageName)
	if utils.FileExists(testdataPath) {
		return testdataPath
	}
	return imageName // fallback, will error in converter
}

func main() {
	// CLI flags
	width := flag.Uint("width", 80, "output width in characters")
	color := flag.Bool("color", false, "enable ANSI color output (true/false)")
	outFile := flag.String("out", "", "output file path (if empty, print to stdout)")
	charset := flag.String("charset", "@%#*+=-:. ", "custom ASCII character ramp (dark→light)")

	flag.Parse()

	// Require image argument
	if flag.NArg() < 1 {
		fmt.Println("Usage: asciiart [options] <image-path>")
		flag.PrintDefaults()
		return
	}

	// Resolve image path
	imagePath := resolveImagePath(flag.Arg(0))

	// Run converter
	result, err := ascii.ConvertImageToASCII(imagePath, *width, *color, *charset)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Output handling
	if *outFile != "" {
		if err := os.WriteFile(*outFile, []byte(result), 0644); err != nil {
			fmt.Println("Error writing file:", err)
			os.Exit(1)
		}
		fmt.Println("✅ ASCII art saved to", *outFile)
	} else {
		fmt.Println(result)
	}
}
