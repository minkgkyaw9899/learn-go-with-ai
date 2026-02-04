package main

import "fmt"

// Create a simple guessing game:
// - Set a secret number (e.g., 7)
// - Use a for loop to give the user 3 attempts
// - Ask user to guess the number
// - Tell them if their guess is too high, too low, or correct
// - If correct, break the loop
// - If they use all attempts, tell them they lost

func main() {
	secretNumber := 7

	for i := 1; i <= 3; i++ {
		var guess int

		fmt.Println("Guess number 1-30: ")
		fmt.Scan(&guess)

		if guess == secretNumber {
			fmt.Println("You guessed it!")
			break
		} else if guess < secretNumber {
			fmt.Println("Too low")
		} else {
			fmt.Println("Too high")
		}

		if i == 3 {
			fmt.Println("You lost!")
		}
	}
}
