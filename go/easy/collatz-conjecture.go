// https://exercism.org/tracks/go/exercises/collatz-conjecture
package easy

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0

	if n <= 0 {
		return 0, errors.New("input must be a positive integer")
	}

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}

		count++
	}

	return count, nil
}
