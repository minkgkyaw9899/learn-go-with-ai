package main

import (
	"fmt"
	"sync"
	"time"
)

// Lesson 22: WaitGroups (Better Goroutine Synchronization)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func fetchURL(url string, result chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(len(url)*10) * time.Millisecond)
	result <- fmt.Sprintf("Fetched: %s", url)
}

type Image struct {
	Name     string
	Size     int
	Filtered bool
}

func applyFilter(img *Image, filterName string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Applying %s filter to %s...\n", filterName, img.Name)
	time.Sleep(time.Duration(img.Size*100) * time.Millisecond)
	img.Filtered = true
	fmt.Printf("Finished %s filter on %s\n", filterName, img.Name)
}

func main() {
	// Example 1
	// wg := sync.WaitGroup{}

	// for i := 1; i <= 5; i++ {
	// 	wg.Add(1)
	// 	go worker(i, &wg)
	// }

	// wg.Wait()
	// fmt.Println("All workers completed")

	// Example 2
	// urls := []string{
	// 	"example.com",
	// 	"golang.org",
	// 	"github.com",
	// 	"stackoverflow.com",
	// }

	// results := make(chan string, len(urls))

	// wg := sync.WaitGroup{}

	// for _, url := range urls {
	// 	wg.Add(1)
	// 	go fetchURL(url, results, &wg)
	// }

	// go func() {
	// 	wg.Wait()
	// 	close(results)
	// }()

	// // Receive results
	// for result := range results {
	// 	fmt.Println(result)
	// }

	// Example 3
	// images := []Image{
	// 	{Name: "photo1.jpg", Size: 5},
	// 	{Name: "photo2.jpg", Size: 3},
	// 	{Name: "photo3.jpg", Size: 7},
	// }

	// var wg sync.WaitGroup
	// start := time.Now()

	// for i := range images {
	// 	wg.Add(3)
	// 	go applyFilter(&images[i], "Blur", &wg)
	// 	go applyFilter(&images[i], "Sharpen", &wg)
	// 	go applyFilter(&images[i], "Sepia", &wg)
	// }

	// wg.Wait()

	// fmt.Printf("\nAll processing completed in %v\n", time.Since(start))

	// for _, img := range images {
	// 	fmt.Printf("%s - Filtered: %v\n", img.Name, img.Filtered)
	// }

	// Exercises 1: Create a program that spawns 10 goroutines, each printing numbers 1-5. Use WaitGroup to ensure all complete before main exits.
	// wg := sync.WaitGroup{}

	// for i := 1; i <= 10; i++ {
	// 	wg.Add(1)
	// 	go func(id int) {
	// 		defer wg.Done()
	// 		for j := 1; j <= 5; j++ {
	// 			fmt.Printf("Goroutine %d: %d\n", id, j)
	// 		}
	// 	}(i)
	// }
	// wg.Wait()

	// Exercises 2: Build a concurrent calculator: create goroutines for add, subtract, multiply, divide operations on the same two numbers. Use WaitGroup and channels to collect results.
	// calculator(10, 5)

	// Exercises 3: Simulate a download manager: create 5 files to "download" concurrently, use WaitGroup to wait for all, and display total time.
	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
		"file4.txt",
		"file5.txt",
	}

	var wg sync.WaitGroup
	start := time.Now()

	for _, file := range files {
		wg.Add(1)
		go downloadFile(file, &wg)
	}

	wg.Wait()

	fmt.Printf("\nAll downloads completed in %v\n", time.Since(start))
}

// Exercises 2: Build a concurrent calculator: create goroutines for add, subtract, multiply, divide operations on the same two numbers. Use WaitGroup and channels to collect results.
func calculator(a, b float64) {
	result := make(chan string, 4)

	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer wg.Done()
		result <- fmt.Sprintf("Add: %.2f + %.2f = %.2f", a, b, a+b)
	}()

	go func() {
		defer wg.Done()
		result <- fmt.Sprintf("Subtract: %.2f - %.2f = %.2f", a, b, a-b)
	}()

	go func() {
		defer wg.Done()
		result <- fmt.Sprintf("Multiply: %.2f * %.2f = %.2f", a, b, a*b)
	}()

	go func() {
		defer wg.Done()
		result <- fmt.Sprintf("Divide: %.2f / %.2f = %.2f", a, b, a/b)
	}()

	wg.Wait()
	close(result)

	for res := range result {
		fmt.Println(res)
	}
}

// Exercises 3: Simulate a download manager: create 5 files to "download" concurrently, use WaitGroup to wait for all, and display total time.

func downloadFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Downloading %s...\n", filename)
	time.Sleep(time.Second)
	fmt.Printf("Finished downloading %s\n", filename)
}
