// This script generates a pre-signed URL for an S3 object and downloads it to your local machine for example a png file.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <bucket-name> <object-key>")
		return
	}

	bucket := os.Args[1]
	key := os.Args[2]

	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	fmt.Println(" AWS configuration loaded successfully")

	client := s3.NewFromConfig(cfg)
	// Create a presigner
	presigner := s3.NewPresignClient(client)

	// Create the presigned URL request

	urlOut, err := presigner.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(1*time.Hour))

	if err != nil {
		log.Fatalf(" Failed to generate pre-signed URL: %v", err)
	}

	fmt.Println(" Pre-signed URL generated successfully:")
	fmt.Println(urlOut.URL)
	fmt.Println("You can use this URL to access the object in S3 for the next hour.")

}
