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
	// Direct check
	if utils.FileExists(imageName) {
		return imageName
	}

	// Fallback to testdata/
	testdataPath := filepath.Join("testdata", imageName)
	if utils.FileExists(testdataPath) {
		return testdataPath
	}

	// Return original (converter will error)
	return imageName
}

func main() {
	// CLI flags
	width := flag.Uint("width", 80, "output width in characters")
	color := flag.Bool("color", false, "enable ANSI color output (true/false)")
	flag.Parse()

	// Require image arg
	if flag.NArg() < 1 {
		fmt.Println("Usage: asciiart [options] <image-path>")
		flag.PrintDefaults()
		return
	}

	// Resolve image path with fallback
	imagePath := resolveImagePath(flag.Arg(0))

	// Run converter
	result, err := ascii.ConvertImageToASCII(imagePath, *width, *color)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
