// Day 1 Challenge of building a bacic Go application  and CLI tool

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, Go!")

	// Check if a name is provided as a command-line argument
	if len(os.Args) > 1 {
		name := os.Args[1]
		fmt.Printf("Hello, %s!\n", name)
	} else {
		fmt.Println("No name provided. Please provide a name as a command-line argument.")
	}

	name := os.Args[1]
	fmt.Printf("Hello, %s! Welcome to Day 1 of the 20-day Go challenge + AWS.\n", name)
}
