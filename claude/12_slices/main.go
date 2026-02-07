package main

import (
	"fmt"
)

// Explanation
// A slice is a dynamic, flexible view into an array. Unlike arrays, slices can grow and shrink.

// Key Operations:

// append() - add elements
// len() - current length
// cap() - capacity
// Slicing: slice[start:end]

func main() {
	// Example 1: Creating and Using Slices
	// numbers := []int{10, 20, 30, 40, 50}

	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))

	// numbers = append(numbers, 60)

	// fmt.Println("after append Slice: ", numbers)
	// fmt.Println("after append Length:", len(numbers))
	// fmt.Println("after append Capacity:", cap(numbers))

	// Example 2: Slice Operations
	// numbers := []int{10, 20, 30, 40, 50, 60}

	// fmt.Println("Full slice: ", numbers)
	// fmt.Println("First 3: ", numbers[0:3])
	// fmt.Println("Start from index 2: ", numbers[2:])
	// fmt.Println("Up to index 4: ", numbers[:4])
	// fmt.Println("Middle: ", numbers[2:5])

	// Example 3: Make Function

	slice1 := make([]int, 5)     // length=5, capacity=5
	slice2 := make([]int, 3, 10) // length=3, capacity=10

	fmt.Println("Slice1:", slice1, "Len:", len(slice1), "Cap:", cap(slice1))
	fmt.Println("Slice2:", slice2, "Len:", len(slice2), "Cap:", cap(slice2))

	for i := 0; i < len(slice2); i++ {
		slice2[i] = (i + 1) * 10
	}

	fmt.Println("Filled Slice2:", slice2)

	slice2 = append(slice2, 40)

	fmt.Println("New Slice2:", slice2)

	// Exercises 1: Create a slice of your favorite movies, add 2 more using append, then print all movies.
	movies := []string{"Godfather", "Iron Man", "The Dark Knight"}

	fmt.Println("Movies:", movies)

	movies = append(movies, "Avengers", "The Matrix")

	fmt.Println("Movies:", movies)

	// Exercises 2: Create a slice of numbers from 1-10, then create a new slice containing only numbers greater than 5.
	numbersSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Numbers: ", numbersSlice)

	graterThanFive := make([]int, 0)

	for _, num := range numbersSlice {
		if num > 5 {
			graterThanFive = append(graterThanFive, num)
		}
	}

	fmt.Println("GraterThanFive: ", graterThanFive)

	// Exercises 3: Write a program that removes an element at a specific index from a slice (hint: use slicing and append).
	nums := []int{10, 20, 30, 40, 50, 60}
	removeIndex := 2

	nums = append(nums[:removeIndex], nums[removeIndex+1:]...)

	fmt.Println("Nums: ", nums)

}
