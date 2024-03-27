// describe_vpc_subnets.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
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

	// Specify the VPC ID
	vpcID := "your-vpc-id"

	// Describe VPC subnets
	result, err := svc.DescribeSubnets(context.TODO(), &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{vpcID},
			},
		},
	})
	if err != nil {
		fmt.Println("Error describing VPC subnets:", err)
		return
	}

	// Print subnet details
	for _, subnet := range result.Subnets {
		fmt.Printf("Subnet ID: %s\n", *subnet.SubnetId)
		fmt.Printf("CIDR Block: %s\n", *subnet.CidrBlock)
		fmt.Printf("Availability Zone: %s\n", *subnet.AvailabilityZone)
		fmt.Println("---")
	}
}
