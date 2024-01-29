// https://exercism.org/tracks/go/exercises/luhn
package easy

func Valid(id string) bool {
	var sum, valid, i, current int

	for i = len(id) - 1; i >= 0; i-- {
		c := id[i]

		switch {
		case c == ' ':
			continue
		case c >= '0' && c <= '9':
			current = int(c - '0')

			if valid%2 == 1 {
				current <<= 1
			}

			if current > 9 {
				current -= 9
			}

			sum += current
			valid++
		default:
			return false
		}
	}

	return valid > 1 && sum%10 == 0
}
