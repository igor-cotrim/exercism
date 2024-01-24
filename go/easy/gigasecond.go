// https://exercism.org/tracks/go/exercises/gigasecond
package easy

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
