// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acltoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclTokenConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The uuid of the token. If omitted, Consul will generate a random uuid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#accessor_id AclToken#accessor_id}
	AccessorId *string `field:"optional" json:"accessorId" yaml:"accessorId"`
	// The description of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#description AclToken#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set this represents the point after which a token should be considered revoked and is eligible for destruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#expiration_time AclToken#expiration_time}
	ExpirationTime *string `field:"optional" json:"expirationTime" yaml:"expirationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#id AclToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The flag to set the token local to the current datacenter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#local AclToken#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// The namespace to create the token within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#namespace AclToken#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// node_identities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#node_identities AclToken#node_identities}
	NodeIdentities interface{} `field:"optional" json:"nodeIdentities" yaml:"nodeIdentities"`
	// The partition the ACL token is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#partition AclToken#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// The list of policies attached to the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#policies AclToken#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// The list of roles attached to the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#roles AclToken#roles}
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// service_identities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#service_identities AclToken#service_identities}
	ServiceIdentities interface{} `field:"optional" json:"serviceIdentities" yaml:"serviceIdentities"`
	// templated_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/acl_token#templated_policies AclToken#templated_policies}
	TemplatedPolicies interface{} `field:"optional" json:"templatedPolicies" yaml:"templatedPolicies"`
}

