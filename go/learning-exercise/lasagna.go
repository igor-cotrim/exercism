// https://exercism.org/tracks/go/exercises/lasagna

package learningexercise

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	OverTime := 40
	remainingTime := OverTime - actualMinutesInOven

	return remainingTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	preparationTime := numberOfLayers * 2

	return preparationTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	preparationTime := numberOfLayers * 2
	elapsedTime := preparationTime + actualMinutesInOven

	return elapsedTime
}
