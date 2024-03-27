// describe_security_group_rules.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
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

	// Specify the security group ID
	securityGroupID := "your-security-group-id"

	// Describe security group rules
	result, err := svc.DescribeSecurityGroups(context.TODO(), &ec2.DescribeSecurityGroupsInput{
		GroupIds: []string{securityGroupID},
	})
	if err != nil {
		fmt.Println("Error describing security group rules:", err)
		return
	}

	// Print security group rules
	for _, group := range result.SecurityGroups {
		fmt.Printf("Security Group ID: %s\n", *group.GroupId)
		fmt.Printf("Group Name: %s\n", *group.GroupName)

		fmt.Println("Inbound Rules:")
		for _, permission := range group.IpPermissions {
			fmt.Printf("- Protocol: %s\n", *permission.IpProtocol)
			fmt.Printf("  Port Range: %d-%d\n", *permission.FromPort, *permission.ToPort)
			fmt.Printf("  Source: %d\n", permission.IpRanges)
		}

		fmt.Println("Outbound Rules:")
		for _, permission := range group.IpPermissionsEgress {
			fmt.Printf("- Protocol: %s\n", *permission.IpProtocol)
			fmt.Printf("  Port Range: %d-%d\n", *permission.FromPort, *permission.ToPort)
			fmt.Printf("  Destination: %d\n", permission.IpRanges)
		}

		fmt.Println("---")
	}
}
