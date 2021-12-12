// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cit

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"sync"
)

// Ensure, that CloudFormationInterfaceMock does implement CloudFormationInterface.
// If this is not the case, regenerate this file with moq.
var _ CloudFormationInterface = &CloudFormationInterfaceMock{}

// CloudFormationInterfaceMock is a mock implementation of CloudFormationInterface.
//
//     func TestSomethingThatUsesCloudFormationInterface(t *testing.T) {
//
//         // make and configure a mocked CloudFormationInterface
//         mockedCloudFormationInterface := &CloudFormationInterfaceMock{
//             DescribeStackResourceFunc: func(ctx context.Context, params *cloudformation.DescribeStackResourceInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error) {
// 	               panic("mock out the DescribeStackResource method")
//             },
//             GetTemplateFunc: func(ctx context.Context, params *cloudformation.GetTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error) {
// 	               panic("mock out the GetTemplate method")
//             },
//         }
//
//         // use mockedCloudFormationInterface in code that requires CloudFormationInterface
//         // and then make assertions.
//
//     }
type CloudFormationInterfaceMock struct {
	// DescribeStackResourceFunc mocks the DescribeStackResource method.
	DescribeStackResourceFunc func(ctx context.Context, params *cloudformation.DescribeStackResourceInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error)

	// GetTemplateFunc mocks the GetTemplate method.
	GetTemplateFunc func(ctx context.Context, params *cloudformation.GetTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error)

	// calls tracks calls to the methods.
	calls struct {
		// DescribeStackResource holds details about calls to the DescribeStackResource method.
		DescribeStackResource []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *cloudformation.DescribeStackResourceInput
			// OptFns is the optFns argument value.
			OptFns []func(*cloudformation.Options)
		}
		// GetTemplate holds details about calls to the GetTemplate method.
		GetTemplate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *cloudformation.GetTemplateInput
			// OptFns is the optFns argument value.
			OptFns []func(*cloudformation.Options)
		}
	}
	lockDescribeStackResource sync.RWMutex
	lockGetTemplate           sync.RWMutex
}

// DescribeStackResource calls DescribeStackResourceFunc.
func (mock *CloudFormationInterfaceMock) DescribeStackResource(ctx context.Context, params *cloudformation.DescribeStackResourceInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error) {
	if mock.DescribeStackResourceFunc == nil {
		panic("CloudFormationInterfaceMock.DescribeStackResourceFunc: method is nil but CloudFormationInterface.DescribeStackResource was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *cloudformation.DescribeStackResourceInput
		OptFns []func(*cloudformation.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockDescribeStackResource.Lock()
	mock.calls.DescribeStackResource = append(mock.calls.DescribeStackResource, callInfo)
	mock.lockDescribeStackResource.Unlock()
	return mock.DescribeStackResourceFunc(ctx, params, optFns...)
}

// DescribeStackResourceCalls gets all the calls that were made to DescribeStackResource.
// Check the length with:
//     len(mockedCloudFormationInterface.DescribeStackResourceCalls())
func (mock *CloudFormationInterfaceMock) DescribeStackResourceCalls() []struct {
	Ctx    context.Context
	Params *cloudformation.DescribeStackResourceInput
	OptFns []func(*cloudformation.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *cloudformation.DescribeStackResourceInput
		OptFns []func(*cloudformation.Options)
	}
	mock.lockDescribeStackResource.RLock()
	calls = mock.calls.DescribeStackResource
	mock.lockDescribeStackResource.RUnlock()
	return calls
}

// GetTemplate calls GetTemplateFunc.
func (mock *CloudFormationInterfaceMock) GetTemplate(ctx context.Context, params *cloudformation.GetTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error) {
	if mock.GetTemplateFunc == nil {
		panic("CloudFormationInterfaceMock.GetTemplateFunc: method is nil but CloudFormationInterface.GetTemplate was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *cloudformation.GetTemplateInput
		OptFns []func(*cloudformation.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockGetTemplate.Lock()
	mock.calls.GetTemplate = append(mock.calls.GetTemplate, callInfo)
	mock.lockGetTemplate.Unlock()
	return mock.GetTemplateFunc(ctx, params, optFns...)
}

// GetTemplateCalls gets all the calls that were made to GetTemplate.
// Check the length with:
//     len(mockedCloudFormationInterface.GetTemplateCalls())
func (mock *CloudFormationInterfaceMock) GetTemplateCalls() []struct {
	Ctx    context.Context
	Params *cloudformation.GetTemplateInput
	OptFns []func(*cloudformation.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *cloudformation.GetTemplateInput
		OptFns []func(*cloudformation.Options)
	}
	mock.lockGetTemplate.RLock()
	calls = mock.calls.GetTemplate
	mock.lockGetTemplate.RUnlock()
	return calls
}