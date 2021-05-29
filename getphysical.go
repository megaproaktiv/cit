package cit

import (
	aws "github.com/aws/aws-sdk-go-v2/aws"

)

// PhysicalID get AWS physical id from CDK contruct ID
// stack - Stackname
// constructID -  awssns.NewTopic(stack, jsii.String("MyTopic")  << MyTopic
func PhysicalID(stack *string, constructId *string) *string {
	return aws.String("not implemented")
}
