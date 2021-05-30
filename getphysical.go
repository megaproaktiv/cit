package cit

import (
	"fmt"
	"log"
	"strings"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	goformation "github.com/awslabs/goformation/v4"
	"github.com/awslabs/goformation/v4/intrinsics"
)

// PhysicalIDfromCID get AWS physical id from CDK contruct ID
// stack - Stackname
// constructID -  awssns.NewTopic(stack, jsii.String("MyTopic")  << MyTopic
func PhysicalIDfromCID(stack *string, constructId *string) *string {
	return aws.String("not implemented")
}

// LogicalIDfromCID - get logicalID
func LogicalIDfromCID(stackfile *string, constructID *string) *string{

	// Do not process, goformation can`t handle cdk generated data
	template, err := goformation.OpenWithOptions(*stackfile, &intrinsics.ProcessorOptions{
		NoProcess: true,
		EvaluateConditions: false,
	})
	if err != nil {
		log.Fatalf("There was an error processing the template: %s", err)
	}

	for _, resource := range template.Resources {
		fmt.Println(resource.AWSCloudFormationType())
	}

	

	return nil
}

func ExtractLogicalName( path *string) string{
	var parts = strings.Split(*path, "/")
	return parts[1]
}