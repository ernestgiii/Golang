package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks []string

func main() {
	for {
		fmt.Println("To-Do List")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanf("%d\n", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			var task string
			fmt.Print("Enter the task: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan() // Read the entire line, including spaces
			task = scanner.Text()
			task = strings.TrimSpace(task)

			if task == "" {
				fmt.Println("Task cannot be empty.")
			} else {
				tasks = append(tasks, task)
				fmt.Println("Task added.")
			}
		case 2:
			fmt.Println("Tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
