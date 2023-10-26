// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults


type ConfigEntryServiceDefaultsEnvoyExtensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#arguments ConfigEntryServiceDefaults#arguments}.
	Arguments *map[string]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#consul_version ConfigEntryServiceDefaults#consul_version}.
	ConsulVersion *string `field:"optional" json:"consulVersion" yaml:"consulVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#envoy_version ConfigEntryServiceDefaults#envoy_version}.
	EnvoyVersion *string `field:"optional" json:"envoyVersion" yaml:"envoyVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#name ConfigEntryServiceDefaults#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#required ConfigEntryServiceDefaults#required}.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

