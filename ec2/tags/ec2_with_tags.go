package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading default config:", err)
		return
	}

	// Create an EC2 client
	svc := ec2.NewFromConfig(cfg)

	// Define the tag key and value to filter instances
	tagKey := "Environment"
	tagValue := "Production"

	// Create a filter for the tag
	filter := types.Filter{
		Name:   aws.String(fmt.Sprintf("tag:%s", tagKey)),
		Values: []string{tagValue},
	}

	// Describe EC2 instances with the specified tag
	result, err := svc.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		Filters: []types.Filter{filter},
	})
	if err != nil {
		fmt.Println("Error describing instances:", err)
		return
	}

	// Print instance metadata
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Instance ID: %s\n", aws.ToString(instance.InstanceId))
			fmt.Printf("Instance Type: %s\n", aws.ToString((*string)(&instance.InstanceType)))
			fmt.Printf("Public IP: %s\n", aws.ToString(instance.PublicIpAddress))
			fmt.Printf("Private IP: %s\n", aws.ToString(instance.PrivateIpAddress))
			fmt.Println("---")
		}
	}
}
