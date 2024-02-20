// https://exercism.org/tracks/go/exercises/roman-numerals
package easy

import (
	"bytes"
	"errors"
)

type arabicToRoman struct {
	arabic int
	roman  string
}

func ToRomanNumeral(input int) (string, error) {
	buffer := bytes.NewBufferString("")

	if input <= 0 || input >= 4000 {
		return "", errors.New("invalid input")
	}

	dictionary := []arabicToRoman{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, item := range dictionary {
		for input >= item.arabic {
			buffer.WriteString(item.roman)
			input -= item.arabic
		}
	}

	return buffer.String(), nil
}
