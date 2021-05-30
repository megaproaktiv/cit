package cit_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/megaproaktiv/cit"
	"github.com/stretchr/testify/assert"
)

// Deploy testData stack before
func TestIntegPhysicalID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping aws api access")
	  }
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	clientSsm := ssm.NewFromConfig(cfg)
	clientCfn := cloudformation.NewFromConfig(cfg)

	params := &ssm.GetParameterInput{
		Name: aws.String("/cit/test/snsid"),
	}

	res, err := clientSsm.GetParameter(context.TODO(), params)
	if err != nil {
		panic("Cant get SSM parameter - deploy testdata/cdksns stack ")
	}
	ssmid := res.Parameter.Value

	pID := cit.PhysicalIDfromCID(clientCfn, aws.String("CdksnsStack"), aws.String("MyTopic"))

	assert.Equal(t, *ssmid, *pID, "PhysicalID should match ConstructID")

}
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

	actualPhid := cit.PhysicalIDfromCID(mockedCloudFormationInterface,
		aws.String("CdksnsStack"), 
		aws.String("MyTopic"))
	expectedPhid := "arn:aws:sns:eu-central-1:703466486373:CdksnsStack-MyTopic86869434-8W2KJU1PQTX0"

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