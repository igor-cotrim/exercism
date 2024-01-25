// https://exercism.org/tracks/go/exercises/isogram
package easy

import (
	"strings"
	"unicode"
)

// The IsIsogram function checks if a given word is an isogram, meaning it does not
// contain any repeated letters.
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i, letter := range word {
		if unicode.IsLetter(letter) && strings.ContainsRune(word[i+1:], letter) {
			return false
		}
	}
	return true
}
