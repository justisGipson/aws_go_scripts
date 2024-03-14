package s3buckets

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Create an AWS session
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create an S3 client
	svc := s3.New(s3.Options{
		Region: cfg.Region,
	})

	// List S3 buckets
	result, err := svc.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Println("Error listing buckets:", err)
		return
	}

	// Print bucket names
	for _, bucket := range result.Buckets {
		fmt.Println(aws.ToString(bucket.Name))
	}
}
