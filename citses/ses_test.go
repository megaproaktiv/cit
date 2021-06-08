package citses

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/stretchr/testify/assert"
	
)

func TestSesInterfaceTrue(t *testing.T) {

	// make and configure a mocked SesInterface
	mockedSesInterfaceTrue := &SesInterfaceMock{
		ListIdentitiesFunc: func(ctx context.Context, params *ses.ListIdentitiesInput, optFns ...func(*ses.Options)) (*ses.ListIdentitiesOutput, error) {
				
				var output ses.ListIdentitiesOutput
				data, err := ioutil.ReadFile("testdata/list-identities-true.json")
				if err != nil {
					t.Error("Cant read input testdata")
					t.Error(err)
				}
				json.Unmarshal(data, &output);
				return &output,nil
		},
	}

	os.Setenv("AUTO_INIT", "false")
	SetClient(mockedSesInterfaceTrue)

	identity := "info@megaproaktiv.de"
	actualValue,_ := IdentityExists(&identity)
	assert.Equal(t, true, actualValue)

	

}

func TestSesInterfaceFalse(t *testing.T) {

	// make and configure a mocked SesInterface
	mockedSesInterfaceTrue := &SesInterfaceMock{
		ListIdentitiesFunc: func(ctx context.Context, params *ses.ListIdentitiesInput, optFns ...func(*ses.Options)) (*ses.ListIdentitiesOutput, error) {
				
				var output ses.ListIdentitiesOutput
				data, err := ioutil.ReadFile("testdata/list-identities-false.json")
				if err != nil {
					t.Error("Cant read input testdata")
					t.Error(err)
				}
				json.Unmarshal(data, &output);
				return &output,nil
		},
	}

	os.Setenv("AUTO_INIT", "false")
	SetClient(mockedSesInterfaceTrue)

	identity := "info@megaproaktiv.de"
	actualValue,_ := IdentityExists(&identity)
	assert.Equal(t, false, actualValue)

	

}