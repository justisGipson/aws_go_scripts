// list_cloudformation_stack_resources.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func main() {
	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading default config:", err)
		return
	}

	// Create a CloudFormation client
	svc := cloudformation.NewFromConfig(cfg)

	// Specify the stack name
	stackName := "your-stack-name"

	// List stack resources
	result, err := svc.ListStackResources(context.TODO(), &cloudformation.ListStackResourcesInput{
		StackName: &stackName,
	})
	if err != nil {
		fmt.Println("Error listing stack resources:", err)
		return
	}

	// Print stack resource details
	for _, resource := range result.StackResourceSummaries {
		fmt.Printf("Resource Type: %s\n", *resource.ResourceType)
		fmt.Printf("Logical ID: %s\n", *resource.LogicalResourceId)
		fmt.Printf("Physical ID: %s\n", *resource.PhysicalResourceId)
		fmt.Printf("Status: %s\n", resource.ResourceStatus)
		fmt.Println("---")
	}
}
