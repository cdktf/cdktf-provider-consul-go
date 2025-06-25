// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acltoken


type AclTokenTemplatedPoliciesTemplateVariables struct {
	// The name of node, workload identity or service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/acl_token#name AclToken#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

