// https://exercism.org/tracks/go/exercises/bob
package easy

import (
	"strings"
)

// Hey makes bob say hey on different ways based on a remark
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	upper := strings.ToUpper(remark)
	lower := strings.ToLower(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if strings.HasSuffix(remark, "?") {
		if upper == remark && upper != lower {
			return "Calm down, I know what I'm doing!"
		}

		return "Sure."
	}

	if upper == remark && upper != lower {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
