// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type ConsulProviderAuthJwt struct {
	// The name of the auth method to use for login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#auth_method ConsulProvider#auth_method}
	AuthMethod *string `field:"required" json:"authMethod" yaml:"authMethod"`
	// The bearer token to present to the auth method during login for authentication purposes.
	//
	// For the Kubernetes auth method this is a [Service Account Token (JWT)](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#service-account-tokens).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#bearer_token ConsulProvider#bearer_token}
	BearerToken *string `field:"optional" json:"bearerToken" yaml:"bearerToken"`
	// Specifies arbitrary KV metadata linked to the token. Can be useful to track origins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#meta ConsulProvider#meta}
	Meta *map[string]*string `field:"optional" json:"meta" yaml:"meta"`
	// Whether to use a [Terraform Workload Identity token](https://developer.hashicorp.com/terraform/cloud-docs/workspaces/dynamic-provider-credentials/workload-identity-tokens). The token will be read from the `TFC_WORKLOAD_IDENTITY_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#use_terraform_cloud_workload_identity ConsulProvider#use_terraform_cloud_workload_identity}
	UseTerraformCloudWorkloadIdentity interface{} `field:"optional" json:"useTerraformCloudWorkloadIdentity" yaml:"useTerraformCloudWorkloadIdentity"`
}

