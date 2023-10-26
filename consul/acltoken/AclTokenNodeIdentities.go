// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acltoken


type AclTokenNodeIdentities struct {
	// The datacenter of the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/acl_token#datacenter AclToken#datacenter}
	Datacenter *string `field:"required" json:"datacenter" yaml:"datacenter"`
	// The name of the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/acl_token#node_name AclToken#node_name}
	NodeName *string `field:"required" json:"nodeName" yaml:"nodeName"`
}

