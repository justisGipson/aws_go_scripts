// list_rds_snapshots.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func main() {
	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading default config:", err)
		return
	}

	// Create an RDS client
	svc := rds.NewFromConfig(cfg)

	// List RDS database snapshots
	result, err := svc.DescribeDBSnapshots(context.TODO(), nil)
	if err != nil {
		fmt.Println("Error describing RDS snapshots:", err)
		return
	}

	// Print RDS snapshot details
	for _, snapshot := range result.DBSnapshots {
		fmt.Printf("Snapshot ID: %s\n", *snapshot.DBSnapshotIdentifier)
		fmt.Printf("Instance ID: %s\n", *snapshot.DBInstanceIdentifier)
		fmt.Printf("Snapshot Type: %d\n", snapshot.SnapshotType)
		fmt.Printf("Creation Time: %s\n", snapshot.SnapshotCreateTime)
		fmt.Println("---")
	}
}
