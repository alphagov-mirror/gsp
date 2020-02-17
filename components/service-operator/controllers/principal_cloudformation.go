package controllers

import (
	access "github.com/alphagov/gsp/components/service-operator/apis/access/v1beta1"
	"github.com/alphagov/gsp/components/service-operator/internal/aws/cloudformation"
	"github.com/alphagov/gsp/components/service-operator/internal/aws/sdk"
	"github.com/alphagov/gsp/components/service-operator/internal/env"
	"github.com/aws/aws-sdk-go/aws"
	"strings"
)

func PrincipalCloudFormationController(c sdk.Client) Controller {
	return &cloudformation.Controller{
		Kind:              &access.Principal{},
		PrincipalListKind: &access.PrincipalList{},
		CloudFormationClient: &cloudformation.Client{
			Client: c,
		},
		Parameters: []*cloudformation.Parameter{
			{
				ParameterKey:   aws.String(access.IAMPermissionsBoundaryParameterName),
				ParameterValue: aws.String(env.AWSPrincipalPermissionsBoundaryARN()),
			},
			{
				ParameterKey:   aws.String(access.ServiceOperatorIAMRoleArn),
				ParameterValue: aws.String(env.AWSRoleArn()),
			},
			{
				ParameterKey:   aws.String(access.IAMOIDCProviderURLParameterName),
				ParameterValue: aws.String(strings.Replace(env.AWSOIDCProviderURL(), "https://", "", 1)),
			},
			{
				ParameterKey:   aws.String(access.IAMOIDCProviderARNParameterName),
				ParameterValue: aws.String(env.AWSOIDCProviderARN()),
			},
		},
		RequeueOnSuccess:           true,
		ReconcileSuccessRetryDelay: env.ImageRegistryCredentialsRenewalInterval(),
	}
}
