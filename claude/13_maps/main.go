package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Example 1: Creating and Using Maps
	// ages := make(map[string]int)

	// ages["Alice"] = 30
	// ages["Bob"] = 25
	// ages["Charlie"] = 35

	// fmt.Println("Ages:", ages)
	// fmt.Println("Alice's age:", ages["Alice"])
	// fmt.Println("Map length:", len(ages))

	// Example 2: Map Initialization and Checking Existence
	// prices := map[string]float64{
	// 	"Apple":  1.50,
	// 	"Banana": 0.75,
	// 	"Orange": 2.00,
	// }

	// Check if key exists
	// price, exist := prices["Mango"]
	// if exist {
	// 	fmt.Println("Mango coasts: ", price)
	// } else {
	// 	fmt.Println("Mango not found in the map")
	// }

	// Access existing key
	// applePrice, _ := prices["Apple"]
	// fmt.Println("Apple costs:", applePrice)

	// Example 3: Looping and Deleting
	// students := map[string]int{
	// 	"John":  85,
	// 	"Sarah": 92,
	// 	"Mike":  78,
	// 	"Emma":  95,
	// }

	// // Loop through map
	// fmt.Println("All students:")
	// for name, score := range students {
	// 	fmt.Printf("%s: %d\n", name, score)
	// }

	// // Delete a key
	// delete(students, "Mike")
	// fmt.Println("\nAfter deleting Mike:")
	// fmt.Println(students)

	// Exercises 1: Create a map of country capitals (at least 5), then ask the user for a country and display its capital.
	// countries := map[string]string{
	// 	"Ukraine": "Kyiv",
	// 	"USA":     "Washington",
	// 	"France":  "Paris",
	// 	"Germany": "Berlin",
	// 	"Italy":   "Rome",
	// }

	// country := ""
	// fmt.Println("Enter a country name: ")
	// fmt.Scan(&country)

	// capital, exist := countries[country]
	// if exist {
	// 	fmt.Println("The capital of", country, "is", capital)
	// } else {
	// 	fmt.Println("Capital not found for", country)
	// }

	// Exercises 2: Create a phonebook: map of names to phone numbers. Add at least 3 entries, display all, then delete one.
	// phonebook := map[string]string{
	// 	"John":  "123-456-7890",
	// 	"Sarah": "987-654-3210",
	// 	"Mike":  "555-123-4567",
	// }

	// fmt.Println("Phonebook:")
	// for name, phone := range phonebook {
	// 	fmt.Printf("%s: %s\n", name, phone)
	// }

	// delete(phonebook, "John")
	// fmt.Println("\nAfter deleting John:")
	// fmt.Println(phonebook)

	// Exercises 3: Create a word counter: ask user for a sentence, count how many times each word appears using a map.
	wordCounter := map[string]int{}

	var sentence string
	fmt.Println("Enter a sentence: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	sentence = strings.TrimSpace(input)

	words := strings.Split(sentence, " ")

	for _, word := range words {
		wordCounter[word]++
	}

	fmt.Println("Word counts:")
	for word, count := range wordCounter {
		fmt.Printf("%s: %d\n", word, count)
	}
}
