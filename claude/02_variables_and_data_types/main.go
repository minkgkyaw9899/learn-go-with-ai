package main

import "fmt"

func main() {
	// Declare variables with var keyword and data type (uncomment following code)
	// var name string = "Julian"
	// var age int = 26
	// var salary float64 = 100000
	// var isEmployed bool = true

	// fmt.Println("Name: ", name)
	// fmt.Println("Age: ", age)
	// fmt.Println("Salary: ", salary)
	// fmt.Println("Is Employed: ", isEmployed)

	// Declare variables with short variable declaration operator (uncomment following code)
	// name := "Julian"
	// age := 26
	// salary := 100000
	// isEmployed := true

	// fmt.Println("Name: ", name)
	// fmt.Println("Age: ", age)
	// fmt.Println("Salary: ", salary)
	// fmt.Println("Is Employed: ", isEmployed)

	// multiple variable declaration (uncomment following code)
	// var name, age, salary, isEmployed = "Julian", 26, 100000, true

	// fmt.Println("Name: ", name)
	// fmt.Println("Age: ", age)
	// fmt.Println("Salary: ", salary)
	// fmt.Println("Is Employed: ", isEmployed)

	// multiple variable declaration with short variable declaration operator (uncomment following code)
	// name, age, salary, isEmployed := "Julian", 26, 100000, true

	// fmt.Println("Name: ", name)
	// fmt.Println("Age: ", age)
	// fmt.Println("Salary: ", salary)
	// fmt.Println("Is Employed: ", isEmployed)

	// multiple variable declaration with var keyword and data type (uncomment following code)
	var (
		name       string  = "Julian"
		age        int     = 26
		salary     float64 = 100000
		isEmployed bool    = true
	)

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Salary: ", salary)
	fmt.Println("Is Employed: ", isEmployed)

	// Exercises 1: Create variables for a book: title (string), pages (int), price (float64), and available (bool). Print them all.
	var title string = "The Go Programming Language"
	var pages int = 400
	var price float64 = 50.0
	var available bool = true

	fmt.Println("Title: ", title)
	fmt.Println("Pages: ", pages)
	fmt.Println("Price: ", price)
	fmt.Println("Available: ", available)

	// Exercises 2: Use short declaration to create variables for your favorite movie's name, release year, and rating.
	movieName := "The Matrix"
	releaseYear := 1999
	rating := 8.7

	fmt.Println("Movie Name: ", movieName)
	fmt.Println("Release Year: ", releaseYear)
	fmt.Println("Rating: ", rating)

	// Exercises 3: Declare three variables on one line: x, y, z all as integers with values 10, 20, 30. Print their sum.
	x, y, z := 10, 20, 30
	fmt.Println("Sum: ", x+y+z)
}
