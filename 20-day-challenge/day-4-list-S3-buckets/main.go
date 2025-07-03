// This code lists all S3 buckets in your AWS account using the AWS SDK for Go v2.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Load your AWS credentials and default region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1")) // Change "us-east-1" to your desired region
	if err != nil {
		log.Fatalf("❌ Failed to load AWS config: %v", err)
	}
	fmt.Println("✅ AWS config loaded")

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// List S3 buckets
	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("❌ Failed to list S3 buckets: %v", err)
	}

	// Print the results of our bucket listing

	fmt.Println("📦 S3 Buckets:")
	for _, bucket := range result.Buckets {
		// bucket.Name is a pointer, so we dereference it using *bucket.Name
		fmt.Printf("- %s (created on %s)\n", *bucket.Name, bucket.CreationDate.Format(time.RFC1123))
	}
	fmt.Printf("Total buckets: %d\n", len(result.Buckets))
	fmt.Println("✅ Successfully listed S3 buckets")
}
