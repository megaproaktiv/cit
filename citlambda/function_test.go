package citlambda_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/megaproaktiv/cit"
	"github.com/megaproaktiv/cit/citlambda"
	"github.com/stretchr/testify/assert"
)

func TestGetFunctionConfiguration(t *testing.T) {
	
	mockedLambdaInterface := &citlambda.LambdaInterfaceMock{
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
				data, err := ioutil.ReadFile("testdata/function-test-describe-stack-resource.json")
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
				var templateOutput cloudformation.GetTemplateOutput
				
				data, err := ioutil.ReadFile("testdata/function-test-template.json")
				if err != nil {
					fmt.Println("File reading error: ", err)
				}
				content := string(data)
				templateOutput.TemplateBody = &content
	
				return &templateOutput, nil
			},
	}

	os.Setenv("AUTO_INIT", "false")
	citlambda.SetClient(mockedLambdaInterface)
	cit.SetClient(mockedCloudFormationInterface)
	
	got, err := citlambda.GetFunctionConfiguration(aws.String("LambdaSimpleStack"), aws.String("HelloHandler"))
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
