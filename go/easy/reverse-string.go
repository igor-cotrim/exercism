// https://exercism.org/tracks/go/exercises/reverse-string
package easy

func Reverse(input string) string {
	runeSlice := []rune(input)
	var output string

	for i := len(runeSlice) - 1; i >= 0; i-- {
		output += string(runeSlice[i])
	}

	return output
}
