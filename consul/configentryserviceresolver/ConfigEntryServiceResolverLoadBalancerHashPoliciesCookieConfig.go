// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver


type ConfigEntryServiceResolverLoadBalancerHashPoliciesCookieConfig struct {
	// Specifies the path to set for the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_resolver#path ConfigEntryServiceResolver#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Directs Consul to generate a session cookie with no expiration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_resolver#session ConfigEntryServiceResolver#session}
	Session interface{} `field:"optional" json:"session" yaml:"session"`
	// Specifies the TTL for generated cookies. Cannot be specified for session cookies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_resolver#ttl ConfigEntryServiceResolver#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

