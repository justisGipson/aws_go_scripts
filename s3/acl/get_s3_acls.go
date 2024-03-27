// get_s3_bucket_acl.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading default config:", err)
		return
	}

	// Create an S3 client
	svc := s3.NewFromConfig(cfg)

	// Specify the bucket name
	bucketName := "" // replace with the name of the bucket

	// Get the bucket ACL
	result, err := svc.GetBucketAcl(context.TODO(), &s3.GetBucketAclInput{
		Bucket: &bucketName,
	})
	if err != nil {
		fmt.Println("Error getting bucket ACL:", err)
		return
	}

	// Print the bucket ACL
	fmt.Printf("Bucket ACL for %s:\n", bucketName)
	for _, grant := range result.Grants {
		fmt.Printf("- Grantee: %s\n", *grant.Grantee.DisplayName)
		fmt.Printf("  Permission: %s\n", grant.Permission)
	}
}
