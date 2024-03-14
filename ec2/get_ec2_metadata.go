package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	// Create an AWS session
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create an EC2 client
	svc := ec2.New(ec2.Options{
		Region: cfg.Region,
	})

	// Describe EC2 instances
	result, err := svc.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
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
