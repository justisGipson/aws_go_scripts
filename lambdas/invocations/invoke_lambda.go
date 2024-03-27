// invoke_lambda_function.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func main() {
	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading default config:", err)
		return
	}

	// Create a Lambda client
	svc := lambda.NewFromConfig(cfg)

	// Specify the Lambda function name
	functionName := "" // replace with function name

	// Invoke the Lambda function
	result, err := svc.Invoke(context.TODO(), &lambda.InvokeInput{
		FunctionName: &functionName,
	})
	if err != nil {
		fmt.Println("Error invoking Lambda function:", err)
		return
	}

	// Print the Lambda function response
	fmt.Println("Lambda function response:")
	fmt.Println(string(result.Payload))
}
