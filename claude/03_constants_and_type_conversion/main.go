package main

import "fmt"

func main() {
	// int(value)      // convert to int
	// float64(value)  // convert to float64
	// string(value)   // convert to string (with limitations)

	const PI = 3.142
	const CompanyName = "TechCorp"
	const MaxUsers = 100

	// Example 1: Constants
	fmt.Println("PI:", PI)
	fmt.Println("CompanyName:", CompanyName)
	fmt.Println("MaxUsers:", MaxUsers)

	// Example 2: Type Conversion
	var x int = 32
	var y float64 = float64(x)
	var z int = int(y)

	fmt.Printf("x: %d (type: int)\n", x)
	fmt.Printf("y: %f (type: float64)\n", y)
	fmt.Printf("z: %d (type: int)\n", z)

	// Example 3: Mixed Type Operations
	a := 10
	b := 2.5

	result := float64(a) + b

	fmt.Printf("result: %f (type: float64)\n", result)

	// Exercises 1: Create constants for days in a week, hours in a day, and minutes in an hour. Calculate total minutes in a week.
	const daysInWeek = 7
	const hoursInDay = 24
	const minutesInHour = 60

	totalMinutesInWeek := daysInWeek * hoursInDay * minutesInHour

	fmt.Printf("Total minutes in a week: %d mins\n", totalMinutesInWeek)

	// Exercises 2: Convert an integer to float64, multiply by 2.5, then convert back to int. Print each step.
	var intNum int = 10
	var floatNum float64 = float64(intNum)
	var resultNum int = int(floatNum * 2.1)

	fmt.Printf("result is %d\n", resultNum)

	// Exercises 3: Create a constant for your birth year, calculate your age using the current year (use a variable), and print it.
	const birthYear = 1999
	currentYear := 2026
	age := currentYear - birthYear

	fmt.Printf("Age is %d\n", age)
}
