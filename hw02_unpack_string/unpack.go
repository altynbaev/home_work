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

func isDigit(r rune, digits map[rune]int) bool {
	_, ok := digits[r]
	if ok {
		return true
	}
	return false
}

func Unpack(s string) (string, error) {
	var output strings.Builder

	if len(s) == 0 {
		return "", nil
	}

	chars := []rune(s)

	if isDigit(chars[0], digitsMap) {
		return "", ErrInvalidString
	}

	for i := 1; i < len(s); i++ {
		prevChar := chars[i-1]
		digit, ok := digitsMap[chars[i]]
		if ok {
			if isDigit(prevChar, digitsMap) {
				return "", ErrInvalidString
			}
			if digit == 0 {
				continue
			}
			output.WriteString(strings.Repeat(string(prevChar), digit))
		} else {
			if !isDigit(prevChar, digitsMap) {
				output.WriteRune(prevChar)
			}
			if i == len(s)-1 {
				output.WriteRune(chars[i])
			}
		}
	}

	return output.String(), nil
}
