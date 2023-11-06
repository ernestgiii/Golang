package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func listFilesInDirectory(directoryPath string) {
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	directoryPath := "Golang"
	listFilesInDirectory(directoryPath)
}
