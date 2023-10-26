// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicerouter


type ConfigEntryServiceRouterRoutesMatchHttpQueryParam struct {
	// Specifies that a request matches when the query parameter with the given name is this exact value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#exact ConfigEntryServiceRouter#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Specifies the name of the HTTP query parameter to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#name ConfigEntryServiceRouter#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies that a request matches when the value in the `name` argument is present anywhere in the HTTP query parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#present ConfigEntryServiceRouter#present}
	Present interface{} `field:"optional" json:"present" yaml:"present"`
	// Specifies that a request matches when the query parameter with the given name matches this regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_router#regex ConfigEntryServiceRouter#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

