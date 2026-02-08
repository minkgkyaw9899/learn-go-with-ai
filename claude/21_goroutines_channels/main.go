package main

import (
	"fmt"
	"time"
)

// Example 1: Basic Channel Usage

func sendData(ch chan string) {
	ch <- "Hello"
}

// Example 2: Channels with Multiple Values
func worker(id int, jobs chan int, result chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		result <- job * 2
	}
}

func main() {
	message := make(chan string)

	go sendData(message)

	msg := <-message
	fmt.Println(msg)

	// Example 2
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	fmt.Println("Starting workers...")

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Println("Result:", result)
	}

	// Example 3: Buffered Channels
	ch1 := make(chan int)

	ch2 := make(chan int, 3)

	// This would deadlock (uncomment to see):
	// ch1 <- 1 // No one is receiving, so this blocks forever!

	// But buffered channels allow sending without immediate receive
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3

	fmt.Println("Sent 3 values to buffered channel")

	// Now receive them
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	go func() {
		ch1 <- 1
	}()

	fmt.Println("Received from unbuffered:", <-ch1)

	// Example 4: Select Statement (Like Switch for Channels)

	ch3 := make(chan string)
	ch4 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch3 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch4 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch3:
			fmt.Println("received", msg1)
		case msg2 := <-ch4:
			fmt.Println("received", msg2)
		}
	}

	// Exercises 1: Create a goroutine that generates numbers 1-10 and sends them through a channel. The main function receives and prints them.

	ch5 := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch5 <- i
		}
		close(ch5)
	}()

	for num := range ch5 {
		fmt.Println(num)
	}

	// Exercises 2: Create a "ping-pong" program: two goroutines send messages back and forth through two channels, incrementing a counter each time. Stop after 10 exchanges.

	ping := make(chan string)
	pong := make(chan string)

	go func() {
		for range 10 {
			ping <- "ping"
			msg := <-pong
			fmt.Println("message :", msg)
		}
	}()

	go func() {
		for range 10 {
			msg := <-ping
			pong <- "pong"
			fmt.Println("message :", msg)
		}
	}()

	// Exercises 3: Implement a timeout using select and time.After(). If a channel doesn't receive data within 2 seconds, print "Timeout!"

	ch6 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch6 <- "data"
	}()

	select {
	case msg := <-ch6:
		fmt.Println("received", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}

}
