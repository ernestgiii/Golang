package main

import (
	"encoding/json"
	"fmt"
)

type ListBucket struct {
	Buckets []struct {
		Name         string `json:"name"`
		CreationDate string `json:"creationDate"`
	} `json:"buckets"`
	Owner struct {
		DisplayName string `json:"displayName"`
		ID          string `json:"id"`
	} `json:"owner"`
}

func main() {
	data := `{
		"buckets": [
			{
				"name": "demo-objectlock1",
				"creationDate": "2023-01-01T00:00:00Z"
			},
			{
				"name": "dev-irsa-test-bucket",
				"creationDate": "2023-02-01T00:00:00Z"
			}
		],
		"owner": {
			"displayName": "John Doe",
			"id": "12345"
		}
	}`

	var response ListBucket
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Buckets:")
	for _, bucket := range response.Buckets {
		fmt.Printf("Name: %s, Creation Date: %s\n", bucket.Name, bucket.CreationDate)
	}
	fmt.Printf("Owner: %s (ID: %s)\n", response.Owner.DisplayName, response.Owner.ID)
}
