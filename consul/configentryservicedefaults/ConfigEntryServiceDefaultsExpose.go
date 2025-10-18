// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults


type ConfigEntryServiceDefaultsExpose struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_defaults#checks ConfigEntryServiceDefaults#checks}.
	Checks interface{} `field:"optional" json:"checks" yaml:"checks"`
	// paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_defaults#paths ConfigEntryServiceDefaults#paths}
	Paths interface{} `field:"optional" json:"paths" yaml:"paths"`
}

