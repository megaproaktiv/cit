package citlambda

import "testing"

func TestInvokeFunction(t *testing.T) {
	
	mockedLambdaInterface := &LambdaInterfaceMock{
		GetFunctionFunc: func(ctx context.Context, params *lambda.GetFunctionInput, optFns ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error) {
			panic("mock out the GetFunction method")
		},
		InvokeFunc: func(ctx context.Context, params *lambda.InvokeInput, optFns ...func(*lambda.Options)) (*lambda.InvokeOutput, error) {
			panic("mock out the Invoke method")
		},
	}

	//got, err := InvokeFunction(tt.args.stackname, tt.args.constructID, tt.args.eventFile)
	//if got != want {
			
}
