//https://exercism.org/tracks/go/exercises/weather-forecast

// Package weather provides tools to forecast weather showing the location and its weather condition.
package learningexercise

// CurrentCondition represents a current condition in a text.
var CurrentCondition string

// CurrentLocation represents a current location in a text.
var CurrentLocation string

// Forecast returns an text with a current location and the weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
