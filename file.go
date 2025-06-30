package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating the file", err)
		return
	}
	defer file.Close()

	content := "I love DevOps!\n "
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to the file", err)
	}
	fmt.Println("Content written to 'output.txt'")
}
