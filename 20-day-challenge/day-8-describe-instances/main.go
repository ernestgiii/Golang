// This script describes EC2 instances

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func retry(attempts int, sleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		log.Printf(" Attempt %d failed: %v", i+1, err)
		time.Sleep(sleep)
	}
	return fmt.Errorf(" all %d attempts failed: %v", attempts, err)
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf(" Failed to load AWS config: %v", err)
	}
	fmt.Println(" AWS config loaded")

	client := ec2.NewFromConfig(cfg)

	var result *ec2.DescribeInstancesOutput

	// Use retry wrapper for DescribeInstances
	err = retry(3, 2*time.Second, func() error {
		output, callErr := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
		if callErr == nil {
			result = output
		}
		return callErr
	})

	if err != nil {
		log.Fatalf(" Failed to describe EC2 instances: %v", err)
	}

	//  Print instance info
	fmt.Println(" EC2 Instances:")
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instanceID := *instance.InstanceId
			state := string(instance.State.Name)

			// Get the Name tag if it exists
			name := "(no name)"
			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
					break
				}
			}

			fmt.Printf("- %s | %s | %s\n", instanceID, name, state)
		}
	}
}
