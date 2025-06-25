// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclrole


type AclRoleServiceIdentities struct {
	// The name of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/acl_role#service_name AclRole#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The datacenters the effective policy is valid within.
	//
	// When no datacenters are provided the effective policy is valid in all datacenters including those which do not yet exist but may in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/acl_role#datacenters AclRole#datacenters}
	Datacenters *[]*string `field:"optional" json:"datacenters" yaml:"datacenters"`
}

