// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver


type ConfigEntryServiceResolverLoadBalancerRingHashConfig struct {
	// Determines the maximum number of entries in the hash ring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#maximum_ring_size ConfigEntryServiceResolver#maximum_ring_size}
	MaximumRingSize *float64 `field:"optional" json:"maximumRingSize" yaml:"maximumRingSize"`
	// Determines the minimum number of entries in the hash ring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#minimum_ring_size ConfigEntryServiceResolver#minimum_ring_size}
	MinimumRingSize *float64 `field:"optional" json:"minimumRingSize" yaml:"minimumRingSize"`
}

