// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults


type ConfigEntryServiceDefaultsUpstreamConfig struct {
	// defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_defaults#defaults ConfigEntryServiceDefaults#defaults}
	Defaults interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_defaults#overrides ConfigEntryServiceDefaults#overrides}
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

