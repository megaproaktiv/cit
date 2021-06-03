package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	awsiam "github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/megaproaktiv/cit"
)

var client *awsiam.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client  = awsiam.NewFromConfig(cfg)
}

func GetUser(stackname *string, constructID *string)(*types.User, error){
	id, err:= cit.PhysicalIDfromCID(cit.CfnClient, stackname, constructID)
	if err != nil {
		return nil, err
	}
	parms := &awsiam.GetUserInput{
		UserName: id,
	}
	info,err := client.GetUser(context.TODO(), parms)
	if err != nil {
		return nil, err
	}
	object := info.User
	return object, nil
}