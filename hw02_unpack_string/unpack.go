package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var character rune
	charIsSet := false
	var output strings.Builder

	for _, buffer := range input {
		if unicode.IsDigit(buffer) {
			if !charIsSet {
				return "", ErrInvalidString
			}
			count := int(buffer - '0')
			output.WriteString(strings.Repeat(string(character), count))
			charIsSet = false
		} else {
			if charIsSet {
				output.WriteRune(character)
			}
			character = buffer
			charIsSet = true
		}
	}

	if charIsSet {
		output.WriteRune(character)
	}

	return output.String(), nil
}
