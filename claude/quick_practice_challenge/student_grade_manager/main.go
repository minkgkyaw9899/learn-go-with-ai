package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Challenge: Build a student grade management system:

// Use a map where key=student name, value=grade (int)
// Menu options:

// 1: Add student and grade
// 2: View all students
// 3: Search for a student's grade
// 4: Calculate class average
// 5: Find highest and lowest grades
// 6: Delete a student
// 7: Exit

// Handle cases where student doesn't exist
// Display results in a formatted way

// Bonus: Keep track of how many students have A (>=90), B (80-89), C (70-79), etc.

func main() {
	grades := map[string]int{}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nStudent Grade Management System")
		fmt.Println("1. Add student and grade")
		fmt.Println("2. View all students")
		fmt.Println("3. Search for a student's grade")
		fmt.Println("4. Calculate class average")
		fmt.Println("5. Find highest and lowest grades")
		fmt.Println("6. Delete a student")
		fmt.Println("7. Exit")

		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter student name: ")
			var name string
			input, _ := reader.ReadString('\n')
			name = strings.TrimSpace(input)
			fmt.Print("Enter student grade: ")
			var grade int
			fmt.Scan(&grade)
			grades[name] = grade
		case 2:
			fmt.Println("All students:")
			for name, grade := range grades {
				fmt.Printf("%s: %d\n", name, grade)
			}
		case 3:
			fmt.Print("Enter student name: ")
			var name string
			input, _ := reader.ReadString('\n')
			name = strings.TrimSpace(input)
			grade, exist := grades[name]
			if exist {
				fmt.Printf("%s's grade: %d\n", name, grade)
			} else {
				fmt.Println("Student not found")
			}
		case 4:
			var sum int
			for _, grade := range grades {
				sum += grade
			}
			fmt.Printf("Class average: %d\n", sum/len(grades))
		case 5:
			var highest int
			var lowest int
			for _, grade := range grades {
				if grade > highest {
					highest = grade
				}
				if grade < lowest {
					lowest = grade
				}
			}
			fmt.Printf("Highest grade: %d\n", highest)
			fmt.Printf("Lowest grade: %d\n", lowest)
		case 6:
			fmt.Print("Enter student name: ")
			var name string
			input, _ := reader.ReadString('\n')
			name = strings.TrimSpace(input)
			delete(grades, name)
		case 7:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
