// https://exercism.org/tracks/go/exercises/party-robot
package learningexercise

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	firstPhase := fmt.Sprintf("Happy birthday %s! ", name)
	secondPhase := fmt.Sprintf("You are now %d years old!", age)

	return firstPhase + secondPhase
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	firstPhase := fmt.Sprintf("\nWelcome to my party, %s!\n", name)
	secondPhase := fmt.Sprintf("You have been assigned to table %03d. ", table)
	thirdPhase := fmt.Sprintf("Your table is %s, exactly %.1f meters from here.\n", direction, distance)
	fourthPhase := fmt.Sprintf("You will be sitting next to %s.", neighbor)

	return firstPhase + secondPhase + thirdPhase + fourthPhase
}
