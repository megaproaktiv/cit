package citsts

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	
)

var client *sts.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client  = sts.NewFromConfig(cfg)
}

// GetAccount - Number
func GetAccount() (*string, error ){
	params := &sts.GetCallerIdentityInput{

	}
	iamwhatiam, err := client.GetCallerIdentity(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	account := iamwhatiam.Account
	return account, err
}
