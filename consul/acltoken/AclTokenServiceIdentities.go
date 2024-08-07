// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acltoken


type AclTokenServiceIdentities struct {
	// The name of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/acl_token#service_name AclToken#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Specifies the datacenters the effective policy is valid within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/acl_token#datacenters AclToken#datacenters}
	Datacenters *[]*string `field:"optional" json:"datacenters" yaml:"datacenters"`
}

