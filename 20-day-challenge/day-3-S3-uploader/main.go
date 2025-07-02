// This code uploads a file to an S3 bucket using the AWS SDK for Go v2.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <file name here> <S3 bucket name here>")
		return
	}

	filePath := os.Args[1]
	bucketName := os.Args[2]
	fileName := filepath.Base(filePath)

	// Opens the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// Load AWS credentials and config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1")) // If your bucket is in a different region, change "us-east-1" to your bucket's region
	if err != nil {
		log.Fatalf("unable to load SDK config: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("failed to upload file: %v", err)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", fileName, bucketName)
}
