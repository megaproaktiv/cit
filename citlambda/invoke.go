package citlambda

import (
	// core
	"context"
	"io/ioutil"

	// sdk
	"github.com/aws/aws-sdk-go-v2/service/lambda"

)

// InvokeFunction
// stackName - name of an CDK deployed Cloudformation Stack
// constructID - name of the CDK Construct of the Lambda Resource
// eventFile - local json file, which will be synchronosly send to the Lambda function
func InvokeFunction(stackname *string, constructID *string, eventFile *string)( *string, error){
	functionConfiguration, err := GetFunctionConfiguration(
		stackname,
		constructID,
	)
	if err != nil {
		return nil, err
	}
	functionName := functionConfiguration.FunctionName

	data, err := ioutil.ReadFile(*eventFile)
	if err != nil {
		return nil, err
	}
	params := &lambda.InvokeInput{
		FunctionName:   functionName,
		Payload:        data,
	}
	res, err := client.Invoke(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	responseString := string(res.Payload)
	return &responseString, nil
}