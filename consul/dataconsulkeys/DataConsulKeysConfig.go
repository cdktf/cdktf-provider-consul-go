// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataconsulkeys

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataConsulKeysConfig struct {
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
	// The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/data-sources/keys#datacenter DataConsulKeys#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Whether to return an error when a key is absent from the KV store and no default is configured.
	//
	// This defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/data-sources/keys#error_on_missing_keys DataConsulKeys#error_on_missing_keys}
	ErrorOnMissingKeys interface{} `field:"optional" json:"errorOnMissingKeys" yaml:"errorOnMissingKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/data-sources/keys#id DataConsulKeys#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/data-sources/keys#key DataConsulKeys#key}
	Key interface{} `field:"optional" json:"key" yaml:"key"`
	// The namespace to lookup the keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/data-sources/keys#namespace DataConsulKeys#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The partition to lookup the keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/data-sources/keys#partition DataConsulKeys#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// The ACL token to use. This overrides the token that the agent provides by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/data-sources/keys#token DataConsulKeys#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

