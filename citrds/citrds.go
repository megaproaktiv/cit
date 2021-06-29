package citrds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/megaproaktiv/cit"
)

var client RdsInterface
//go:generate moq -out rds_moq_test.go . RdsInterface

type RdsInterface interface {
	DescribeDBInstances(ctx context.Context,
		 params *rds.DescribeDBInstancesInput, 
		 optFns ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error)
}

func init() {
	autoinit := cit.Autoinit()
	if autoinit {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("configuration error, " + err.Error())
		}
		client = rds.NewFromConfig(cfg)
	}
}