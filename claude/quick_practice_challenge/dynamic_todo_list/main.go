package main

import "fmt"

// Challenge: Create a simple to-do list manager:

// Start with an empty slice
// Show a menu:

// 1: Add task
// 2: View all tasks
// 3: Remove task (by index)
// 4: Exit

// Use a loop to keep the program running until user chooses to exit
// Use slices to store tasks dynamically

// Bonus: Number each task when displaying!

func main() {
	var tasks []string

	for {
		fmt.Println()
		fmt.Println("1: Add task")
		fmt.Println("2: View all tasks")
		fmt.Println("3: Remove task (by index)")
		fmt.Println("4: Exit")
		fmt.Println()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			{
				fmt.Println("Enter task:")
				var task string
				fmt.Scan(&task)
				tasks = append(tasks, task)
			}
		case 2:
			{
				fmt.Println("Tasks:")
				for i, task := range tasks {
					fmt.Println(i, task)
				}
			}
		case 3:
			{
				fmt.Println("Enter index:")
				var index int
				fmt.Scan(&index)
				tasks = append(tasks[:index], tasks[index+1:]...)
			}
		case 4:
			{
				return
			}
		default:
			{
				fmt.Println("Invalid choice")
			}
		}
	}
}
