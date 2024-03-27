// get_iam_user_access_keys.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func main() {
	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading default config:", err)
		return
	}

	// Create an IAM client
	svc := iam.NewFromConfig(cfg)

	// Specify the IAM user name
	userName := "your-user-name"

	// List access keys for the IAM user
	result, err := svc.ListAccessKeys(context.TODO(), &iam.ListAccessKeysInput{
		UserName: &userName,
	})
	if err != nil {
		fmt.Println("Error listing access keys:", err)
		return
	}

	// Print access key details
	for _, accessKey := range result.AccessKeyMetadata {
		fmt.Printf("Access Key ID: %s\n", *accessKey.AccessKeyId)
		fmt.Printf("Status: %s\n", accessKey.Status)
		fmt.Printf("Creation Date: %s\n", accessKey.CreateDate)
		fmt.Println("---")
	}
}
