// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver


type ConfigEntryServiceResolverLoadBalancer struct {
	// hash_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_resolver#hash_policies ConfigEntryServiceResolver#hash_policies}
	HashPolicies interface{} `field:"optional" json:"hashPolicies" yaml:"hashPolicies"`
	// least_request_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_resolver#least_request_config ConfigEntryServiceResolver#least_request_config}
	LeastRequestConfig interface{} `field:"optional" json:"leastRequestConfig" yaml:"leastRequestConfig"`
	// Specifies the type of load balancing policy for selecting a host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_resolver#policy ConfigEntryServiceResolver#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// ring_hash_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_resolver#ring_hash_config ConfigEntryServiceResolver#ring_hash_config}
	RingHashConfig interface{} `field:"optional" json:"ringHashConfig" yaml:"ringHashConfig"`
}

