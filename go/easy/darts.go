// https://exercism.org/tracks/go/exercises/darts
package easy

import "math"

func ScoreDarts(x, y float64) int {
	r := math.Sqrt(x*x + y*y)

	switch {
	case r <= 1:
		return 10
	case r <= 5:
		return 5
	case r <= 10:
		return 1
	default:
		return 0
	}
}
