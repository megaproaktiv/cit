package cit_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/megaproaktiv/cit"
	"github.com/stretchr/testify/assert"
)

// Deploy testData stack before

func TestPhysicalIDBoolError(t *testing.T) {

	data, err := ioutil.ReadFile("testdata/dsl.template.json")
	if err != nil {
		fmt.Println("File reading error: ", err)
	}
	content := string(data)

	actualPhid, err := cit.LogicalIDfromCID(&content,aws.String("myHandler") )
	
	expectedPhid := "myHandler0D56A5FA"
	assert.Nil(t, err, "PhysicalIDfromCID should give no error" )
	assert.Equal(t, expectedPhid, *actualPhid, "PhysicalID should match ConstructID")

}
