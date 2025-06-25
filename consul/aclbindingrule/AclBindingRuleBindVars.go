// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclbindingrule


type AclBindingRuleBindVars struct {
	// The name of node, workload identity or service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/acl_binding_rule#name AclBindingRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

