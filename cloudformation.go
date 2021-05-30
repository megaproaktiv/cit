package cit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

//go:generate moq -out cloudformation_moq_test.go . CloudFormationInterface

type CloudFormationInterface interface {
	GetTemplate(ctx context.Context, params *cloudformation.GetTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error)
	DescribeStackResource(ctx context.Context, params *cloudformation.DescribeStackResourceInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error)
}