// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults


type ConfigEntryServiceDefaultsUpstreamConfigDefaults struct {
	// Sets the strategy for allocating outbound connections from upstreams across Envoy proxy threads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#balance_outbound_connections ConfigEntryServiceDefaults#balance_outbound_connections}
	BalanceOutboundConnections *string `field:"optional" json:"balanceOutboundConnections" yaml:"balanceOutboundConnections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#connect_timeout_ms ConfigEntryServiceDefaults#connect_timeout_ms}.
	ConnectTimeoutMs *float64 `field:"optional" json:"connectTimeoutMs" yaml:"connectTimeoutMs"`
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#limits ConfigEntryServiceDefaults#limits}
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
	// mesh_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#mesh_gateway ConfigEntryServiceDefaults#mesh_gateway}
	MeshGateway interface{} `field:"optional" json:"meshGateway" yaml:"meshGateway"`
	// passive_health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#passive_health_check ConfigEntryServiceDefaults#passive_health_check}
	PassiveHealthCheck interface{} `field:"optional" json:"passiveHealthCheck" yaml:"passiveHealthCheck"`
	// Specifies the default protocol for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_defaults#protocol ConfigEntryServiceDefaults#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

