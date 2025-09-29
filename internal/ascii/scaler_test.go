package ascii

import (
	"image"
	"image/color"
	"testing"
)

func TestImageToASCII(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
	}{
		{"2x2 image", 2, 2},
		{"4x4 image", 4, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test grayscale image
			img := image.NewRGBA(image.Rect(0, 0, tt.width, tt.height))
			for y := 0; y < tt.height; y++ {
				for x := 0; x < tt.width; x++ {
					gray := uint8((x + y) * 255 / (tt.width + tt.height))
					img.Set(x, y, color.Gray{Y: gray})
				}
			}

			// Call new function signature (with charset)
			result := imageToASCII(img, "@%#*+=-:. ")
			if len(result) == 0 {
				t.Error("expected ASCII output, got empty string")
			}
		})
	}
}
