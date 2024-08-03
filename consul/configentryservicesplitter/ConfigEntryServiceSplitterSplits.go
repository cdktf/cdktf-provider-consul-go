// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicesplitter


type ConfigEntryServiceSplitterSplits struct {
	// Specifies the name of the service to resolve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_splitter#service ConfigEntryServiceSplitter#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Specifies the percentage of traffic sent to the set of service instances specified in the `service` field.
	//
	// Each weight must be a floating integer between `0` and `100`. The smallest representable value is `.01`. The sum of weights across all splits must add up to `100`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_splitter#weight ConfigEntryServiceSplitter#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Specifies the namespace to use in the FQDN when resolving the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_splitter#namespace ConfigEntryServiceSplitter#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the admin partition to use in the FQDN when resolving the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_splitter#partition ConfigEntryServiceSplitter#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// request_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_splitter#request_headers ConfigEntryServiceSplitter#request_headers}
	RequestHeaders *ConfigEntryServiceSplitterSplitsRequestHeaders `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
	// response_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_splitter#response_headers ConfigEntryServiceSplitter#response_headers}
	ResponseHeaders *ConfigEntryServiceSplitterSplitsResponseHeaders `field:"optional" json:"responseHeaders" yaml:"responseHeaders"`
	// Specifies a subset of the service to resolve.
	//
	// A service subset assigns a name to a specific subset of discoverable service instances within a datacenter, such as `version2` or `canary`. All services have an unnamed default subset that returns all healthy instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_splitter#service_subset ConfigEntryServiceSplitter#service_subset}
	ServiceSubset *string `field:"optional" json:"serviceSubset" yaml:"serviceSubset"`
}

