package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func main() {
	jsonStr := `[{"name": "John", "age": 32, "country": "USA"},
	{"name": "Bri", "age": 31, "country": "USA"}
	]`

	var people []Person

	err := json.Unmarshal([]byte(jsonStr), &people)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print information for each person
	for _, person := range people {
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
		fmt.Println("Country:", person.Country)
		fmt.Println()
	}
}
