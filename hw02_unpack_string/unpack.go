package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var output strings.Builder

	if len(s) == 0 {
		return "", nil
	}
	digitsMap := map[rune]int{
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

	chars := []rune(s)

	_, ok := digitsMap[chars[0]]
	if ok {
		return "", ErrInvalidString
	}

	for i := 1; i < len(s); i++ {
		prevChar := chars[i-1]
		digit, ok := digitsMap[chars[i]]
		if ok {
			_, ok := digitsMap[prevChar]
			if ok {
				return "", ErrInvalidString
			}
			if digit == 0 {
				continue
			}
			output.WriteString(strings.Repeat(string(prevChar), digit))
		} else {
			_, ok := digitsMap[prevChar]
			if !ok {
				output.WriteRune(prevChar)
			}
			if i == len(s)-1 {
				output.WriteRune(chars[i])
			}
		}
	}

	return output.String(), nil
}
