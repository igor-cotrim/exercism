// https://exercism.org/tracks/go/exercises/bird-watcher
package learningexercise

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0

	for _, birds := range birdsPerDay {
		total += birds
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weeksDays := 7
	start := (week - 1) * weeksDays
	end := start + weeksDays
	total := 0

	for _, birds := range birdsPerDay[start:end] {
		total += birds
	}

	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}

	return birdsPerDay
}
