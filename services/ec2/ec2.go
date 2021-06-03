package ec2

import (
	"context"

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
	vpcId,err := cit.PhysicalIDfromCID(cit.CfnClient,stackname, constructID)
	if err != nil {
		return nil, err
	}
	parms := &awsec2.DescribeVpcsInput{
		VpcIds:     []string{*vpcId},
	}
	vpcInfo,err := client.DescribeVpcs(context.TODO(), parms)
	if err != nil {
		return nil, err
	}
	vpc := vpcInfo.Vpcs[0]
	return &vpc, nil
}

// GetSecurityGroup - VPC Object
func GetSecurityGroup(stackname *string, constructID *string) (*types.SecurityGroup, error ){
	securityGroupId,err := cit.PhysicalIDfromCID(cit.CfnClient, stackname, constructID)
	if err != nil {
		return nil, err
	}
	parms := &awsec2.DescribeSecurityGroupsInput{
		GroupIds:   []string{*securityGroupId},
	}
	sgInfo,err := client.DescribeSecurityGroups(context.TODO(), parms)
	if err != nil {
		return nil, err
	}
	sg := sgInfo.SecurityGroups[0]
	return &sg, nil
}

func GetInstance(stackname *string, constructID *string) (*types.Instance, error ){
	id, err:= cit.PhysicalIDfromCID(cit.CfnClient, stackname, constructID)
	if err != nil {
		return nil, err
	}
	parms := &awsec2.DescribeInstancesInput{
		InstanceIds: []string{*id},
	}
	info,err := client.DescribeInstances(context.TODO(), parms)
	if err != nil {
		return nil, err
	}
	object := info.Reservations[0].Instances[0]
	return &object, nil
}
