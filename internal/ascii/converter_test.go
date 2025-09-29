package ascii

import (
	"os"
	"testing"
)

func TestConvertImageToASCII(t *testing.T) {
	tests := []struct {
		name     string
		imgPath  string
		width    uint
		useColor bool
		charset  string
		scaleY   float64
		wantErr  bool
	}{
		{"grayscale default scale", "../../testdata/test.png", 40, false, "@%#*+=-:. ", 0.5, false},
		{"grayscale custom scale", "../../testdata/test.png", 40, false, "@%#*+=-:. ", 1.0, false},
		{"color mode", "../../testdata/test.png", 40, true, "@%#*+=-:. ", 0.5, false},
		{"nonexistent file", "../../testdata/missing.png", 40, false, "@%#*+=-:. ", 0.5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertImageToASCII(tt.imgPath, tt.width, tt.useColor, tt.charset, tt.scaleY)

			if (err != nil) != tt.wantErr {
				t.Errorf("expected error=%v, got %v", tt.wantErr, err)
			}
			if !tt.wantErr && len(result) == 0 {
				t.Error("expected non-empty ASCII result, got empty string")
			}
		})
	}
}

func TestTestImageExists(t *testing.T) {
	if _, err := os.Stat("../../testdata/test.png"); os.IsNotExist(err) {
		t.Skip("test image not found, skipping image conversion tests")
	}
}
