package main

import "fmt"

// Explanation
// An array is a fixed-size collection of elements of the same type. Once declared, the size cannot change.

func main() {
	// Example 1: Declaring and Accessing Arrays
	// var numbers [5]int

	// numbers[0] = 10
	// numbers[1] = 20
	// numbers[2] = 30
	// numbers[3] = 40
	// numbers[4] = 50

	// fmt.Println("First element:", numbers[0])
	// fmt.Println("Full array:", numbers)
	// fmt.Println("Array length:", len(numbers))

	// Example 2: Array Initialization
	// Method 1: Declare with values
	fruits := [3]string{"Apple", "Banana", "Orange"}

	// Method 2: Let Go count the size
	colors := [...]string{"Red", "Green", "Blue"}

	// Method 3: Specific indices
	scores := [5]int{0: 100, 2: 85, 4: 90}

	fmt.Println("Fruits:", fruits)
	fmt.Println("Colors:", colors)
	fmt.Println("Scores:", scores)

	// Example 3: Looping Through Arrays
	// Method 1: Traditional for loop
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Index:", i, "Value:", numbers[i])
	// }

	// Method 2: Range (cleaner)
	for index, value := range fruits {
		fmt.Println("Index:", index, "Value:", value)
	}

	for _, value := range colors {
		fmt.Println("Color Value:", value)
	}

	// Exercises 1: Create an array of 7 days of the week, then print only the weekend days.
	weekdays := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println("Method 1")
	for i := 5; i < len(weekdays); i++ {
		fmt.Println(weekdays[i])
	}

	fmt.Println("Method 2")
	for _, value := range weekdays[5:] {
		fmt.Println(value)
	}

	// Exercises 2: Create an array of 10 integers, fill it with user input, then find and print the largest number.
	var numbers [10]int

	for i := range len(numbers) {
		fmt.Println("Enter number:")
		fmt.Scan(&numbers[i])
	}

	var max int = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	fmt.Println("Largest number:", max)

	// Exercises 3: Create two arrays of 5 integers each, then create a third array that contains the sum of corresponding elements.
	var array1 [5]int
	var array2 [5]int
	var array3 [5]int

	for i := range len(array1) {
		fmt.Println("Enter number:")
		fmt.Scan(&array1[i])
	}

	for i := range len(array2) {
		fmt.Println("Enter number:")
		fmt.Scan(&array2[i])
	}

	for i := range len(array3) {
		array3[i] = array1[i] + array2[i]
	}

	fmt.Println("Array 1:", array1)
	fmt.Println("Array 2:", array2)
	fmt.Println("Array 3:", array3)
}
