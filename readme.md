# CIT

CDK Infrastructure Testing

Project to support Infrastructure Testing with AWS CDK.
You get physical Resource objects from the name of the cdk stack and the constructid.

![Overview](img/mapping.png)

See [Blog1](https://aws-blog.de/2021/05/cit-build-cdk-infrastructure-testing-part-1-terratest-and-the-integrated-integration.html).


## Example

### Testing the resulting cidr of an vpc

In Typescript you create a stack named "vpc":

```ts
new VpcProjektStack(app, 'vpc', {
```

And name the construct "baseVPC".

```ts
var vpc = new Vpc(this, "baseVPC", {
```

Now you can test things on the created AWS vpc with a small go programm:

```go
package main

import (
	"testing"
	aws "github.com/aws/aws-sdk-go-v2/aws"
	ec2 "github.com/megaproaktiv/cit/services/ec2"
	"github.com/stretchr/testify/assert"
)

func TestVpcPro(t *testing.T) {
	
	vpc,err := ec2.GetVpc(aws.String("vpc"), aws.String("baseVPC"))
	assert.Nil(t,err)
	assert.Equal(t, "10.0.96.0/21",*vpc.CidrBlock)
}
```

Which you call with 

```bash
go test
```

Or more verbose


```bash
go test -v
```

"PASS" or "FAIL" will indicate whether the created vpc exists and have this cdr block.

