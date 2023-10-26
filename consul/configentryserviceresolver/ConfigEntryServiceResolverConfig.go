// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceResolverConfig struct {
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
	// Specifies a name for the configuration entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#name ConfigEntryServiceResolver#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the timeout duration for establishing new network connections to this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#connect_timeout ConfigEntryServiceResolver#connect_timeout}
	ConnectTimeout *string `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// Specifies a defined subset of service instances to use when no explicit subset is requested.
	//
	// If this parameter is not specified, Consul uses the unnamed default subset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#default_subset ConfigEntryServiceResolver#default_subset}
	DefaultSubset *string `field:"optional" json:"defaultSubset" yaml:"defaultSubset"`
	// failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#failover ConfigEntryServiceResolver#failover}
	Failover interface{} `field:"optional" json:"failover" yaml:"failover"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#id ConfigEntryServiceResolver#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#load_balancer ConfigEntryServiceResolver#load_balancer}
	LoadBalancer interface{} `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// Specifies key-value pairs to add to the KV store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#meta ConfigEntryServiceResolver#meta}
	Meta *map[string]*string `field:"optional" json:"meta" yaml:"meta"`
	// Specifies the namespace that the service resolver applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#namespace ConfigEntryServiceResolver#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the admin partition that the service resolver applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#partition ConfigEntryServiceResolver#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#redirect ConfigEntryServiceResolver#redirect}
	Redirect interface{} `field:"optional" json:"redirect" yaml:"redirect"`
	// Specifies the timeout duration for receiving an HTTP response from this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#request_timeout ConfigEntryServiceResolver#request_timeout}
	RequestTimeout *string `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// subsets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#subsets ConfigEntryServiceResolver#subsets}
	Subsets interface{} `field:"optional" json:"subsets" yaml:"subsets"`
}

