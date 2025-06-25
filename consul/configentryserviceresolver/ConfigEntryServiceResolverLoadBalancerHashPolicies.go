// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver


type ConfigEntryServiceResolverLoadBalancerHashPolicies struct {
	// cookie_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_resolver#cookie_config ConfigEntryServiceResolver#cookie_config}
	CookieConfig interface{} `field:"optional" json:"cookieConfig" yaml:"cookieConfig"`
	// Specifies the attribute type to hash on. You cannot specify the Field parameter if SourceIP is also configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_resolver#field ConfigEntryServiceResolver#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Specifies the value to hash, such as a header name, cookie name, or a URL query parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_resolver#field_value ConfigEntryServiceResolver#field_value}
	FieldValue *string `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// Determines if the hash type should be source IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_resolver#source_ip ConfigEntryServiceResolver#source_ip}
	SourceIp interface{} `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// Determines if Consul should stop computing the hash when multiple hash policies are present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_resolver#terminal ConfigEntryServiceResolver#terminal}
	Terminal interface{} `field:"optional" json:"terminal" yaml:"terminal"`
}

