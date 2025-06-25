// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicerouter


type ConfigEntryServiceRouterRoutesMatch struct {
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_router#http ConfigEntryServiceRouter#http}
	Http *ConfigEntryServiceRouterRoutesMatchHttp `field:"optional" json:"http" yaml:"http"`
}

