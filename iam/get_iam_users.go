package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func main() {
	// Create an AWS session
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create an IAM client
	svc := iam.New(iam.Options{
		Region: cfg.Region,
	})

	// List IAM users
	result, err := svc.ListUsers(context.TODO(), &iam.ListUsersInput{})
	if err != nil {
		fmt.Println("Error listing IAM users:", err)
		return
	}

	// Print IAM user details
	for _, user := range result.Users {
		fmt.Printf("User Name: %s\n", aws.ToString(user.UserName))
		fmt.Printf("User ARN: %s\n", aws.ToString(user.Arn))
		fmt.Println("---")
	}
}
