package main

import "fmt"

func main() {
	// Example 1: Continue Statement
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Example 2: Nested Loops (Multiplication Table)
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// Example 3: Pattern Printing
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Exercises 1: Print all numbers from 1 to 50, but skip multiples of 5 (use continue).
	for i := 1; i <= 50; i++ {
		if i%5 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Exercises 2: Create a program that prints a pyramid pattern with numbers:
	// ```
	// 1
	// 1 2
	// 1 2 3
	// 1 2 3 4
	// 1 2 3 4 5
	// ```
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}

	// Exercises 3:Find all prime numbers between 1 and 50 (use nested loops to check divisibility).
	for i := 1; i <= 50; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				break
			}
			fmt.Println(i)
		}
	}
}
