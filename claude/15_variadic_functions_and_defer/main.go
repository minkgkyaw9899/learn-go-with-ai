package main

import (
	"fmt"
	"strings"
)

// Example 1: Variadic Function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Example 2: Print Multiple Items
func printAll(items ...string) {
	fmt.Println("Items received:", len(items))
	for i, item := range items {
		fmt.Printf("%d: %s\n", i+1, item)
	}
}

func main() {
	fmt.Println(sum(1, 2, 3))        // 6
	fmt.Println(sum(10, 20, 30, 40)) // 100
	fmt.Println(sum(5))

	printAll("Apple", "Banana", "Orange")
	printAll("Go", "is", "awesome", "!")

	// Example 3: Defer Statement
	defer fmt.Println("World") // Executes last
	defer fmt.Println("Beautiful")

	fmt.Println("Hello") // Executes first

	fmt.Println("Average:", average(1, 2, 3, 4, 5))
	printPyramid(5)
	fmt.Println("Joined:", joinStrings(", ", "Apple", "Banana", "Orange"))

}

// Exercises 1: Create a variadic function average(numbers ...float64) float64 that calculates the average.
func average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

// Exercises 2: Write a function printPyramid(height int) that prints a number pyramid using defer for each line.
func printPyramid(height int) {
	for i := 1; i <= height; i++ {
		defer fmt.Println(i)
	}
}

// Exercises 3: Create a function joinStrings(separator string, words ...string) string that joins words with a separator.
func joinStrings(separator string, words ...string) string {
	return strings.Join(words, separator)
}
