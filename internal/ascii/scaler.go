package ascii

import (
	"fmt"
	"image"
	"strings"
)

// imageToASCII converts an image to grayscale ASCII art (no color).
// Uses a custom charset for brightness mapping.
func imageToASCII(img image.Image, charset string) string {
	var builder strings.Builder
	bounds := img.Bounds()
	chars := parseCharset(charset)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			builder.WriteRune(mapBrightnessToChar(gray, chars))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

// imageToASCIIColor converts an image to ANSI-colored ASCII art.
// Each pixel is represented as a colored block █ using 24-bit escape codes.
func imageToASCIIColor(img image.Image) string {
	var builder strings.Builder
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			R, G, B := r>>8, g>>8, b>>8

			// ANSI escape code: \033[38;2;R;G;Bm sets foreground color
			builder.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm█", R, G, B))
		}
		builder.WriteString("\033[0m\n") // reset colors
	}
	return builder.String()
}
