package main

import "fmt"

func main() {
	// method 1: using 'var' keyword
	var name string = "Julian"
	var age int = 25
	var country string = "Myanmar"

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Country: ", country)

	var name2 = "Ko Ko"
	var age2 = 20
	fmt.Println("Name2: ", name2)
	fmt.Println("Age2: ", age2)

	// method 2: using ':= ' operator
	name3 := "Aung Aung"
	age3 := 30
	country3 := "Myanmar"
	fmt.Println("Name3: ", name3)
	fmt.Println("Age3: ", age3)
	fmt.Println("Country3: ", country3)

	// method 3: using 'const' keyword
	const name4 string = "Ko Ko"
	const age4 int = 20
	const country4 string = "Myanmar"
	fmt.Println("Name4: ", name4)
	fmt.Println("Age4: ", age4)
	fmt.Println("Country4: ", country4)

	// data type: string, int, float64, bool
	var city string = "Yangon"
	var temperature1 float64 = 25.5
	var isRaining bool = false
	fmt.Println("City: ", city)
	fmt.Println("Temperature: ", temperature1)
	fmt.Println("Is Raining: ", isRaining)

	// update variable

	name5 := "Kyaw Kyaw"
	fmt.Println("Name5: ", name5)

	name5 = "Aung Aung"
	fmt.Println("Name5: ", name5)

	// Exercise 1: The Profile Creator Create three variables using the Short Declaration (:=) method inside main:
	firstName := "Min"
	currentAge := 26
	isStudent := false

	fmt.Println("First Name: ", firstName)
	fmt.Println("Current Age: ", currentAge)
	fmt.Println("Is Student: ", isStudent)

	// Exercise 2: The Temperature Reporter Use the Long Declaration (var) method to create a variable named temperature with a type of float64 (e.g., 32.5). Print it.
	var temperature float64 = 32.5
	fmt.Println("Temperature: ", temperature)

	//	Exercise 3: The Swap
	//
	// Create a variable fruit and set it to "Apple".
	// Print fruit.
	// Change the value of fruit to "Orange".
	// Print fruit again.
	fruit := "Apple"
	fmt.Println("Fruit: ", fruit)
	fruit = "Orange"
	fmt.Println("Fruit: ", fruit)
}
