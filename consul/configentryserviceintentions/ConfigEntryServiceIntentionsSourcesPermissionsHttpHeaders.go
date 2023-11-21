// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceintentions


type ConfigEntryServiceIntentionsSourcesPermissionsHttpHeaders struct {
	// Specifies the name of the header to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_intentions#name ConfigEntryServiceIntentions#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies a value for the header key set in the Name field.
	//
	// If the request header value matches the Exact value, Consul applies the permission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_intentions#exact ConfigEntryServiceIntentions#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Inverts the matching logic configured in the Header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_intentions#invert ConfigEntryServiceIntentions#invert}
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// Specifies a prefix value for the header key set in the Name field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_intentions#prefix ConfigEntryServiceIntentions#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Enables a match if the header configured in the Name field appears in the request.
	//
	// Consul matches on any value as long as the header key appears in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_intentions#present ConfigEntryServiceIntentions#present}
	Present interface{} `field:"optional" json:"present" yaml:"present"`
	// Specifies a regular expression pattern as the value for the header key set in the Name field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_intentions#regex ConfigEntryServiceIntentions#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Specifies a suffix value for the header key set in the Name field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_intentions#suffix ConfigEntryServiceIntentions#suffix}
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

