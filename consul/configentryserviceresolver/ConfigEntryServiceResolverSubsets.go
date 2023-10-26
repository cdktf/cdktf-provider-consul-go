// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver


type ConfigEntryServiceResolverSubsets struct {
	// Specifies an expression that filters the DNS elements of service instances that belong to the subset.
	//
	// If empty, all healthy instances of a service are returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#filter ConfigEntryServiceResolver#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Name of subset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#name ConfigEntryServiceResolver#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Determines if instances that return a warning from a health check are allowed to resolve a request.
	//
	// When set to false, instances with passing and warning states are considered healthy. When set to true, only instances with a passing health check state are considered healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#only_passing ConfigEntryServiceResolver#only_passing}
	OnlyPassing interface{} `field:"required" json:"onlyPassing" yaml:"onlyPassing"`
}

