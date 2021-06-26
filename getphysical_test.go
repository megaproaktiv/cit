package cit_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/megaproaktiv/cit"
	"gotest.tools/assert"
)

// Deploy testData stack before

func TestPhysicalID(t *testing.T) {

	// make and configure a mocked CloudFormationInterface
	mockedCloudFormationInterface := &cit.CloudFormationInterfaceMock{
		GetTemplateFunc: func(ctx context.Context, 
			params *cloudformation.GetTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error) {
			var templateOutput cloudformation.GetTemplateOutput
			data, err := ioutil.ReadFile("testdata/template-body.json")
			if err != nil {
				fmt.Println("File reading error: ", err)
			}
			content := string(data)
			templateOutput.TemplateBody = &content

			return &templateOutput, nil
		},
		DescribeStackResourceFunc: func(ctx context.Context, 
			params *cloudformation.DescribeStackResourceInput, 
			optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error) {
			var describeOutput cloudformation.DescribeStackResourceOutput
			data, err := ioutil.ReadFile("testdata/stack-resource.json")
			if err != nil {
				fmt.Println("File reading error: ", err)
			}
			json.Unmarshal(data, &describeOutput);
	
			return &describeOutput, nil
		},
	}

	actualPhid,err := cit.PhysicalIDfromCID(mockedCloudFormationInterface,
		aws.String("CdksnsStack"), 
		aws.String("MyTopic"))
	expectedPhid := "arn:aws:sns:eu-central-1:111122225555:CdksnsStack-MyTopic86869434-8W2KJU1PQTX0"
	assert.Nil(t, err, "PhysicalIDfromCID should give no error" )
	assert.Equal(t, expectedPhid, *actualPhid, "PhysicalID should match ConstructID")

}

func TestLogicalIDfromCID(t *testing.T) {
	testfile := "testdata/template-raw.json"
	data, err := ioutil.ReadFile(testfile)
	if err != nil {
		fmt.Println("File reading error: ", err)
	}
	content := string(data)
	logicalID,err := cit.LogicalIDfromCID(&content, aws.String("MyTopic"))
	assert.Nil(t, err, "LogicalIDfromCID should return no error")
	if err != nil{
		t.Fatal(err)
	}
	assert.Equal(t, "MyTopic86869434", *logicalID)

}

func TestMarshallTemplate(t *testing.T){
	testfile := "testdata/template-raw.json"

	stack := &cit.Template{}

	data, err := ioutil.ReadFile(testfile)
	
	if err != nil {
		panic("Cannot read "+testfile)
	}

	err = json.Unmarshal(data, stack); 
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err, "Unmarshal should work wo errors")

	logicalName := "MyTopic86869434"
	for key := range stack.Resources { 
		assert.Equal(t, logicalName, key)	
	}
	resource := stack.Resources[logicalName]
	assert.Equal(t, resource.Metadata["aws:cdk:path"], "CdksnsStack/MyTopic/Resource")
}

func TestExtractLogicalName(t *testing.T){
	path := "CdksnsStack/MyTopic/Resource"
	assert.Equal(t, "MyTopic", cit.ExtractConstructID(&path))
}