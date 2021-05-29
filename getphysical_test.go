package cit_test

import (
	"github.com/megaproaktiv/cit"
	"context"
	"testing"
	"github.com/stretchr/testify/assert"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
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
	ssmid := res.Parameter.Value

	pID := cit.PhysicalID(aws.String("CdksnsStack"),aws.String("MyTopic"))

	assert.Equal(t, *ssmid, *pID, "PhysicalID should match ConstructID")
	
}
