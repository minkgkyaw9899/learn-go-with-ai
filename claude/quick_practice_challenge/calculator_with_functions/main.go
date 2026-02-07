package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Challenge: Build a complete calculator using functions:
// Required Functions:

// add(a, b float64) float64
// subtract(a, b float64) float64
// multiply(a, b float64) float64
// divide(a, b float64) (float64, error) - handle division by zero
// power(base, exponent float64) float64
// squareRoot(n float64) (float64, error) - handle negative numbers
// calculate(a, b float64, operator string) (float64, error) - main calculator function

// Menu:

// Basic operations (+, -, *, /)
// Power
// Square root
// Exit

// Bonus: Keep history of calculations using a slice of strings!

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Error, can't divide by 0")
	}
	return a / b, nil
}

func power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("Error, can't take square root of negative number")
	}
	return math.Sqrt(n), nil
}

func calculate(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return add(a, b), nil
	case "-":
		return subtract(a, b), nil
	case "*":
		return multiply(a, b), nil
	case "/":
		return divide(a, b)
	case "^":
		return power(a, b), nil
	case "sqrt":
		return squareRoot(a)
	default:
		return 0, fmt.Errorf("Error, invalid operator")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var history []string
	for {

		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Power")
		fmt.Println("6. Square Root")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var a, b float64
			fmt.Print("Enter first number: ")
			s1, _ := reader.ReadString('\n')
			a, _ = strconv.ParseFloat(strings.TrimSpace(s1), 64)
			fmt.Print("Enter second number: ")
			s2, _ := reader.ReadString('\n')
			b, _ = strconv.ParseFloat(strings.TrimSpace(s2), 64)
			result := add(a, b)
			fmt.Println("Result:", result)
			history = append(history, fmt.Sprintf("%f + %f = %f", a, b, result))

		case 2:
			var a, b float64
			fmt.Print("Enter first number: ")
			s1, _ := reader.ReadString('\n')
			a, _ = strconv.ParseFloat(strings.TrimSpace(s1), 64)
			fmt.Print("Enter second number: ")
			s2, _ := reader.ReadString('\n')
			b, _ = strconv.ParseFloat(strings.TrimSpace(s2), 64)
			result := subtract(a, b)
			fmt.Println("Result:", result)
			history = append(history, fmt.Sprintf("%f - %f = %f", a, b, result))
		case 3:
			var a, b float64
			fmt.Print("Enter first number: ")
			s1, _ := reader.ReadString('\n')
			a, _ = strconv.ParseFloat(strings.TrimSpace(s1), 64)
			fmt.Print("Enter second number: ")
			s2, _ := reader.ReadString('\n')
			b, _ = strconv.ParseFloat(strings.TrimSpace(s2), 64)
			result := multiply(a, b)
			fmt.Println("Result:", result)
			history = append(history, fmt.Sprintf("%f * %f = %f", a, b, result))
		case 4:
			var a, b float64
			fmt.Print("Enter first number: ")
			s1, _ := reader.ReadString('\n')
			a, _ = strconv.ParseFloat(strings.TrimSpace(s1), 64)
			fmt.Print("Enter second number: ")
			s2, _ := reader.ReadString('\n')
			b, _ = strconv.ParseFloat(strings.TrimSpace(s2), 64)
			result, err := divide(a, b)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Result:", result)
				history = append(history, fmt.Sprintf("%f / %f = %f", a, b, result))
			}
		case 5:
			var base, exponent float64
			fmt.Print("Enter base: ")
			s1, _ := reader.ReadString('\n')
			base, _ = strconv.ParseFloat(strings.TrimSpace(s1), 64)
			fmt.Print("Enter exponent: ")
			s2, _ := reader.ReadString('\n')
			exponent, _ = strconv.ParseFloat(strings.TrimSpace(s2), 64)
			result := power(base, exponent)
			fmt.Println("Result:", result)
			history = append(history, fmt.Sprintf("%f ^ %f = %f", base, exponent, result))
		case 6:
			var n float64
			fmt.Print("Enter a number: ")
			s1, _ := reader.ReadString('\n')
			n, _ = strconv.ParseFloat(strings.TrimSpace(s1), 64)
			result, err := squareRoot(n)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Result:", result)
				history = append(history, fmt.Sprintf("sqrt(%f) = %f", n, result))
			}
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
