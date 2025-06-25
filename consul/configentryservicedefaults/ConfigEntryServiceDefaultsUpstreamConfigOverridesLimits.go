// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults


type ConfigEntryServiceDefaultsUpstreamConfigOverridesLimits struct {
	// Specifies the maximum number of concurrent requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#max_concurrent_requests ConfigEntryServiceDefaults#max_concurrent_requests}
	MaxConcurrentRequests *float64 `field:"optional" json:"maxConcurrentRequests" yaml:"maxConcurrentRequests"`
	// Specifies the maximum number of connections a service instance can establish against the upstream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#max_connections ConfigEntryServiceDefaults#max_connections}
	MaxConnections *float64 `field:"optional" json:"maxConnections" yaml:"maxConnections"`
	// Specifies the maximum number of requests that are queued while waiting for a connection to establish.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#max_pending_requests ConfigEntryServiceDefaults#max_pending_requests}
	MaxPendingRequests *float64 `field:"optional" json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

