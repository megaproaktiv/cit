package citiam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/megaproaktiv/cit"

)

var client *iam.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client  = iam.NewFromConfig(cfg)
}

func GetUser(stackname *string, constructID *string)(*types.User, error){
	id, err:= cit.PhysicalIDfromCID(cit.CfnClient, stackname, constructID)
	if err != nil {
		return nil, err
	}
	parms := &iam.GetUserInput{
		UserName: id,
	}
	info,err := client.GetUser(context.TODO(), parms)
	if err != nil {
		return nil, err
	}
	object := info.User
	return object, nil
}