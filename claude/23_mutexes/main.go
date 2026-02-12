package main

import (
	"fmt"
	"sync"
)

// Example 1: Race Condition Problem
// func main() {
// 	counter := 0

// 	var wg sync.WaitGroup

// 	// This has a race condition!
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			counter++
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Counter:", counter) // Probably not 1000!
// }

// Example 2: Fixing with Mutex
// type SafeCounter struct {
// 	mu      sync.Mutex
// 	counter int
// }

// func (c *SafeCounter) Increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.counter++
// }

// func (c *SafeCounter) Value() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.counter
// }

// func main() {
// 	counter := SafeCounter{}
// 	var wg sync.WaitGroup

// 	for range 1000 {
// 		wg.Go(func() {
// 			counter.Increment()
// 		})
// 	}

// 	wg.Wait()
// 	fmt.Println("Counter:", counter.Value())
// }

// Example 3: RWMutex for Read/Write Optimization
// type Cache struct {
// 	mu   sync.Mutex
// 	data map[string]string
// }

// func (c *Cache) Set(key, value string) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.data[key] = value
// }

// func (c *Cache) Get(key string) (string, bool) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	value, ok := c.data[key]
// 	return value, ok
// }

// func main() {
// 	cache := Cache{
// 		data: make(map[string]string),
// 	}

// 	var wg sync.WaitGroup

// 	// One writer
// 	wg.Go(func() {
// 		for i := range 5 {
// 			cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	})

// 	// Multiple readers
// 	for i := range 10 {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			for j := range 5 {
// 				if value, ok := cache.Get(fmt.Sprintf("key%d", j)); ok {
// 					fmt.Printf("Reader %d: Got %s\n", id, value)
// 				}
// 				time.Sleep(50 * time.Millisecond)
// 			}
// 		}(i)
// 	}

// 	wg.Wait()
// }

// Exercises 1: Create a thread-safe bank account with Deposit() and Withdraw() methods that use a mutex.

// type BankAccount struct {
// 	mu      sync.Mutex
// 	balance int
// }

// func (b *BankAccount) Deposit(amount int) {
// 	b.mu.Lock()
// 	defer b.mu.Unlock()
// 	b.balance += amount
// }

// func (b *BankAccount) Withdraw(amount int) bool {
// 	b.mu.Lock()
// 	defer b.mu.Unlock()
// 	if b.balance >= amount {
// 		b.balance -= amount
// 		return true
// 	}
// 	return false
// }

// func (b *BankAccount) Balance() int {
// 	b.mu.Lock()
// 	defer b.mu.Unlock()
// 	return b.balance
// }

// func main() {
// 	account := BankAccount{}
// 	var wg sync.WaitGroup

// 	for range 1000 {
// 		wg.Go(func() {
// 			account.Deposit(1)
// 		})
// 	}

// 	wg.Wait()
// 	fmt.Println("Balance:", account.Balance())
// }

// Exercises 2: Build a concurrent hit counter for a website that tracks total hits and hits per page using a mutex.
// type HitCounter struct {
// 	mu        sync.Mutex
// 	totalHits int
// 	pageHits  map[string]int
// }

// func NewHitCounter() *HitCounter {
// 	return &HitCounter{
// 		pageHits: make(map[string]int),
// 	}
// }

// func (h *HitCounter) Hit(page string) {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()
// 	h.totalHits++
// 	h.pageHits[page]++
// }

// func (h *HitCounter) TotalHits() int {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()
// 	return h.totalHits
// }

// func (h *HitCounter) PageHits(page string) int {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()
// 	return h.pageHits[page]
// }

// func main() {
// 	counter := NewHitCounter()
// 	var wg sync.WaitGroup

// 	for range 1000 {
// 		wg.Go(func() {
// 			counter.Hit("home")
// 		})
// 	}

// 	wg.Wait()
// 	fmt.Println("Total Hits:", counter.TotalHits())
// 	fmt.Println("Home Page Hits:", counter.PageHits("home"))
// }

// Exercises 3: Create a thread-safe cache system that stores key-value pairs with expiration times. Use RWMutex.

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func main() {
	cache := NewCache()
	var wg sync.WaitGroup

	for range 1000 {
		wg.Go(func() {
			cache.Set("key1", "value1")
		})
	}

	wg.Wait()
	value, found := cache.Get("key1")
	fmt.Println("Value:", value, "Found:", found)
}
