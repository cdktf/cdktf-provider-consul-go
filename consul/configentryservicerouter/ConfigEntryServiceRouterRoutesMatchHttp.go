// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicerouter


type ConfigEntryServiceRouterRoutesMatchHttp struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_router#header ConfigEntryServiceRouter#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Specifies HTTP methods that the match applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_router#methods ConfigEntryServiceRouter#methods}
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// Specifies the exact path to match on the HTTP request path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_router#path_exact ConfigEntryServiceRouter#path_exact}
	PathExact *string `field:"optional" json:"pathExact" yaml:"pathExact"`
	// Specifies the path prefix to match on the HTTP request path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_router#path_prefix ConfigEntryServiceRouter#path_prefix}
	PathPrefix *string `field:"optional" json:"pathPrefix" yaml:"pathPrefix"`
	// Specifies a regular expression to match on the HTTP request path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_router#path_regex ConfigEntryServiceRouter#path_regex}
	PathRegex *string `field:"optional" json:"pathRegex" yaml:"pathRegex"`
	// query_param block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_router#query_param ConfigEntryServiceRouter#query_param}
	QueryParam interface{} `field:"optional" json:"queryParam" yaml:"queryParam"`
}

