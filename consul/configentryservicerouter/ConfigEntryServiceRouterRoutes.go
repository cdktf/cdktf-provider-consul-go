// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicerouter


type ConfigEntryServiceRouterRoutes struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_router#destination ConfigEntryServiceRouter#destination}
	Destination *ConfigEntryServiceRouterRoutesDestination `field:"optional" json:"destination" yaml:"destination"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_router#match ConfigEntryServiceRouter#match}
	Match *ConfigEntryServiceRouterRoutesMatch `field:"optional" json:"match" yaml:"match"`
}

