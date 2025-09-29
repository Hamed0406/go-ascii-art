package utils

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// GenerateTestImage creates a simple 10x10 gradient PNG in testdata/test.png
func GenerateTestImage() error {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			gray := uint8((x + y) * 255 / 20)
			img.Set(x, y, color.Gray{Y: gray})
		}
	}
	file, err := os.Create("testdata/test.png")
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}
