// https://exercism.org/tracks/go/exercises/meteorology
package learningexercise

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (temp TemperatureUnit) String() string {
	units := []string{"°C", "°F"}

	return units[temp]
}

// Add a String method to the Temperature type
func (d Temperature) String() string {
	return fmt.Sprintf("%v %v", d.degree, d.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (temp SpeedUnit) String() string {
	units := []string{"km/h", "mph"}

	return units[temp]
}

// Add a String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (md MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", md.location, md.temperature, md.windDirection, md.windSpeed, md.humidity)
}
