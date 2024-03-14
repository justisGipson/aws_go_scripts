package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func main() {
	// Create an AWS session
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create a Lambda client
	svc := lambda.New(lambda.Options{
		Region: cfg.Region,
	})

	// List Lambda functions
	result, err := svc.ListFunctions(context.TODO(), &lambda.ListFunctionsInput{})
	if err != nil {
		fmt.Println("Error listing Lambda functions:", err)
		return
	}

	// Print Lambda function details
	for _, function := range result.Functions {
		fmt.Printf("Function Name: %s\n", aws.ToString(function.FunctionName))
		fmt.Printf("Runtime: %s\n", aws.ToString((*string)(&function.Runtime)))
		fmt.Printf("Memory: %d\n", aws.ToInt32(function.MemorySize))
		fmt.Printf("Handler: %s\n", aws.ToString(function.Handler))
		fmt.Println("---")
	}
}
