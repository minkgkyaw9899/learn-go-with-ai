package main

import "fmt"

func main() {
	// Example 1: Simple If-Else
	age := 20

	if age >= 18 {
		println("You are an adult")
	} else {
		println("You are a minor")
	}

	// Example 2: If-Else If-Else
	score := 85

	if score >= 90 {
		println("Grade: A")
	} else if score >= 80 {
		println("Grade: B")
	} else if score >= 70 {
		println("Grade: C")
	} else {
		println("Grade: F")
	}

	// Example 3: Short Statement with If
	if age := 25; age >= 18 {
		fmt.Println("Age:", age, "- You can vote!")
	} else {
		fmt.Println("Age:", age, "- Too young to vote")
	}

	// Exercises 1: Ask for a number and print whether it's positive, negative, or zero.
	fmt.Println("Enter a number:")
	var num int
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

	// Exercise 2: Create a BMI calculator that takes weight (kg) and height (m), calculates BMI = weight/(height*height), and classifies it (Underweight <18.5, Normal 18.5-24.9, Overweight 25-29.9, Obese >=30).
	fmt.Println("Enter your weight in kg:")
	var weight float64
	fmt.Scan(&weight)

	fmt.Println("Enter your height in meters:")
	var height float64
	fmt.Scan(&height)

	bmi := weight / (height * height)

	if bmi < 18.5 {
		fmt.Println("Underweight")
	} else if bmi < 25 {
		fmt.Println("Normal")
	} else if bmi < 30 {
		fmt.Println("Overweight")
	} else {
		fmt.Println("Obese")
	}

	// Exercises 3: Write a password checker: length >= 8 characters (use len(password)) and print "Strong" or "Weak".
	fmt.Println("Enter your password:")
	var password string
	fmt.Scan(&password)

	if len(password) >= 8 {
		fmt.Println("Strong")
	} else {
		fmt.Println("Weak")
	}
}
