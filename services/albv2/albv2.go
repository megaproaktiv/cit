package albv2

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	awsalbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/megaproaktiv/cit"
)

var client *awsalbv2.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client  = awsalbv2.NewFromConfig(cfg)
}



func GetLoadBalancer(stackname *string, constructID *string) (*types.LoadBalancer, error ){
	id,err := cit.PhysicalIDfromCID(cit.CfnClient, stackname, constructID)
	if err != nil {
		log.Fatal("Error with PhysicalIDfromCID")
		return nil, err
	}
	parms := &awsalbv2.DescribeLoadBalancersInput{
		LoadBalancerArns: []string{*id},
	}
	objects,err := client.DescribeLoadBalancers(context.TODO(), parms)
	
	if err != nil {
		return nil, err
	}
	return &objects.LoadBalancers[0],nil
}
