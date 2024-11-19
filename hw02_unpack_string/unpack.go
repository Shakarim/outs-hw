package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(receivedString string) (string, error) {
	if receivedString == "" {
		return "", nil
	}

	var result []string = make([]string, len(receivedString))
	var currentLetter string = ""

	for _, char := range receivedString {
		if unicode.IsDigit(char) {
			if currentLetter == "" {
				return "", ErrInvalidString
			}
			intValue, err := strconv.Atoi(string(char))
			if err != nil {
				return "", ErrInvalidString
			}
			result = append(result, strings.Repeat(string(currentLetter), intValue))
			currentLetter = ""
		} else {
			if currentLetter != "" {
				result = append(result, string(currentLetter))
			}
			currentLetter = string(char)
		}
	}

	if currentLetter != "" {
		result = append(result, string(currentLetter))
	}

	return strings.Join(result, ""), nil
}
