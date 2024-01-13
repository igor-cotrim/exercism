package main

import (
	"fmt"

	learningexercise "github.com/igor-cotrim/exercism-go/learning-exercise"
)

func main() {
	//Lasagna
	fmt.Println("LASAGNA")
	fmt.Println("Remaining Oven Time:", learningexercise.RemainingOvenTime(20))
	fmt.Println("Preparation Time:", learningexercise.PreparationTime(3))
	fmt.Println("Elapsed Time:", learningexercise.ElapsedTime(3, 20))

	fmt.Println("--------------------")

	// annalyns-infiltration
	fmt.Println("ANNALYNS-INFILTRATION")
	fmt.Println("Can Fast Attack:", learningexercise.CanFastAttack(true))
	fmt.Println("Can Spy:", learningexercise.CanSpy(false, true, false))
	fmt.Println("Can Signal Prisoner:", learningexercise.CanSignalPrisoner(false, true))
	fmt.Println("Can Free Prisoner:", learningexercise.CanFreePrisoner(false, true, false, false))

	fmt.Println("--------------------")

	// welcome-to-tech-palace
	fmt.Println("WELCOME-TO-TECH-PALACE")
	fmt.Println("Calculate Cost:", learningexercise.CleanupMessage(`
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`))

	fmt.Println("--------------------")

	// card-tricks
	fmt.Println("CARD-TRICKS")
	fmt.Println("Favorite Cards:", learningexercise.FavoriteCards())

	fmt.Println("--------------------")

	// lasagna-master
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}

	fmt.Println("Preparation Time Lasagna Master:", learningexercise.PreparationTimeLasagnaMaster(layers, 3))
	fmt.Println("Preparation Time Lasagna Master:", learningexercise.PreparationTimeLasagnaMaster(layers, 0))

	fmt.Println("--------------------")

	// booking-up-for-beauty
	fmt.Println("Schedule:", learningexercise.Schedule("7/25/2019 13:45:00"))
	fmt.Println("Has Passed:", learningexercise.HasPassed("December 9, 2112 11:59:59"))
	fmt.Println("Is Afternoon Appointment:", learningexercise.IsAfternoonAppointment("Friday, March 8, 1974 12:02:02"))
	fmt.Println("Description:", learningexercise.Description("7/25/2019 13:45:00"))

	fmt.Println("--------------------")

	// animal-magic
	fmt.Println("Roll A Die:", learningexercise.RollADie())
	fmt.Println("Roll A Die:", learningexercise.GenerateWandEnergy())
}
