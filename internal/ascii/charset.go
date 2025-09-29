package ascii

// Default character ramp (dark → light).
var defaultCharset = []rune("@%#*+=-:. ")

// mapBrightnessToChar maps brightness (0–255) to a character from the charset.
func mapBrightnessToChar(brightness uint8, charset []rune) rune {
	index := int((float64(brightness) / 255.0) * float64(len(charset)-1))
	return charset[index]
}

// parseCharset converts a string into a rune slice.
// If the input is empty, falls back to default.
func parseCharset(charset string) []rune {
	if len(charset) == 0 {
		return defaultCharset
	}
	return []rune(charset)
}
