package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// func readFile() {
// 	file, err := os.Open("output.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Read line by line using scanner
// 	scanner := bufio.NewScanner(file)
// 	lineNum := 1

// 	fmt.Println("\nReading file:")
// 	for scanner.Scan() {
// 		fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
// 		lineNum++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error reading file:", err)
// 	}
// }

// func writeFile() {
// 	file, err := os.Create("output.txt")
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	file.WriteString("Hello, File!\n")
// 	file.WriteString("This is line 2\n")

// 	// Using fmt.Fprintf for formatted writing
// 	fmt.Fprintf(file, "This is line %d\n", 3)

// 	fmt.Println("File written successfully!")
// }

// func appendFile() {
// 	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	file.WriteString("This is an appended line.\n")
// 	fmt.Println("Appended to file successfully!")
// }

// func main() {
// 	writeFile()

// 	appendFile()

// 	readFile()

// 	// Clean up
// 	os.Remove("output.txt")
// }

type Student struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Grade   float64  `json:"grade"`
	Courses []string `json:"courses"`
}

func saveStudents(fileName string, students []Student) error {
	data, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("writing file failed: %w", err)
	}

	return nil
}

func loadStudents(filename string) ([]Student, error) {
	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading file failed: %w", err)
	}

	// Unmarshal JSON
	var students []Student
	err = json.Unmarshal(data, &students)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling failed: %w", err)
	}

	return students, nil
}

func main() {
	students := []Student{
		{ID: 1, Name: "Alice", Grade: 95.5,
			Courses: []string{"Math", "Physics"}},
		{ID: 2, Name: "Bob", Grade: 87.3,
			Courses: []string{"English", "History"}},
		{ID: 3, Name: "Charlie", Grade: 92.1,
			Courses: []string{"Biology", "Chemistry"}},
	}

	const filename = "students.json"

	// Save to file
	if err := saveStudents(filename, students); err != nil {
		fmt.Println("Error saving:", err)
		return
	}
	fmt.Println("Students saved!")

	// Load from file
	loaded, err := loadStudents(filename)
	if err != nil {
		fmt.Println("Error loading:", err)
		return
	}

	fmt.Println("\nLoaded students:")
	for _, s := range loaded {
		fmt.Printf("  [%d] %s - Grade: %.1f - Courses: %v\n",
			s.ID, s.Name, s.Grade, s.Courses)
	}

	os.Remove(filename)
}
