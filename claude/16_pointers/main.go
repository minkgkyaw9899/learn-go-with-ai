package main

import "fmt"

func doubleValue(n int) {
	n = n * 2
}

func doublePointer(n *int) {
	*n = *n * 2
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	// Example 1: Basic Pointers
	x := 10
	p := &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of p:", p)
	fmt.Println("Value at address p:", *p)

	// Example 2: Pointers in Functions
	num := 10

	doubleValue(num)
	fmt.Println("After doubleValue:", num) // Still 10

	doublePointer(&num)
	fmt.Println("After doublePointer:", num)

	// Example 3: Swapping with Pointers
	x, y := 5, 10
	fmt.Printf("Before: x=%d, y=%d\n", x, y)

	swap(&x, &y)
	fmt.Printf("After: x=%d, y=%d\n", x, y)

	// Exercises 1: Create a function increment(n *int) that increases a number by 1 using pointers.
	num1 := 10
	increment(&num1)
	fmt.Println("After increment:", num1)

	// Exercises 2: Write a function updatePrice(price *float64, percentage float64) that increases price by a percentage.
	price := 100.0
	updatePrice(&price, 10)
	fmt.Println("After updatePrice:", price)

	// Exercises 3: Create a function reset(values *[]int) that sets all slice elements to zero using pointers.
	values := []int{1, 2, 3, 4, 5}
	reset(&values)
	fmt.Println("After reset:", values)
}

func increment(n *int) {
	*n += 1
}

func updatePrice(price *float64, percentage float64) {
	*price = *price * (1 + percentage/100)
}

func reset(values *[]int) {
	for i := range *values {
		(*values)[i] = 0
	}
}
