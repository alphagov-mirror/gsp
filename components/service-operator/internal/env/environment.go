// package env keeps calls for global flags or configuration in one place
package env

import (
	"fmt"
	"os"
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

// AWSPrincipalServerRoleARN is the arn of the role that can assume any provisioned roles
func AWSPrincipalServerRoleARN() string {
	return MustGet("AWS_PRINCIPAL_SERVER_ROLE_ARN")
}

// AWSPrincipalPermissionsBoundaryARN is the arn of the policy that limits permissions
func AWSPrincipalPermissionsBoundaryARN() string {
	return MustGet("AWS_PRINCIPAL_PERMISSIONS_BOUNDARY_ARN")
}

// MustGet is a panicy version of os.Getenv
func MustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Errorf("required environment variable '%s' not found", key))
	}
	return v
}
