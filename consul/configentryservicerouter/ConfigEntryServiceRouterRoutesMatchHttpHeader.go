// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicerouter


type ConfigEntryServiceRouterRoutesMatchHttpHeader struct {
	// Specifies that a request matches when the header with the given name is this exact value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#exact ConfigEntryServiceRouter#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Specifies that the logic for the HTTP header match should be inverted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#invert ConfigEntryServiceRouter#invert}
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// Specifies the name of the HTTP header to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#name ConfigEntryServiceRouter#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies that a request matches when the header with the given name has this prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#prefix ConfigEntryServiceRouter#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies that a request matches when the value in the `name` argument is present anywhere in the HTTP header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#present ConfigEntryServiceRouter#present}
	Present interface{} `field:"optional" json:"present" yaml:"present"`
	// Specifies that a request matches when the header with the given name matches this regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#regex ConfigEntryServiceRouter#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Specifies that a request matches when the header with the given name has this suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#suffix ConfigEntryServiceRouter#suffix}
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

