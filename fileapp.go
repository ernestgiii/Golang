package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	directoryPath := "Golang"

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			fmt.Printf("File Name: %s\n", info.Name())
			fmt.Printf("File Size: %d bytes\n", info.Size())
			fmt.Println("---")
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning directory: %v\n", err)
	}
}
