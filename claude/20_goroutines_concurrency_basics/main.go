package main

import (
	"fmt"
	"time"
)

// Example 1: Basic Goroutines

func sayHello() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func sayWorld() {
	for i := 1; i <= 5; i++ {
		fmt.Println("World", i)
		time.Sleep(200 * time.Millisecond)
	}
}

// Example 3: Real-World Example - Concurrent Web Scraping Simulation

type WebPage struct {
	URL   string
	Title string
}

func fetchPage(url string) WebPage {
	// Simulate network delay
	time.Sleep(time.Duration(100+len(url)*10) * time.Millisecond)

	return WebPage{
		URL:   url,
		Title: "Page: " + url,
	}
}

func main() {
	// // Without goroutines - runs sequentially
	// fmt.Println("--- Sequential Execution ---")
	// sayHello()
	// sayWorld()

	// fmt.Println("\n--- Concurrent Execution ---")

	// // With goroutines - runs concurrently
	// go sayHello()
	// go sayWorld()

	// // Example 2: Goroutines with Anonymous Functions
	// // Launch multiple goroutines
	// for i := 1; i <= 5; i++ {
	// 	go func(n int) {
	// 		fmt.Printf("Goroutine %d starting\n", n)
	// 		time.Sleep(time.Duration(n*100) * time.Millisecond)
	// 		fmt.Printf("Goroutine %d finished\n", n)
	// 	}(i) // Pass i as argument to avoid closure issues
	// }

	// // Wait for goroutines to finish
	// // (We'll learn a better way with channels soon)
	// time.Sleep(2 * time.Second)

	// fmt.Println("Main function ending")

	// Example 3

	// urls := []string{
	// 	"example.com",
	// 	"golang.org",
	// 	"github.com",
	// 	"stackoverflow.com",
	// 	"reddit.com",
	// }

	// fmt.Println("--- Sequential Fetching ---")

	// start := time.Now()

	// for _, url := range urls {
	// 	page := fetchPage(url)
	// 	fmt.Printf("Fetched: %s - %s\n", page.URL, page.Title)
	// }

	// fmt.Printf("Time taken Sequential Fetching: %v\n\n", time.Since(start))

	// fmt.Println("--- Concurrent Fetching ---")
	// start = time.Now()

	// for _, url := range urls {
	// 	go func(u string) {
	// 		page := fetchPage(u)
	// 		fmt.Printf("Fetched: %s - %s\n", page.URL, page.Title)
	// 	}(url)
	// }

	// // Wait for all goroutines to complete
	// time.Sleep(1 * time.Second)

	// fmt.Printf("Time taken Concurrent Fetching: %v\n", time.Since(start))

	// Exercises 1
	go countTo(5, "A")
	go countTo(5, "B")
	go countTo(5, "C")

	// Wait for goroutines to finish
	time.Sleep(1 * time.Second)

	// Exercises 2
	go run("A", time.Millisecond*100)
	go run("B", time.Millisecond*200)
	go run("C", time.Millisecond*300)
	go run("D", time.Millisecond*400)
	go run("E", time.Millisecond*500)

	// Wait for goroutines to finish
	time.Sleep(1 * time.Second)

	// Exercises 3
	for i := 1; i <= 10; i++ {
		go downloadFile(fmt.Sprintf("file%d.txt", i))
	}

	// Wait for goroutines to finish
	time.Sleep(1 * time.Second)
}

// Exercises 1: Create a function countTo(n int, name string) that counts from 1 to n with a small delay. Run 3 instances concurrently with different values.

func countTo(n int, name string) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(time.Millisecond * 100)
	}
}

// Exercises 2: Simulate a race: create 5 goroutines that each "run" for a random time (use time.Sleep with different durations), and print when each finishes.

func run(name string, duration time.Duration) {
	time.Sleep(duration)
	fmt.Printf("%s finished after %v\n", name, duration)
}

// Exercises 3: Create a program that downloads 10 "files" concurrently (simulate with time.Sleep). Print progress messages from each goroutine.

func downloadFile(filename string) {
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("Downloaded: %s\n", filename)
}
