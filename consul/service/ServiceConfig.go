// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceConfig struct {
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
	// The name of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#name Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the node the to register the service on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#node Service#node}
	NodeAttribute *string `field:"required" json:"nodeAttribute" yaml:"nodeAttribute"`
	// The address of the service. Defaults to the address of the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#address Service#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#check Service#check}
	Check interface{} `field:"optional" json:"check" yaml:"check"`
	// The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#datacenter Service#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Specifies to disable the anti-entropy feature for this service's tags. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#enable_tag_override Service#enable_tag_override}
	EnableTagOverride interface{} `field:"optional" json:"enableTagOverride" yaml:"enableTagOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#external Service#external}.
	External interface{} `field:"optional" json:"external" yaml:"external"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#id Service#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A map of arbitrary KV metadata linked to the service instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#meta Service#meta}
	Meta *map[string]*string `field:"optional" json:"meta" yaml:"meta"`
	// The namespace to create the service within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#namespace Service#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The partition the service is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#partition Service#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// The port of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#port Service#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// If the service ID is not provided, it will be defaulted to the value of the `name` attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#service_id Service#service_id}
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
	// A list of values that are opaque to Consul, but can be used to distinguish between services or nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#tags Service#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

