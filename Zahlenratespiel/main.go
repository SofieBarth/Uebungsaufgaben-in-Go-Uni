package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var randomnumber int
	var numberguess int
	var feedback string
	var numberoftryes int
	var lastnumberguess int

	fmt.Println("Try to guess a randomly chosen natural number between 0 and 100.")

	randomnumber = rand.Intn(100)
	//fmt.Println("Random number: ", randomnumber)

	numberoftryes = 1
	fmt.Printf("Enter your %v. guess:\n", numberoftryes)
	fmt.Scan(&numberguess)
	lastnumberguess = numberguess

	if randomnumber == numberguess {
		fmt.Println("CONGRATULATION! You found the right number.")
		fmt.Println("Your first guess was right!")
	} else if (numberguess < (randomnumber - 5)) || (numberguess > (randomnumber + 5)) {
		fmt.Println("cold")
	} else {
		fmt.Println("hot")
	}

	for i := false; i != true; {
		numberoftryes = numberoftryes + 1
		lastnumberguess = numberguess

		fmt.Printf("Enter your %v. guess:\n", numberoftryes)
		fmt.Scan(&numberguess)

		switch {
		case randomnumber == numberguess:
			feedback = "CONGRATULATION! You found the right number."
			i = true
		case distance(numberguess, randomnumber) < 6:
			feedback = "hot"
		case distance(numberguess, randomnumber) == distance(lastnumberguess, randomnumber):
			feedback = "same distance"
		case distance(numberguess, randomnumber) < distance(lastnumberguess, randomnumber):
			feedback = "warmer"
		case distance(numberguess, randomnumber) > distance(lastnumberguess, randomnumber):
			feedback = "colder"
		}
		fmt.Println(feedback)

		if feedback == "CONGRATULATION! You found the right number." {
			fmt.Println("You needed", numberoftryes, "tryes.")
			i = true
		}
	}
}

func distance(a, b int) int {
	var x int
	x = a - b
	if x < 0 {
		return -x
	}
	return x
}
