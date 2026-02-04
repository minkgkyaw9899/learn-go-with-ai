package main

import "fmt"

// Explanation
// fmt.Scan() reads input from the user. fmt.Printf() allows formatted output using verbs:

// %d - decimal integer
// %f - float
// %s - string
// %t - boolean
// %v - default format
// %.2f - float with 2 decimal places


func main() {
	// Example 1: Basic Input
	// var name string
	// fmt.Println("Enter your name:")
	// fmt.Scan(&name)
	// fmt.Println("Hello,", name)

	// Example 2: Multiple Inputs
	// var (
	// 	age int
	// 	height float64
	// )

	// fmt.Println("Enter your age:")
	// fmt.Scanln(&age)

	// fmt.Println("Enter your height (in meters):")
	// fmt.Scanln(&height)

	// fmt.Printf("You are %d years old and %.2f meters tall\n", age, height)

	// Example 3: Formatted Output
	// name := "Bob"
    // score := 95.678
    // passed := true
    
    // fmt.Printf("Student: %s\n", name)
    // fmt.Printf("Score: %.2f%%\n", score)
    // fmt.Printf("Passed: %t\n", passed)

	// Exercises 1: Ask the user for their name and favorite number, then print a message like "Hi [name], your lucky number is [number]!"
	var (
		name string
		luckyNum int
	)

	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	fmt.Println("Enter your lucky number:")
	fmt.Scanln(&luckyNum)

	fmt.Printf("Hi %s, your lucky number is %d", name, luckyNum)

	// Exercises 2: Create a simple calculator that asks for two numbers and prints their sum, formatted to 2 decimal places.

	var (
		num1 float64
		num2 float64
	)

	fmt.Println("Enter number 1")
	fmt.Scanln(&num1)

	fmt.Println("Enter number 2")
	fmt.Scanln(&num2)

	fmt.Printf("total sum of %.2f + %.2f is %.2f", num1, num2, num1 + num2)

	// Exercises 3: Ask for product name, quantity, and price per unit. Calculate and display the total cost.
	var productName string
	var unit int
	var price float64

	fmt.Println("Enter product name: ")
	fmt.Scanln(&productName)

	fmt.Println("Enter unit:")
	fmt.Scanln(&unit)

	fmt.Printf("Enter %s price: ", productName)
	fmt.Println()
	fmt.Scanln(&price)

	fmt.Printf("Total cost of %s is %.2f", productName, float64(unit) * price)
}