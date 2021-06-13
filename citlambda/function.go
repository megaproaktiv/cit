package citlambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"github.com/megaproaktiv/cit"
)

var client LambdaInterface
//go:generate moq -out lambda_moq.go . LambdaInterface

type LambdaInterface interface {
	GetFunction(ctx context.Context, 
		params *lambda.GetFunctionInput, 
		optFns ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error)
	Invoke(ctx context.Context, 
		params *lambda.InvokeInput, 
		optFns ...func(*lambda.Options)) (*lambda.InvokeOutput, error)	
}

func init() {
	autoinit := cit.Autoinit()
	if autoinit {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("configuration error, " + err.Error())
		}
		client = lambda.NewFromConfig(cfg)
	}
}

func GetFunctionConfiguration(stackname *string, constructID *string)( *types.FunctionConfiguration, error){
	id,err := cit.PhysicalIDfromCID(cit.CfnClient,stackname, constructID)
	if err != nil {
		return nil, err
	}
	parms := &lambda.GetFunctionInput{
		FunctionName: id,
	}
	info, err := client.GetFunction(context.TODO(), parms)
	if err != nil {
		return nil, err
	}

	return info.Configuration, nil
}

func SetClient(c LambdaInterface) {
	client = c
}
