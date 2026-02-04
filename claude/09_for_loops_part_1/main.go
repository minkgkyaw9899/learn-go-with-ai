package main

import "fmt"

// Explanation
// Go has only ONE loop keyword: for. But it can work in different ways:

// Traditional for loop with initialization, condition, post
// While-style loop (just condition)
// Infinite loop
func main() {
	// Example 1: Traditional For Loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Example 2: While-Style Loop
	count := 1
	for count <= 10 {
		fmt.Println(count)
		count++
	}

	// Example 3: Infinite Loop with Break
	x := 1
	for {
		x++
		fmt.Println(x)
		if x == 10 {
			break
		}
	}

	// Exercises 1: Print all even numbers from 2 to 20.
	for i := 2; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Exercises 2: Calculate the sum of numbers from 1 to 100 using a for loop.
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// Exercises 3: Create a countdown timer from 10 to 1, then print "Blast off!"
	timer := 10

	for {
		fmt.Println(timer)
		timer--
		if timer == 0 {
			break
		}
	}

}
