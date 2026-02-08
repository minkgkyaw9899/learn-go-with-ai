package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Challenge: Build a concurrent file processing simulation that demonstrates goroutines and channels.
// Requirements:
// Create a program that simulates processing multiple files concurrently:

// File struct:

// type File struct {
//     Name string
//     Size int  // in MB
// }

// Functions to implement:

// generateFiles(count int) []File - creates dummy files
// processFile(file File, results chan string) - simulates processing (use time.Sleep based on size)
// worker(id int, files chan File, results chan string) - worker that processes files from channel

// Program flow:

// Generate 20 files with random sizes (1-10 MB)
// Create 5 worker goroutines
// Send all files through a channel to workers
// Workers process files concurrently
// Collect and display results showing which worker processed which file

// Output should show:

// Which worker is processing which file
// How long each file took to process
// Total time for all files
// Comparison with sequential processing time

// Bonus: Add a progress tracker that shows how many files have been processed out of the total.

type File struct {
	Name string
	Size int // in MB
}

func generateFiles(count int) []File {
	files := make([]File, count)
	for i := range count {
		files[i] = File{
			Name: fmt.Sprintf("file_%d.dat", i),
			Size: rand.Intn(10) + 1,
		}
	}
	return files
}

// processFile(file File, results chan string) - simulates processing (use time.Sleep based on size)
func processFile(file File, results chan string) {
	startTime := time.Now()
	// Simulate processing time: 100ms per MB
	time.Sleep(time.Duration(file.Size) * 100 * time.Millisecond)
	duration := time.Since(startTime)
	results <- fmt.Sprintf("Processed %s (%d MB) in %v", file.Name, file.Size, duration)
}

func worker(id int, files chan File, results chan string) {
	for file := range files {
		// Bonus: Progress tracking is simulated by the result message including the worker ID
		// which allows the main function to display who did what.
		// For the bonus requirement "Add a progress tracker that shows how many files have been processed out of the total"
		// we can handle that in the main loop when receiving results.
		startTime := time.Now()
		// Simulate processing time: 100ms per MB
		time.Sleep(time.Duration(file.Size) * 100 * time.Millisecond)
		duration := time.Since(startTime)
		results <- fmt.Sprintf("Worker %d processed %s (%d MB) in %v", id, file.Name, file.Size, duration)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	files := generateFiles(20)
	totalFiles := len(files)

	// Calculate sequential processing time (approximation)
	var sequentialTime time.Duration
	for _, f := range files {
		sequentialTime += time.Duration(f.Size) * 100 * time.Millisecond
	}

	fmt.Printf("Starting concurrent processing of %d files...\n", totalFiles)
	fmt.Printf("Estimated sequential processing time: %v\n\n", sequentialTime)

	// Create channels
	filesChan := make(chan File, totalFiles)
	resultsChan := make(chan string, totalFiles)

	// Start workers
	numWorkers := 5
	for i := 1; i <= numWorkers; i++ {
		go worker(i, filesChan, resultsChan)
	}

	// Send files to workers
	startTime := time.Now()
	for _, f := range files {
		filesChan <- f
	}
	close(filesChan)

	// Collect results
	for i := 0; i < totalFiles; i++ {
		result := <-resultsChan
		fmt.Printf("[%d/%d] %s\n", i+1, totalFiles, result)
	}

	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal concurrent processing time: %v\n", totalTime)
	fmt.Printf("Speedup factor: %.2fx\n", float64(sequentialTime)/float64(totalTime))
}
