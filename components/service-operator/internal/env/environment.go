// package env keeps calls for global flags or configuration in one place
package env

import (
	"fmt"
	"os"
	"time"
)

// ClusterName returns the name of the cluster. panics if missing.
func ClusterName() string {
	return MustGet("CLUSTER_NAME")
}

// AWSIntegrationTestEnabled returns true if we should run the integration tests
func AWSIntegrationTestEnabled() bool {
	return os.Getenv("AWS_INTEGRATION") == "true"
}

// AWSRDSSecurityGroupID returns security group to assign to RDS provisioned resources
func AWSRDSSecurityGroupID() string {
	return MustGet("AWS_RDS_SECURITY_GROUP_ID")
}

// AWSRDSSubnetGroupName returns the subnet to provision RDS resources into
func AWSRDSSubnetGroupName() string {
	return MustGet("AWS_RDS_SUBNET_GROUP_NAME")
}

// AWSRedisSecurityGroupID returns security group to assign to Redis provisioned resources
func AWSRedisSecurityGroupID() string {
	return MustGet("AWS_REDIS_SECURITY_GROUP_ID")
}

// AWSRedisSubnetGroupName returns the subnet to provision Redis resources into
func AWSRedisSubnetGroupName() string {
	return MustGet("AWS_REDIS_SUBNET_GROUP_NAME")
}

// AWSOIDCProviderURL is the URL of the OIDC provider for our EKS cluster
func AWSOIDCProviderURL() string {
	return MustGet("AWS_OIDC_PROVIDER_URL")
}

// AWSOIDCProviderARN is the URL of the OIDC provider for our EKS cluster
func AWSOIDCProviderARN() string {
	return MustGet("AWS_OIDC_PROVIDER_ARN")
}

// AWSPrincipalPermissionsBoundaryARN is the arn of the policy that limits permissions
func AWSPrincipalPermissionsBoundaryARN() string {
	return MustGet("AWS_PRINCIPAL_PERMISSIONS_BOUNDARY_ARN")
}

func AWSRoleArn() string {
	return MustGet("AWS_ROLE_ARN")
}

func ImageRegistryCredentialsRenewalInterval() time.Duration {
	envVarName := "IMAGE_REGISTRY_CREDENTIALS_RENEWAL_INTERVAL"
	envVarValue := os.Getenv(envVarName)
	renewalInterval := time.Minute * 30

	if envVarValue != "" {
		var err error
		renewalInterval, err = time.ParseDuration(envVarValue)
		if err != nil {
			panic(fmt.Errorf("failed to parse duration from %s (%s)", envVarName, envVarValue))
		}
	}

	return renewalInterval
}

// MustGet is a panicy version of os.Getenv
func MustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Errorf("required environment variable '%s' not found", key))
	}
	return v
}
