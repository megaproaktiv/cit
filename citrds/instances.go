package citrds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"

	"github.com/megaproaktiv/cit"
)

func GetDbInstances(stackname *string, constructID *string)( *[]types.DBInstance, error){
	id,err := cit.PhysicalIDfromCID(cit.CfnClient,stackname, constructID)
	if err != nil {
		return nil, err
	}
	parms := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: id,
	}
	info, err := client.DescribeDBInstances(context.TODO(), parms)
	if err != nil {
		return nil, err
	}

	return &info.DBInstances, nil
}

func SetClient(c RdsInterface) {
	client = c
}