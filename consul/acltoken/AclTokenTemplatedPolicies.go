// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acltoken


type AclTokenTemplatedPolicies struct {
	// The name of the templated policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/acl_token#template_name AclToken#template_name}
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// Specifies the datacenters the effective policy is valid within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/acl_token#datacenters AclToken#datacenters}
	Datacenters *[]*string `field:"optional" json:"datacenters" yaml:"datacenters"`
	// template_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/acl_token#template_variables AclToken#template_variables}
	TemplateVariables *AclTokenTemplatedPoliciesTemplateVariables `field:"optional" json:"templateVariables" yaml:"templateVariables"`
}

