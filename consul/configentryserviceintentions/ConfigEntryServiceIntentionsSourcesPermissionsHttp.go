// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceintentions


type ConfigEntryServiceIntentionsSourcesPermissionsHttp struct {
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_intentions#headers ConfigEntryServiceIntentions#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Specifies a list of HTTP methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_intentions#methods ConfigEntryServiceIntentions#methods}
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// Specifies an exact path to match on the HTTP request path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_intentions#path_exact ConfigEntryServiceIntentions#path_exact}
	PathExact *string `field:"optional" json:"pathExact" yaml:"pathExact"`
	// Specifies a path prefix to match on the HTTP request path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_intentions#path_prefix ConfigEntryServiceIntentions#path_prefix}
	PathPrefix *string `field:"optional" json:"pathPrefix" yaml:"pathPrefix"`
	// Defines a regular expression to match on the HTTP request path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_intentions#path_regex ConfigEntryServiceIntentions#path_regex}
	PathRegex *string `field:"optional" json:"pathRegex" yaml:"pathRegex"`
}

