package main

import "fmt"

// Example 1: Basic Function
func greet(name string) {
	fmt.Println("Hello", name)
}

func add(a int, b int) int {
	return a + b
}

// Example 2: Multiple Return Values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Error, can't divide by 0")
	}

	return a / b, nil
}

func swap(a, b int) (int, int) {
	return b, a
}

// Example 3: Named Return Values
func rectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

func main() {
	greet("Alice")

	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	divideResult, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("10 / 2 =", divideResult)
	}

	swapResult1, swapResult2 := swap(10, 20)
	fmt.Println("Swapped values:", swapResult1, swapResult2)

	a, p := rectangle(5, 3)
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", a, p)

	evenResult := isEven(10)
	fmt.Println("Is 10 even?", evenResult)

	maxResult := findMax(10, 20, 30)
	fmt.Println("Max value:", maxResult)

	celsiusResult := celsiusToFahrenheit(25)
	fmt.Println("25 Celsius is", celsiusResult, "Fahrenheit")
}

// Exercises 1: Create a function isEven(n int) bool that returns true if a number is even.
func isEven(n int) bool {
	return n%2 == 0
}

// Exercises 2:Write a function findMax(a, b, c int) int that returns the largest of three numbers.
func findMax(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

// Exercises 3: Create a function celsiusToFahrenheit(celsius float64) float64 that converts temperature.
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}
