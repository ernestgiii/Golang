package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func main() {
	// Load your AWS credentials and default region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("❌ Failed to load AWS config: %v", err)
	}
	fmt.Println("✅ AWS config loaded")

	// Create an IAM client
	client := iam.NewFromConfig(cfg)

	// Call the ListRoles API
	result, err := client.ListRoles(context.TODO(), &iam.ListRolesInput{})
	if err != nil {
		log.Fatalf("❌ Failed to list IAM roles: %v", err)
	}

	// Print the results of our user listing
	fmt.Println("👤 IAM Roles:")
	for _, role := range result.Roles {
		fmt.Printf("- %s (created on %s)\n", *role.RoleName, role.CreateDate)
	}
}
