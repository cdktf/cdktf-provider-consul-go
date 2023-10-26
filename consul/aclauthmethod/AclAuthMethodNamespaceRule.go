// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclauthmethod


type AclAuthMethodNamespaceRule struct {
	// If the namespace rule's `selector` matches then this is used to control the namespace where the token is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/acl_auth_method#bind_namespace AclAuthMethod#bind_namespace}
	BindNamespace *string `field:"required" json:"bindNamespace" yaml:"bindNamespace"`
	// Specifies the expression used to match this namespace rule against valid identities returned from an auth method validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/acl_auth_method#selector AclAuthMethod#selector}
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
}

