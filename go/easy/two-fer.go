// https://exercism.org/tracks/go/exercises/two-fer
package easy

import "fmt"

func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
