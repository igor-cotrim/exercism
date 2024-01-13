// https://exercism.org/tracks/go/exercises/booking-up-for-beauty
package learningexercise

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"

	result, _ := time.Parse(layout, date)

	return result
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

	result, _ := time.Parse(layout, date)

	return result.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"

	result, _ := time.Parse(layout, date)

	return result.Hour() >= 12 && result.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"

	result, _ := time.Parse(layout, date)

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", result.Weekday(), result.Month(), result.Day(), result.Year(), result.Hour(), result.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	anniversary := time.Date(currentYear, 9, 15, 0, 0, 0, 0, time.UTC)

	return anniversary
}
