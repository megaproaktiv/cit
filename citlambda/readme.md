## Example

### Testing if a Lambda Function exists.

In GO you have an  CDK Stack with a simple Lambda Funktion:

```go
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lambdaPath := filepath.Join(path, "lambda")
	awslambda.NewFunction(stack, 
		aws.String("HelloHandler"),
		&awslambda.FunctionProps{
			MemorySize:                   aws.Float64(1024),
			Code:                         awslambda.Code_FromAsset(&lambdaPath, 
				&awss3assets.AssetOptions{}),
			Handler:                      aws.String("hello.handler"),
			Runtime:                      awslambda.Runtime_NODEJS_14_X(),
		})
```

To test the created physical Function you may write a test:

```go
func TestCitLambda(t *testing.T){

	gotFc, err := citlambda.GetFunctionConfiguration(aws.String("LambdaSimpleStack"),
	aws.String("HelloHandler"))
	assert.Nil(t, err, "GetFunctionConfiguration should return no error")
	expectHandler := "hello.handler"
	assert.Equal(t, &expectHandler, gotFc.Handler )
}
```

### Stack is not deployed

```bash
 go test -v
```

gives 

```log
=== RUN   TestCitLambda
FATA[0007] Template LambdaSimpleStack not found
exit status 1
FAIL	lambda-simple	
```

### Stack is deployed

```bash
 go test -v
```

gives

```log
=== RUN   TestCitLambda
--- PASS: TestCitLambda (2.75s)
PASS
ok  	lambda-simple	
```


        
