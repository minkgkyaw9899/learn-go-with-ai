package main

import (
	"errors"
	"fmt"
)

// Example 1: Basic Error Handling
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Function with multiple error conditions
func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	if age > 150 {
		return errors.New("age seems unrealistic")
	}
	if age < 18 {
		return errors.New("must be 18 or older")
	}
	return nil // nil means no error
}

// Example 2: Custom Error Types
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %s in %s", e.Message, e.Field)
}

// Another custom error
type DatabaseError struct {
	Operation string
	Err       error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error during %s: %v", e.Operation, e.Err)
}

// Function using custom errors
func createUser(name, email string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "cannot be empty"}
	}
	if !contains(email, "@") {
		return &ValidationError{Field: "email", Message: "must contain @"}
	}
	if age < 18 {
		return &ValidationError{Field: "age", Message: "must be 18 or older"}
	}
	return nil
}

func contains(s, substr string) bool {
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Example 3: Error Wrapping (Go 1.13+)
func readConfig() error {
	return errors.New("config file not found")
}

func loadDatabase() error {
	err := readConfig()
	if err != nil {
		// Wrap the error with additional context
		return fmt.Errorf("failed to load database: %w", err)
	}
	return nil
}

func startServer() error {
	err := loadDatabase()
	if err != nil {
		return fmt.Errorf("server startup failed: %w", err)
	}
	return nil
}

// Exercises 1: Create a function withdrawMoney(balance, amount float64) (float64, error) that returns an error if amount is negative or exceeds balance.

func withdrawMoney(balance, amount float64) (float64, error) {
	if amount < 0 {
		return balance, errors.New("amount cannot be negative")
	}
	if amount > balance {
		return balance, errors.New("amount exceeds balance")
	}
	return balance - amount, nil
}

// Exercises 2: Create a custom error type FileError with fields Filename and Reason. Use it in a function that simulates file operations.

type FileError struct {
	Filename string
	Reason   string
}

func (e *FileError) Error() string {
	return fmt.Sprintf("file error %s in %s", e.Filename, e.Reason)
}

func openFile(filename string) error {
	return &FileError{Filename: filename, Reason: "file not found"}
}

// Exercises 3: Write a function parseDate(dateStr string) (day, month, year int, err error) that parses a date string "DD/MM/YYYY" and returns appropriate errors for invalid formats.

func parseDate(dateStr string) (day, month, year int, err error) {
	_, err = fmt.Sscanf(dateStr, "%d/%d/%d", &day, &month, &year)
	if err != nil {
		return 0, 0, 0, errors.New("invalid date format, expected DD/MM/YYYY")
	}

	if day < 1 || day > 31 {
		return 0, 0, 0, errors.New("day must be between 1 and 31")
	}
	if month < 1 || month > 12 {
		return 0, 0, 0, errors.New("month must be between 1 and 12")
	}

	return day, month, year, nil
}

func main() {
	// Example 1: Division
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example 2: Division by zero
	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Result:", result2)
	}

	// Example 3: Age validation
	if err := validateAge(15); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Age is valid!")
	}

	err1 := createUser("", "test@example.com", 25)
	if err1 != nil {
		fmt.Println(err1)
	}

	err3 := createUser("John", "invalid-email", 25)
	if err3 != nil {
		fmt.Println(err3)
	}

	err4 := createUser("John", "john@example.com", 15)
	if err4 != nil {
		fmt.Println(err4)
	}

	// Valid user
	err5 := createUser("John", "john@example.com", 25)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println("User created successfully!")
	}

	err6 := startServer()
	if err6 != nil {
		fmt.Println("Error:", err6)

		// Check if the error is or wraps a specific error
		if errors.Is(err6, errors.New("config file not found")) {
			fmt.Println("This is a config error!")
		}
	}
}
