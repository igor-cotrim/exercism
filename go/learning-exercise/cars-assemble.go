// https://exercism.org/tracks/go/exercises/cars-assemble
package learningexercise

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	percentage := successRate / 100

	return float64(productionRate) * percentage
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	percentage := successRate / 100
	productionRatePerHour := float64(productionRate) * percentage
	productionRatePerMinute := int(productionRatePerHour) / 60

	return productionRatePerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupCost := 95000
	individualCost := 10000

	numGroups := carsCount / 10
	remainingCars := carsCount % 10

	totalCost := (numGroups * groupCost) + (remainingCars * individualCost)

	return uint(totalCost)
}
