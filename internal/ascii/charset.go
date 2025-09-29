package ascii

// ASCII character set from darkest to brightest
var asciiChars = []rune("@%#*+=-:. ")

// mapBrightnessToChar maps brightness (0–255) to an ASCII character
func mapBrightnessToChar(brightness uint8) rune {
	index := int((float64(brightness) / 255.0) * float64(len(asciiChars)-1))
	return asciiChars[index]
}
