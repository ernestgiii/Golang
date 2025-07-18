// This Go program demonstrates a retry mechanism for an AWS IAM API call.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
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
	return fmt.Errorf("All %d attempts failed: %v", attempts, err)
}

func main() {
	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}
	fmt.Println("AWS config loaded")

	client := iam.NewFromConfig(cfg)

	var result *iam.ListRolesOutput

	// Retry the IAM ListRoles API call
	err = retry(3, 2*time.Second, func() error {
		output, callErr := client.ListRoles(context.TODO(), &iam.ListRolesInput{})
		if callErr == nil {
			result = output // Save result if successful
		}
		return callErr
	})

	if err != nil {
		log.Fatalf("Failed to list IAM roles after retries: %v", err)
	}

	// Print the IAM role names + creation dates
	fmt.Println("IAM Roles:")
	for _, role := range result.Roles {
		fmt.Printf("- %s (created on %s)\n", *role.RoleName, role.CreateDate)
	}
}
