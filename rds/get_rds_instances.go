package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func main() {
	// Create an AWS session
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create an RDS client
	svc := rds.NewFromConfig(cfg)

	// Describe RDS instances
	result, err := svc.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{})
	if err != nil {
		fmt.Println("Error describing RDS instances:", err)
		return
	}

	// Print RDS instance details
	for _, instance := range result.DBInstances {
		fmt.Printf("Instance ID: %s\n", aws.ToString(instance.DBInstanceIdentifier))
		fmt.Printf("Engine: %s\n", aws.ToString(instance.Engine))
		fmt.Printf("Status: %s\n", aws.ToString(instance.DBInstanceStatus))
		fmt.Println("---")
	}
}
