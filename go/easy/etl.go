// https://exercism.org/tracks/go/exercises/etl
package easy

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)

	for key, value := range in {
		for _, v := range value {
			out[strings.ToLower(v)] = key
		}
	}

	return out
}
