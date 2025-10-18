// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicerouter


type ConfigEntryServiceRouterRoutesDestinationResponseHeaders struct {
	// Defines a set of key-value pairs to add to the header. Use header names as the keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_router#add ConfigEntryServiceRouter#add}
	Add *map[string]*string `field:"optional" json:"add" yaml:"add"`
	// Defines a list of headers to remove.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_router#remove ConfigEntryServiceRouter#remove}
	Remove *[]*string `field:"optional" json:"remove" yaml:"remove"`
	// Defines a set of key-value pairs to add to the response header or to replace existing header values with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_router#set ConfigEntryServiceRouter#set}
	Set *map[string]*string `field:"optional" json:"set" yaml:"set"`
}

