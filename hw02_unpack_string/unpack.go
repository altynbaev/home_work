package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

var digitsMap = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func isDigit(r rune) bool {
	_, ok := digitsMap[r]
	return ok
}

func Unpack(s string) (string, error) {
	var output strings.Builder

	if len(s) == 0 {
		return "", nil
	}

	chars := []rune(s)

	if isDigit(chars[0]) {
		return "", ErrInvalidString
	}

	for i := 1; i < len(s); i++ {
		prevChar := chars[i-1]
		digit, ok := digitsMap[chars[i]]
		if ok {
			if isDigit(prevChar) {
				return "", ErrInvalidString
			}
			if digit == 0 {
				continue
			}
			output.WriteString(strings.Repeat(string(prevChar), digit))
		} else if !isDigit(prevChar) {
			output.WriteRune(prevChar)
		}
	}
	output.WriteRune(chars[len(s)-1])

	return output.String(), nil
}
