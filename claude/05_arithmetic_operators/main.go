package main

import "fmt"

// Explanation
// Go supports standard arithmetic operators:

// + Addition
// - Subtraction
// * Multiplication
// / Division
// % Modulus (remainder)

// Important: Division of integers results in an integer (truncated). For decimal results, use float64

func main() {
	// Example 1: Basic Operations
	a := 10
	b := 3

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Modulus:", a%b)

	// Example 2: Float Division
	// var x float64 = 10.0
	// var y float64 = 3.0

	// result := x / y
	// fmt.Printf("Float Division: %.4f\n", result)

	// Example 3: Compound Assignment
	x := 10

	x += 5
	fmt.Println("After += 5:", x)

	x -= 3
	fmt.Println("After -= 3:", x)

	x *= 2
	fmt.Println("After *= 2:", x)

	x /= 4
	fmt.Println("After /= 4:", x)

	// Exercises 1: Create a simple calculator that takes two numbers and prints all five arithmetic operations.
	num1 := 10
	num2 := 3

	fmt.Println("Addition:", num1+num2)
	fmt.Println("Subtraction:", num1-num2)
	fmt.Println("Multiplication:", num1*num2)
	fmt.Println("Division:", num1/num2)
	fmt.Println("Modulus:", num1%num2)

	// Exercises 2: Calculate the area and perimeter of a rectangle given length and width from user input.
	var length int
	var width int

	fmt.Println("Enter the length of the rectangle:")
	fmt.Scanln(&length)
	fmt.Println("Enter the width of the rectangle:")
	fmt.Scanln(&width)

	fmt.Println("Area:", length*width)
	fmt.Println("Perimeter:", 2*(length+width))

	// Exercises 3: Write a program that converts minutes to hours and remaining minutes (e.g., 125 minutes = 2 hours and 5 minutes). Use division and modulus.

	var minutes int
	fmt.Println("Enter the number of minutes:")
	fmt.Scanln(&minutes)

	hours := minutes / 60
	remainingMinutes := minutes % 60

	fmt.Printf("%d minutes = %d hours and %d minutes\n", minutes, hours, remainingMinutes)
}
