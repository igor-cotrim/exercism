// https://exercism.org/tracks/go/exercises/the-farm
package learningexercise

import (
	"errors"
	"fmt"
)

// This file contains types used in the exercise and tests and should not be modified.

// FodderCalculator provides helper methods to determine the optimal
// amount of fodder to feed cows.
type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (haveACow InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", haveACow.cows, haveACow.message)
}

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, cows int) (float64, error) {
	fodderAmount, err := f.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := f.FatteningFactor()
	if err != nil {
		return 0, err
	}

	amount := fodderAmount * fatteningFactor / float64(cows)
	if amount < 1 {
		return 0, err
	}

	return amount, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(f, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	}

	if cows == 0 {
		return &InvalidCowsError{cows, "no cows don't need food"}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
