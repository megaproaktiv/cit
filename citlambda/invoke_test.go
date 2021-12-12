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
	"github.com/megaproaktiv/cit"
	"github.com/megaproaktiv/cit/citlambda"
	"gotest.tools/assert"
)

func TestInvokeFunction(t *testing.T) {

	mockedLambdaInterface := &citlambda.LambdaInterfaceMock{
		GetFunctionFunc: func(ctx context.Context, params *lambda.GetFunctionInput, optFns ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error) {
			var output lambda.GetFunctionOutput
			data, err := ioutil.ReadFile("testdata/invoke-test-get-function.json")
			if err != nil {
				t.Error("Cant read input testdata")
				t.Error(err)
			}
			json.Unmarshal(data, &output);
			return &output,nil
		},
		InvokeFunc: func(ctx context.Context, params *lambda.InvokeInput, optFns ...func(*lambda.Options)) (*lambda.InvokeOutput, error) {
			var output lambda.InvokeOutput
			output.Payload = []byte("done")
			return &output,nil
			
		},
	}

	mockedCloudFormationInterface := &cit.CloudFormationInterfaceMock{
		GetTemplateFunc: func(ctx context.Context, 
			params *cloudformation.GetTemplateInput, 
			optFns ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error) {

				var templateOutput cloudformation.GetTemplateOutput
				
				data, err := ioutil.ReadFile("testdata/invoke-template.json")
				if err != nil {
					fmt.Println("File reading error: ", err)
				}
				content := string(data)
				templateOutput.TemplateBody = &content
				//t.Log("Resources: ",*templateOutput.TemplateBody)
				return &templateOutput, nil
		},
		DescribeStackResourceFunc: func(ctx context.Context, 
			params *cloudformation.DescribeStackResourceInput, 
			optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error) {
				var output cloudformation.DescribeStackResourceOutput
				data, err := ioutil.ReadFile("testdata/invoke-test-logical-resource-hello-handler.json")
				if err != nil {
					t.Error("Cant read input testdata")
					t.Error(err)
				}
				json.Unmarshal(data, &output);
				t.Log("DescribeStackResourceFunc: ",output)
				return &output,nil
		},
	}
	

	mockedLambdaInterface := &citlambda.LambdaInterfaceMock{
		GetFunctionFunc: func(ctx context.Context, params *lambda.GetFunctionInput, optFns ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error) {
			var output lambda.GetFunctionOutput
				data, err := ioutil.ReadFile("testdata/invoke-test-get-function.json")
				if err != nil {
					t.Error("Cant read input testdata")
					t.Error(err)
				}
				json.Unmarshal(data, &output);
				return &output,nil
		},
		InvokeFunc: func(ctx context.Context, params *lambda.InvokeInput, optFns ...func(*lambda.Options)) (*lambda.InvokeOutput, error) {
			var output lambda.InvokeOutput
			output.Payload = []byte("Done")
			return &output,nil

	os.Setenv("AUTO_INIT", "false")
	citlambda.SetClient(mockedLambdaInterface)
	eventFileName := "./testdata/invoke-test-event.json"
	got, err := citlambda.InvokeFunction( aws.String("LambdaGoStack"), aws.String("HelloHandler"), &eventFileName)
	assert.NilError(t, err)
	want := "Done"
	assert.Equal(t, want, *got)
			
}
