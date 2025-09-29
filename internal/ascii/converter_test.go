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
		wantErr  bool
	}{
		{"valid grayscale small", "../../testdata/test.png", 20, false, false},
		{"valid grayscale large", "../../testdata/test.png", 80, false, false},
		{"valid color small", "../../testdata/test.png", 20, true, false},
		{"valid color large", "../../testdata/test.png", 80, true, false},
		{"nonexistent file", "../../testdata/missing.png", 40, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertImageToASCII(tt.imgPath, tt.width, tt.useColor)

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
