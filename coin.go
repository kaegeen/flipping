package main

import (
	"fmt"
	"math/rand"
	"time"
)

func flipCoin() string {
	if rand.Intn(2) == 0 {
		return "Heads"
	}
	return "Tails"
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Virtual Coin Flip")
	fmt.Println("==================")
	fmt.Println("Type 'flip' to flip the coin or 'exit' to quit.\n")

	for {
		fmt.Print("Command: ")
		var command string
		fmt.Scanln(&command)

		if command == "exit" {
			fmt.Println("Goodbye! Thanks for flipping!")
			break
		} else if command == "flip" {
			result := flipCoin()
			fmt.Printf("The coin landed on: %s\n\n", result)
		} else {
			fmt.Println("Invalid command. Please type 'flip' or 'exit'.\n")
		}
	}
}
