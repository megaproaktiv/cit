package citlambda

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/stretchr/testify/assert"

	"github.com/megaproaktiv/cit"
)

func TestGetFunctionConfiguration(t *testing.T) {
	
	mockedLambdaInterface := &LambdaInterfaceMock{
		GetFunctionFunc: func(ctx context.Context, params *lambda.GetFunctionInput, optFns ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error) {
			var output lambda.GetFunctionOutput
			data, err := ioutil.ReadFile("testdata/get-function-positive.json")
			if err != nil {
				t.Error("Cant read input testdata")
				t.Error(err)
			}
			json.Unmarshal(data, &output);
			return &output,nil
		},
	}
	
	mockedCloudFormationInterface := &cit.CloudFormationInterfaceMock{
		DescribeStackResourceFunc: func(ctx context.Context, 
			params *cloudformation.DescribeStackResourceInput, 
			optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error) {
				var output cloudformation.DescribeStackResourceOutput
				data, err := ioutil.ReadFile("testdata/describe-stack-resource.json")
				if err != nil {
					t.Error("Cant read input testdata")
					t.Error(err)
				}
				json.Unmarshal(data, &output);
				return &output,nil
		},
		GetTemplateFunc: func(ctx context.Context, 
			params *cloudformation.GetTemplateInput, 
			optFns ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error) {
				panic("mock out the GetTemplate method")
		},
	}

	os.Setenv("AUTO_INIT", "false")
	SetClient(mockedLambdaInterface)
	cit.SetClient(mockedCloudFormationInterface)
	
	got, err := GetFunctionConfiguration(aws.String("LambdaSimpleStack"), aws.String("HelloHandler"))
	assert.Nil(t, err, "GetFunction should return no error")
	expect := &types.FunctionConfiguration{
		FunctionName:               aws.String("LambdaSimpleStack-HelloHandler2E4FBA4D-ZqznH9gomexC"),
		Handler:                    aws.String("hello.handler"),
		MemorySize:                 aws.Int32(1024),
	}
	assert.Equal(t, expect.FunctionName, got.FunctionName, "Function name should match")
	assert.Equal(t, expect.Handler, got.Handler, "Handler should match")
	assert.Equal(t, expect.MemorySize, got.MemorySize, "Memory should match")
	
}
