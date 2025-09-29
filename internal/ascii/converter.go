package ascii

import (
	"fmt"
	"image"
	_ "image/jpeg" // enable JPEG decoding
	_ "image/png"  // enable PNG decoding
	"os"

	"github.com/nfnt/resize"
)

// ConvertImageToASCII loads an image, resizes it, and converts it to ASCII art.
// - path: path to the image file
// - width: target width in characters
// - useColor: if true, ANSI color codes are used for output
func ConvertImageToASCII(path string, width uint, useColor bool) (string, error) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	// Decode image (JPEG/PNG supported by default)
	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}

	// Resize to requested width (auto height)
	resized := resize.Resize(width, 0, img, resize.Lanczos3)

	// Convert depending on mode
	if useColor {
		return imageToASCIIColor(resized), nil
	}
	return imageToASCII(resized), nil
}
