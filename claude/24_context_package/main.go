package main

import (
	"context"
	"fmt"
	"time"
)

// Example 1: Context with Cancellation
// func queryDatabase(ctx context.Context, query string) (string, error) {
// 	resultCh := make(chan string, 1)

// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		resultCh <- "Query result: user_data"
// 	}()

// 	select {
// 	case res := <-resultCh:
// 		return res, nil
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

// 	defer cancel() // Always defer cancel to free resources!

// 	fmt.Println("Starting query...")

// 	result, err := queryDatabase(ctx, "SELECT * FROM users")

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Result:", result)
// 	}

// 	// Now try with enough time
// 	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel2()

// 	fmt.Println("\nStarting query with more time...")
// 	result2, err2 := queryDatabase(ctx2, "SELECT * FROM users")
// 	if err2 != nil {
// 		fmt.Println("Error:", err2)
// 	} else {
// 		fmt.Println(result2)
// 	}
// }

// ===========
// end of example 1
// ===========

// Example 2: Context Propagation Through Function Calls
// func handleRequest(ctx context.Context, userId int) error {
// 	fmt.Println("Handling request for user:", userId)

// 	user, err := getUser(ctx, userId)
// 	if err != nil {
// 		return fmt.Errorf("handleRequest failed: %w", err)
// 	}

// 	fmt.Println("Got user:", user)
// 	return nil
// }

// func getUser(ctx context.Context, userId int) (string, error) {
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	default:
// 		return fetchFromDB(ctx, userId)
// 	}
// }

// func fetchFromDB(ctx context.Context, userId int) (string, error) {
// 	done := make(chan string, 1)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		done <- fmt.Sprintf("user_%d", userId)
// 	}()

// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case res := <-done:
// 		return res, nil
// 	}
// }

// func main() {
// 	// Scenario 1: Request completes successfully
// 	ctx1, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	if err := handleRequest(ctx1, 1); err != nil {
// 		fmt.Println("Request 1 failed:", err)
// 	} else {
// 		fmt.Println("Request 1 succeeded!")
// 	}

// 	// Scenario 2: Request times out
// 	ctx2, cancel2 := context.WithTimeout(context.Background(), 100*time.Millisecond)
// 	defer cancel2()

// 	if err := handleRequest(ctx2, 99); err != nil {
// 		fmt.Println("Request 2 failed:", err)
// 	} else {
// 		fmt.Println("Request 2 succeeded!")
// 	}
// }

// ===========
// end of example 2
// ===========

// Example 3: Context with Values

// // Custom key type to avoid collisions
// type contextKey string

// const (
// 	UserIDKey    contextKey = "userID"
// 	RequestIDKey contextKey = "requestID"
// 	RoleKey      contextKey = "role"
// )

// func processRequest(ctx context.Context) {
// 	userID := ctx.Value(UserIDKey)
// 	requestID := ctx.Value(RequestIDKey)
// 	role := ctx.Value(RoleKey)

// 	fmt.Printf("Processing request %v for user %v with role %v\n",
// 		requestID, userID, role)

// 	// Pass to next function
// 	authorizeUser(ctx)
// }

// func authorizeUser(ctx context.Context) {
// 	role, ok := ctx.Value(RoleKey).(string)
// 	if !ok {
// 		fmt.Println("Role not found")
// 		return
// 	}

// 	switch role {
// 	case "admin":
// 		fmt.Println("Full access granted")
// 	case "user":
// 		fmt.Println("Limited access granted")
// 	default:
// 		fmt.Println("Access denied")
// 	}
// }

// func main() {
// 	// Build context with values layer by layer
// 	ctx := context.Background()
// 	ctx = context.WithValue(ctx, UserIDKey, 12345)
// 	ctx = context.WithValue(ctx, RequestIDKey, "req-abc-123")
// 	ctx = context.WithValue(ctx, RoleKey, "admin")

// 	processRequest(ctx)

// 	// Different user context
// 	ctx2 := context.Background()
// 	ctx2 = context.WithValue(ctx2, UserIDKey, 67890)
// 	ctx2 = context.WithValue(ctx2, RequestIDKey, "req-xyz-456")
// 	ctx2 = context.WithValue(ctx2, RoleKey, "user")

// 	processRequest(ctx2)
// }

// ===========
// end of example 3
// ===========

// Exercises 1: Create an HTTP server simulator where each request has a 1-second timeout. Some requests take 0.5s, some take 1.5s. Use context to cancel slow requests.
// func simulateRequest(ctx context.Context, id int, d time.Duration) {
// 	fmt.Printf("Request %d: processing for %v\n", id, d)
// 	select {
// 	case <-time.After(d):
// 		fmt.Printf("Request %d: completed successfully\n", id)
// 	case <-ctx.Done():
// 		fmt.Printf("Request %d: failed: %v\n", id, ctx.Err())
// 	}
// }

// func main() {
// 	durations := []time.Duration{500 * time.Millisecond, 1500 * time.Millisecond}

// 	for i, d := range durations {
// 		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 		simulateRequest(ctx, i+1, d)
// 		cancel()
// 		fmt.Println("---")
// 	}
// }

// Exercises 2: Build a function pipeline where each step checks context and can be cancelled. Stages: validateInput → processData → saveResult.
// func validateInput(ctx context.Context, input string) (string, error) {
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	default:
// 		return processData(ctx, input)
// 	}
// }

// func processData(ctx context.Context, input string) (string, error) {
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	default:
// 		return saveResult(ctx, input)
// 	}
// }

// func saveResult(ctx context.Context, input string) (string, error) {
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	default:
// 		return input, nil
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	result, err := validateInput(ctx, "test")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Result:", result)
// 	}
// }

// Exercises 3: Create a retry function retry(ctx context.Context, maxAttempts int, fn func() error) error that retries a failing function until it succeeds or context is cancelled.
func retry(ctx context.Context, maxAttempts int, fn func() error) error {
	for range maxAttempts {
		err := fn()
		if err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
	return fmt.Errorf("failed after %d attempts", maxAttempts)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := retry(ctx, 3, func() error {
		fmt.Println("Attempting to retry...")
		return fmt.Errorf("failed")
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
