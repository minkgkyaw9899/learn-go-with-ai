package main

import "fmt"

// Explanation
// A slice is a dynamic, flexible view into an array. Unlike arrays, slices can grow and shrink.

// Key Operations:

// append() - add elements
// len() - current length
// cap() - capacity
// Slicing: slice[start:end]

func main() {
	// Example 1: Creating and Using Slices
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	numbers = append(numbers, 60)

	fmt.Println("after append Slice: ", numbers)
	fmt.Println("after append Length:", len(numbers))
	fmt.Println("after append Capacity:", cap(numbers))
}
