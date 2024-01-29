// https://exercism.org/tracks/go/exercises/resistor-color
package easy

// Colors returns the list of all colors.
func Colors() []string {
	return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	colors := Colors()

	for i, c := range colors {
		if color == c {
			return i
		}
	}

	return -1
}
