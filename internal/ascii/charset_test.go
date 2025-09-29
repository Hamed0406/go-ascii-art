package ascii

import "testing"

func TestMapBrightnessToChar(t *testing.T) {
	tests := []struct {
		name       string
		brightness uint8
	}{
		{"darkest", 0},
		{"quarter", 64},
		{"middle", 128},
		{"three quarters", 192},
		{"brightest", 255},
	}

	// Use default charset for testing
	chars := parseCharset("")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mapBrightnessToChar(tt.brightness, chars)
			if r == 0 {
				t.Errorf("unexpected empty rune for brightness %d", tt.brightness)
			}
		})
	}
}
