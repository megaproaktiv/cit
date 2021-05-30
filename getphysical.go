package cit

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	aws "github.com/aws/aws-sdk-go-v2/aws"
)

// PhysicalIDfromCID get AWS physical id from CDK contruct ID
// stack - Stackname
// constructID -  awssns.NewTopic(stack, jsii.String("MyTopic")  << MyTopic
func PhysicalIDfromCID(stack *string, constructId *string) *string {
	return aws.String("not implemented")
}

// LogicalIDfromCID - get logicalID
func LogicalIDfromCID(stackfile *string, constructID *string) (*string, error) {

	stack := &Template{}
	data, err := ioutil.ReadFile(*stackfile)
	if err != nil {
		panic("Cannot read " + *stackfile)
	}

	err = json.Unmarshal(data, stack)
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
