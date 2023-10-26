// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclrole


type AclRoleTemplatedPoliciesTemplateVariables struct {
	// The name of node, workload identity or service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/acl_role#name AclRole#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

