// https://exercism.org/tracks/go/exercises/space-age
package easy

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var year float64

	switch planet {
	case "Mercury":
		year = 0.2408467
	case "Venus":
		year = 0.61519726
	case "Earth":
		year = 1.0
	case "Mars":
		year = 1.8808158
	case "Jupiter":
		year = 11.862615
	case "Saturn":
		year = 29.447498
	case "Uranus":
		year = 84.016846
	case "Neptune":
		year = 164.79132
	default:
		return -1
	}

	return seconds / (year * float64(31557600))
}
