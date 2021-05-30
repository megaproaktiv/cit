package cit

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

// PhysicalIDfromCID get AWS physical id from CDK contruct ID
// stack - Stackname
// constructID -  awssns.NewTopic(stack, jsii.String("MyTopic")  << MyTopic
func PhysicalIDfromCID(client CloudFormationInterface, stack *string, constructId *string) *string {
	// Get Stack
	parameterStack := &cloudformation.GetTemplateInput{
		StackName:     stack,
	}
	resGetTemplate, err := client.GetTemplate(context.TODO(), parameterStack)
	if err != nil {
		panic(err)
	}

	template := resGetTemplate.TemplateBody
	// Find LogicalID
	logicalID,err := LogicalIDfromCID(template, constructId)
	if err != nil {
		panic(err)
	}
	// get stackresources
	parameterResource := &cloudformation.DescribeStackResourceInput{
		LogicalResourceId: logicalID,
		StackName:         stack,
	}
	resourceDetail,err := client.DescribeStackResource(context.TODO(), parameterResource)
	if err != nil {
		panic(err)
	}
	// find physicalid
	physicalId := resourceDetail.StackResourceDetail.PhysicalResourceId
	
	return physicalId
}

// LogicalIDfromCID - get logicalID
func LogicalIDfromCID(stackContent *string, constructID *string) (*string, error) {

	stack := &Template{}
	err := json.Unmarshal([]byte(*stackContent), stack)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal("There was an error processing the template: ", err)
	}

	for key, resource := range stack.Resources {
		if resource.Metadata != nil {
			if resource.Metadata["aws:cdk:path"] != "" {
				meta := resource.Metadata["aws:cdk:path"]
				templateConstructID := ExtractConstructID(&meta)
				if templateConstructID == *constructID {
						return &key, nil
				}
			}
		}
	}

	return aws.String(""), errors.New("ConstructID not found")
}

func ExtractConstructID(path *string) string {
	var parts = strings.Split(*path, "/")
	return parts[1]
}
