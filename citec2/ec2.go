// CDK Infrastructure Testing (cit) - EC2 helper
// 
// You run the cit test after 
// 		cdk deploy
//
// Example:
// 		func TestVpcCit(t *testing.T){
// 
// 			vpc,err := ec2.GetVpc(aws.String("VpcStack"), aws.String("MyVpc"))
// 			assert.Nil(t,err)
// 			assert.Equal(t, "10.0.0.0/16",*vpc.CidrBlock)
// 		}

package citec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/megaproaktiv/cit"
)

var client *ec2.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client  = ec2.NewFromConfig(cfg)
}

// GetVPC - VPC Object
// stackname - the CDK stack name as given in 
// 		cdk ls
// constructId - the ID of the vpc Construct
func GetVpc(stackname *string, constructID *string) (*types.Vpc, error ){
	vpcId,err := cit.PhysicalIDfromCID(cit.CfnClient,stackname, constructID)
	if err != nil {
		return nil, err
	}
	parms := &ec2.DescribeVpcsInput{
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
// stackname - the CDK stack name as given in 
// 		cdk ls
// constructId - the ID of the security group Construct
func GetSecurityGroup(stackname *string, constructID *string) (*types.SecurityGroup, error ){
	securityGroupId,err := cit.PhysicalIDfromCID(cit.CfnClient, stackname, constructID)
	if err != nil {
		return nil, err
	}
	parms := &ec2.DescribeSecurityGroupsInput{
		GroupIds:   []string{*securityGroupId},
	}
	sgInfo,err := client.DescribeSecurityGroups(context.TODO(), parms)
	if err != nil {
		return nil, err
	}
	sg := sgInfo.SecurityGroups[0]
	return &sg, nil
}

// GetSecurityGroup - VPC Object
// stackname - the CDK stack name as given in 
// 		cdk ls
// constructId - the ID of the instance Construct
func GetInstance(stackname *string, constructID *string) (*types.Instance, error ){
	id, err:= cit.PhysicalIDfromCID(cit.CfnClient, stackname, constructID)
	if err != nil {
		return nil, err
	}
	parms := &ec2.DescribeInstancesInput{
		InstanceIds: []string{*id},
	}
	info,err := client.DescribeInstances(context.TODO(), parms)
	if err != nil {
		return nil, err
	}
	object := info.Reservations[0].Instances[0]
	return &object, nil
}
