// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclrole


type AclRoleTemplatedPolicies struct {
	// The name of the templated policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/acl_role#template_name AclRole#template_name}
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// Specifies the datacenters the effective policy is valid within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/acl_role#datacenters AclRole#datacenters}
	Datacenters *[]*string `field:"optional" json:"datacenters" yaml:"datacenters"`
	// template_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/acl_role#template_variables AclRole#template_variables}
	TemplateVariables *AclRoleTemplatedPoliciesTemplateVariables `field:"optional" json:"templateVariables" yaml:"templateVariables"`
}

