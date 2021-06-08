package citses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/megaproaktiv/cit"
)

var client SesInterface

//go:generate moq -out ses_moq_test.go . SesInterface

type SesInterface interface {
	ListIdentities(ctx context.Context,
		params *ses.ListIdentitiesInput,
		optFns ...func(*ses.Options)) (*ses.ListIdentitiesOutput, error)
}

func init() {
	autoinit := cit.Autoinit()
	if autoinit {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("configuration error, " + err.Error())
		}
		client = ses.NewFromConfig(cfg)
	}
}

func IdentityExists(indentity *string) (bool, error) {
	parms := &ses.ListIdentitiesInput{}
	info, err := client.ListIdentities(context.TODO(), parms)
	if err != nil {
		return false, err
	}
	for _, r := range info.Identities {
		if r == *indentity {
			return true, nil
		}
	}
	return false, nil
}

func SetClient(c SesInterface) {
	client = c
}
