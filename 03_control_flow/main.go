package main

import "fmt"

func main() {
	// Example 1: Simple Check (Age)
	// age := 20
	// if age >= 18 {
	// 	println("You are an adult")
	// }

	// Example 2: If / Else (Password)
	password := "123456"
	if password == "123456" {
		println("Access granted")
	} else {
		println("Access denied. Wrong password.")
	}

	// Example 3: Multiple Choices (Else If)
	color := "yellow"
	if color == "red" {
		fmt.Println("Go!")
	} else if color == "yellow" {
		fmt.Println("Slow down!")
	} else if color == "green" {
		fmt.Println("Stop!")
	} else {
		fmt.Println("Invalid color")
	}

	// 	Exercise 1: The Bouncer Create a variable age.
	// If age is 18 or older, print "Welcome to the club."
	// If age is under 18, print "Sorry, you are too young."
	age := 18
	if age >= 18 {
		fmt.Println("Welcome to the club.")
	} else {
		fmt.Println("Sorry, you are too young.")
	}

	// 	Exercise 2: The Exam Grader Create a variable score.
	// If the score is 50 or higher, print "Passed".
	// If the score is lower than 50, print "Failed".
	score := 50
	if score >= 50 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}

	// 	Exercise 3: Positive or Negative? Create a variable number (you can make it positive like 10 or negative like -5).
	// Write logic to check if the number is Positive (greater than 0), Negative (less than 0), or Zero. (Hint: You will need if, else if, and else).
	number := 10
	if number > 0 {
		fmt.Println("Positive")
	} else if number < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}
