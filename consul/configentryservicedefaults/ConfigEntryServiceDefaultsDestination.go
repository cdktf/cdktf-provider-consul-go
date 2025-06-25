// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults


type ConfigEntryServiceDefaultsDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#addresses ConfigEntryServiceDefaults#addresses}.
	Addresses *[]*string `field:"required" json:"addresses" yaml:"addresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#port ConfigEntryServiceDefaults#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

