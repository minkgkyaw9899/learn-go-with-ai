package main

import "fmt"

func main() {
	//Example 1: Basic Switch
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6, 7:
		fmt.Println("Weekend!")
	default:
		fmt.Println("Something is wrong")
	}

	//	Example 2: Switch with Conditions
	score := 72

	switch {
	case score >= 98:
		fmt.Println("Excellent")
	case score >= 80:
		fmt.Println("Very Good")
	case score >= 70:
		fmt.Println("Good")
	case score >= 60:
		fmt.Println("Bad")
	default:
		fmt.Println("Failed")
	}

	// Example 3: Switch with Short Statement
	switch hour := 14; {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 18:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	// Exercises 1: Create a month number to name converter (1=January, 2=February, etc.) using switch.
	switch month := 1; {
	case month == 1:
		fmt.Println("January")
	case month == 2:
		fmt.Println("February")
	case month == 3:
		fmt.Println("March")
	case month == 4:
		fmt.Println("April")
	case month == 5:
		fmt.Println("May")
	case month == 6:
		fmt.Println("June")
	case month == 7:
		fmt.Println("July")
	case month == 8:
		fmt.Println("August")
	case month == 9:
		fmt.Println("September")
	case month == 10:
		fmt.Println("October")
	case month == 11:
		fmt.Println("November")
	case month == 12:
		fmt.Println("December")
	default:
		fmt.Println("Unknown")
	}

	// Exercises 2: Build a simple calculator that asks for two numbers and an operator (+, -, *, /) and performs the operation using switch.
	var num1, num2 float64
	var operator string

	fmt.Println("Enter first number:")
	fmt.Scan(&num1)

	fmt.Println("Enter second number:")
	fmt.Scan(&num2)

	fmt.Println("Enter operator (+, -, *, /):")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println("Result:", num1+num2)
	case "-":
		fmt.Println("Result:", num1-num2)
	case "*":
		fmt.Println("Result:", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Result:", num1/num2)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Error: Invalid operator")
	}

	// Exercises 3: Write a traffic light program: input color (red/yellow/green) and print the action (Stop/Slow Down/Go).
	var color string
	fmt.Println("Enter traffic light color (red/yellow/green):")
	fmt.Scan(&color)

	switch color {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Slow Down")
	case "green":
		fmt.Println("Go")
	default:
		fmt.Println("Invalid color")
	}
}
