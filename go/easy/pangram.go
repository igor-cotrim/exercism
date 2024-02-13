// https://exercism.org/tracks/go/exercises/pangram
package easy

import "unicode"

func IsPangram(input string) bool {
	alphabet := make(map[rune]bool)

	for _, char := range input {
		if unicode.IsLetter(char) {
			alphabet[unicode.ToLower(char)] = true
		}
	}

	return len(alphabet) == 26
}
