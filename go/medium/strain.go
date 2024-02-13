// https://exercism.org/tracks/go/exercises/strain
package medium

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](items []T, pred func(T) bool) []T {
	var kept []T

	for _, item := range items {
		if pred(item) {
			kept = append(kept, item)
		}
	}

	return kept
}

func Discard[T any](items []T, pred func(T) bool) []T {
	var dis []T

	for _, item := range items {
		if !pred(item) {
			dis = append(dis, item)
		}
	}

	return dis
}
