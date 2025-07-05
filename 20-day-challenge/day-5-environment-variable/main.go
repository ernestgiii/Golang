// This code is a simple command-line tool to manage environment variables in a `.env` file.
// It allows you to list, get, and set environment variables using the `godotenv` package.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const envFile = ".env"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  list                List all env vars")
		fmt.Println("  get <key>           Get value of a key")
		fmt.Println("  set <key> <value>   Set or update a key")
		return
	}

	command := os.Args[1]

	// Load existing env variables from file
	envMap, err := godotenv.Read(envFile)
	if err != nil {
		log.Fatalf("❌ Failed to load .env file: %v", err)
	}

	switch command {

	case "list":
		fmt.Println("📜 Environment Variables:")
		for k, v := range envMap {
			fmt.Printf("%s=%s\n", k, v)
		}

	case "get":
		if len(os.Args) < 3 {
			fmt.Println("❗ Usage: get <key>")
			return
		}

		key := os.Args[2]
		value, exists := envMap[key]
		if exists {
			fmt.Printf("%s=%s\n", key, value)
		} else {
			fmt.Printf("❌ Key '%s' not found.\n", key)
		}
	case "set":
		if len(os.Args) < 4 {
			fmt.Println("❗ Usage: set <key> <value>")
			return
		}
		key := os.Args[2]
		value := os.Args[3]

		// Update or add
		envMap[key] = value

		// Save the updated map back to the file
		err := godotenv.Write(envMap, envFile)
		if err != nil {
			log.Fatalf("❌ Failed to write to .env: %v", err)
		}
		fmt.Printf("✅ %s set to %s\n", key, value)

	default:
		fmt.Println("❌ Unknown command. Try: list, get, or set.")
	}
}
