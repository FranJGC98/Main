package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

func play() {
	randomnumber := rand.Intn(100)
	var (
		numInsert int
		tries     int
	)
	const maxTries = 10

	for tries < maxTries {
		tries++
		fmt.Printf("Add a number to guess it (tries : %v ): ", maxTries-tries+1)
		fmt.Scan(&numInsert)
		if randomnumber == numInsert {
			fmt.Println("You Rock!")
			playAgain()
			return
		} else if numInsert < randomnumber {
			fmt.Println("The number is highier!")
		} else if numInsert > randomnumber {
			fmt.Println("The number is lower")
		}
	}
	fmt.Println("Game lost, number was ", randomnumber)
	playAgain()
}

func playAgain() {
	var choice string
	fmt.Println("Do you want play again? (y/n):")
	fmt.Scan(&choice)
	switch choice {
	case "y":
		play()
	case "n":
		fmt.Println("Thanks for playing!")
	default:
		fmt.Println("Wrong choice, please try again.")
		playAgain()
	}
}
