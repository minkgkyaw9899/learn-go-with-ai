package main

import "fmt"

// Challenge: Create a program that:

// Stores temperatures for 7 days in an array
// Takes user input for each day's temperature
// Calculates and displays:

// Average temperature
// Highest temperature
// Lowest temperature
// Number of days above average

// Hint: Use loops, arrays, and conditional statements!

func main() {
	var temperatures [7]int

	for i := range len(temperatures) {
		fmt.Println("Enter temperature for day", i+1)
		fmt.Scan(&temperatures[i])
	}

	sum := 0

	for i := range len(temperatures) {
		sum += temperatures[i]
	}

	average := float64(sum) / float64(len(temperatures))
	fmt.Println("Average temperature:", average)

	var highest int = temperatures[0]
	var lowest int = temperatures[0]
	var aboveAverage int = 0

	for i := range len(temperatures) {
		if temperatures[i] > highest {
			highest = temperatures[i]
		}
		if temperatures[i] < lowest {
			lowest = temperatures[i]
		}
		if temperatures[i] > int(average) {
			aboveAverage++
		}
	}

	fmt.Println("Highest temperature:", highest)
	fmt.Println("Lowest temperature:", lowest)
	fmt.Println("Number of days above average:", aboveAverage)
}
