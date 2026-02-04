package main

import "fmt"

// Explanation
// Comparison Operators (return true/false):

// == Equal to
// != Not equal to
// < Less than
// > Greater than
// <= Less than or equal to
// >= Greater than or equal to

// Logical Operators:

// && AND (both must be true)
// || OR (at least one must be true)
// ! NOT (reverses the boolean)

func main() {
	// Example 1: Comparison Operators
	a := 10
	b := 20

	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a <= b:", a <= b)
	fmt.Println("a >= b:", a >= b)

	// Example 2: Logical Operators
	// age := 25
	// hasLicense := true

	// AND operator
	// canDrive := age >= 18 && hasLicense
	// fmt.Println("Can drive:", canDrive)

	// OR operator
	isWeekend := false
	isHoliday := true
	canRelax := isWeekend || isHoliday
	// fmt.Println("Can relax:", canRelax)

	// NOT operator
	isWorking := !canRelax
	fmt.Println("Is working:", isWorking)

	// Example 3: Complex Conditions
	score := 85
	attendance := 90

	// Pass if score >= 60 AND attendance >= 75
	passed := score >= 60 && attendance >= 75

	// Honor roll if score >= 90 OR (score >= 80 AND attendance >= 95)
	honorRoll := score >= 90 || (score >= 80 && attendance >= 95)

	fmt.Println("Passed:", passed)
	fmt.Println("Honor Roll:", honorRoll)

	// Exercises 1: Ask for a number and check if it's even (use modulus and comparison).
	var num int

	fmt.Println("Enter a number:")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Exercises 2: Create a voting eligibility checker: age >= 18 AND citizenship == true.
	var age int
	var citizenship bool

	fmt.Println("Enter your age:")
	fmt.Scanln(&age)
	fmt.Println("Enter your citizenship:")
	fmt.Scanln(&citizenship)

	if age >= 18 && citizenship == true {
		fmt.Println("Eligible to vote")
	} else {
		fmt.Println("Not eligible to vote")
	}

	// Exercises 3: Write a program that checks if a year is a leap year (divisible by 4 AND (not divisible by 100 OR divisible by 400)).
	var year int

	fmt.Println("Enter a year:")
	fmt.Scanln(&year)

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println("Leap year")
	} else {
		fmt.Println("Not a leap year")
	}
}
