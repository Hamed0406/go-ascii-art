package ascii

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

// ConvertImageToASCII loads an image, rescales it, and converts it to ASCII art.
// - path: image file path
// - width: output width in characters
// - useColor: ANSI color output toggle
// - charset: brightness mapping ramp
// - scaleY: vertical scaling factor (default ~0.5 for terminals)
func ConvertImageToASCII(path string, width uint, useColor bool, charset string, scaleY float64) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}

	// Calculate height with vertical scaling
	origBounds := img.Bounds()
	aspectHeight := uint(float64(origBounds.Dy()) * (float64(width) / float64(origBounds.Dx())) * scaleY)

	// Resize using Lanczos3 interpolation
	resized := resize.Resize(width, aspectHeight, img, resize.Lanczos3)

	if useColor {
		return imageToASCIIColor(resized), nil
	}
	return imageToASCII(resized, charset), nil
}
