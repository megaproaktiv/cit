package cit

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"strings"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

// PhysicalIDfromCID get AWS physical id from CDK contruct ID
// stack - Stackname
// constructID -  awssns.NewTopic(stack, jsii.String("MyTopic")  << MyTopic
func PhysicalIDfromCID(client CloudFormationInterface, stack *string, constructId *string) (*string, error) {
	// Get Stack
	parameterStack := &cloudformation.GetTemplateInput{
		StackName:     stack,
	}

	resGetTemplate, err := client.GetTemplate(context.TODO(), parameterStack)
	if err != nil {
		log.Fatal("Template "+*stack+" not found")
		return nil, err
	}

	template := resGetTemplate.TemplateBody
	// Find LogicalID
	logicalID,err := LogicalIDfromCID(template, constructId)
	if err != nil {
		return nil, err
	}
	// get stackresources
	parameterResource := &cloudformation.DescribeStackResourceInput{
		LogicalResourceId: logicalID,
		StackName:         stack,
	}
	resourceDetail,err := client.DescribeStackResource(context.TODO(), parameterResource)
	if err != nil {
		return nil, err
	}
	// find physicalid
	physicalId := resourceDetail.StackResourceDetail.PhysicalResourceId
	
	return physicalId,nil
}

// LogicalIDfromCID - get logicalID
func LogicalIDfromCID(stackContent *string, constructID *string) (*string, error) {

	stack := &Template{}
	err := json.Unmarshal([]byte(*stackContent), stack)
	
	if err != nil {
		log.Fatal("There was an error processing the template: ", err)
		return nil, err
	}

	for key, resource := range stack.Resources {
		if resource.Metadata != nil {
			if resource.Metadata["aws:cdk:path"] != "" {
				meta := resource.Metadata["aws:cdk:path"]
				log.Debug("Path: ",meta)
				templateConstructID := ExtractConstructID(&meta)
				if templateConstructID == *constructID {
						return &key, nil
				}
			}
		}
	}

	return aws.String(""), errors.New("ConstructID not found")
}

// ExtractCounstructid
//         "aws:cdk:path": "vpc/baseVPC/Resource"
func ExtractConstructID(path *string) string {
	var parts = strings.Split(*path, "/")
	if parts[2] == "Resource"{
		return parts[1]
	}else{
		return ""
	}
}
