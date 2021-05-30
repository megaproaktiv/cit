package cit_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/megaproaktiv/cit"
	"github.com/stretchr/testify/assert"
)

// Deploy testData stack before
func TestPhysicalID(t *testing.T) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ssm.NewFromConfig(cfg)

	params := &ssm.GetParameterInput{
		Name: aws.String("/cit/test/snsid"),
	}

	res, err := client.GetParameter(context.TODO(), params)
	if err != nil {
		panic("Cant get SSM parameter - deploy testdata/cdksns stack ")
	}
	ssmid := res.Parameter.Value

	pID := cit.PhysicalIDfromCID(aws.String("CdksnsStack"), aws.String("MyTopic"))

	assert.Equal(t, *ssmid, *pID, "PhysicalID should match ConstructID")

}

func TestLogicalID(t *testing.T) {
	testfile := "testdata/template.json"

	logicalID := cit.LogicalIDfromCID(&testfile, aws.String("MyTopic"))

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
	assert.Equal(t, "MyTopic", cit.ExtractLogicalName(&path))
}