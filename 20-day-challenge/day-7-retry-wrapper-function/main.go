// This is a Go program that implements a retry mechanism for a function that may fail.
package main

import (
	"fmt"
	"log"
	"time"
)

// Retry wrapper function

func retry(attempts int, sleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		log.Printf("Attempt %d failed: %v", i+1, err)
		time.Sleep(sleep)
	}
	return fmt.Errorf("after %d attempts, last error: %v", attempts, err)
}

func main() {
	counter := 0

	// Simulate a task that fails twice before succeeding
	err := retry(5, 2*time.Second, func() error {
		if counter < 2 {
			counter++
			return fmt.Errorf("simulated error")
		}
		fmt.Printf("✅ Success on attempt #%d\n", counter+1)
		return nil
	})

	if err != nil {
		log.Fatalf("All attempts failed: %v", err)
	}
}
