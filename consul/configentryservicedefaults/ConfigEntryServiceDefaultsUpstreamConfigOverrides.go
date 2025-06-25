// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults


type ConfigEntryServiceDefaultsUpstreamConfigOverrides struct {
	// Sets the strategy for allocating outbound connections from upstreams across Envoy proxy threads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#balance_outbound_connections ConfigEntryServiceDefaults#balance_outbound_connections}
	BalanceOutboundConnections *string `field:"optional" json:"balanceOutboundConnections" yaml:"balanceOutboundConnections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#connect_timeout_ms ConfigEntryServiceDefaults#connect_timeout_ms}.
	ConnectTimeoutMs *float64 `field:"optional" json:"connectTimeoutMs" yaml:"connectTimeoutMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#envoy_listener_json ConfigEntryServiceDefaults#envoy_listener_json}.
	EnvoyListenerJson *string `field:"optional" json:"envoyListenerJson" yaml:"envoyListenerJson"`
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#limits ConfigEntryServiceDefaults#limits}
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
	// mesh_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#mesh_gateway ConfigEntryServiceDefaults#mesh_gateway}
	MeshGateway interface{} `field:"optional" json:"meshGateway" yaml:"meshGateway"`
	// Specifies the name of the service you are setting the defaults for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#name ConfigEntryServiceDefaults#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the namespace containing the upstream service that the configuration applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#namespace ConfigEntryServiceDefaults#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the name of the name of the Consul admin partition that the configuration entry applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#partition ConfigEntryServiceDefaults#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// passive_health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#passive_health_check ConfigEntryServiceDefaults#passive_health_check}
	PassiveHealthCheck interface{} `field:"optional" json:"passiveHealthCheck" yaml:"passiveHealthCheck"`
	// Specifies the peer name of the upstream service that the configuration applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#peer ConfigEntryServiceDefaults#peer}
	Peer *string `field:"optional" json:"peer" yaml:"peer"`
	// Specifies the default protocol for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_defaults#protocol ConfigEntryServiceDefaults#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

