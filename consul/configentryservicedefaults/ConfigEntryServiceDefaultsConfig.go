// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceDefaultsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// expose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#expose ConfigEntryServiceDefaults#expose}
	Expose interface{} `field:"required" json:"expose" yaml:"expose"`
	// Specifies the name of the service you are setting the defaults for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#name ConfigEntryServiceDefaults#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the default protocol for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#protocol ConfigEntryServiceDefaults#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Specifies the strategy for allocating inbound connections to the service across Envoy proxy threads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#balance_inbound_connections ConfigEntryServiceDefaults#balance_inbound_connections}
	BalanceInboundConnections *string `field:"optional" json:"balanceInboundConnections" yaml:"balanceInboundConnections"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#destination ConfigEntryServiceDefaults#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// envoy_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#envoy_extensions ConfigEntryServiceDefaults#envoy_extensions}
	EnvoyExtensions interface{} `field:"optional" json:"envoyExtensions" yaml:"envoyExtensions"`
	// Specifies the TLS server name indication (SNI) when federating with an external system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#external_sni ConfigEntryServiceDefaults#external_sni}
	ExternalSni *string `field:"optional" json:"externalSni" yaml:"externalSni"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#id ConfigEntryServiceDefaults#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the number of milliseconds allowed for establishing connections to the local application instance before timing out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#local_connect_timeout_ms ConfigEntryServiceDefaults#local_connect_timeout_ms}
	LocalConnectTimeoutMs *float64 `field:"optional" json:"localConnectTimeoutMs" yaml:"localConnectTimeoutMs"`
	// Specifies the timeout for HTTP requests to the local application instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#local_request_timeout_ms ConfigEntryServiceDefaults#local_request_timeout_ms}
	LocalRequestTimeoutMs *float64 `field:"optional" json:"localRequestTimeoutMs" yaml:"localRequestTimeoutMs"`
	// Specifies the maximum number of concurrent inbound connections to each service instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#max_inbound_connections ConfigEntryServiceDefaults#max_inbound_connections}
	MaxInboundConnections *float64 `field:"optional" json:"maxInboundConnections" yaml:"maxInboundConnections"`
	// mesh_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#mesh_gateway ConfigEntryServiceDefaults#mesh_gateway}
	MeshGateway interface{} `field:"optional" json:"meshGateway" yaml:"meshGateway"`
	// Specifies a set of custom key-value pairs to add to the Consul KV store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#meta ConfigEntryServiceDefaults#meta}
	Meta *map[string]*string `field:"optional" json:"meta" yaml:"meta"`
	// Specifies a mode for how the service directs inbound and outbound traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#mode ConfigEntryServiceDefaults#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Controls whether mutual TLS is required for incoming connections to this service.
	//
	// This setting is only supported for services with transparent proxy enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#mutual_tls_mode ConfigEntryServiceDefaults#mutual_tls_mode}
	MutualTlsMode *string `field:"optional" json:"mutualTlsMode" yaml:"mutualTlsMode"`
	// Specifies the Consul namespace that the configuration entry applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#namespace ConfigEntryServiceDefaults#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the name of the name of the Consul admin partition that the configuration entry applies to.
	//
	// Refer to Admin Partitions for additional information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#partition ConfigEntryServiceDefaults#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// transparent_proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#transparent_proxy ConfigEntryServiceDefaults#transparent_proxy}
	TransparentProxy interface{} `field:"optional" json:"transparentProxy" yaml:"transparentProxy"`
	// upstream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#upstream_config ConfigEntryServiceDefaults#upstream_config}
	UpstreamConfig interface{} `field:"optional" json:"upstreamConfig" yaml:"upstreamConfig"`
}

