package ec2

import (
	"context"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/megaproaktiv/cit"
)

var client *awsec2.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client  = awsec2.NewFromConfig(cfg)
}

// GetVPC - VPC Object
func GetVpc(stackname *string, constructID *string) (*types.Vpc, error ){
	vpcId := cit.PhysicalIDfromCID(cit.CfnClient, aws.String("vpc"), aws.String("baseVPC"))
	parms := &awsec2.DescribeVpcsInput{
		VpcIds:     []string{*vpcId},
	}
	vpcInfo,err := client.DescribeVpcs(context.TODO(), parms)
	if err != nil {
		panic("DescribeVpcs error, " + err.Error())
	}
	vpc := vpcInfo.Vpcs[0]
	return &vpc, nil
}