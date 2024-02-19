// https://exercism.org/tracks/go/exercises/grains
package easy

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("square out of range")
	}

	return 1 << (number - 1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
