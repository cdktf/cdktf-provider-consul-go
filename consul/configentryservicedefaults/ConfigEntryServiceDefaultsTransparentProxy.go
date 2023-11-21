// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults


type ConfigEntryServiceDefaultsTransparentProxy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_defaults#dialed_directly ConfigEntryServiceDefaults#dialed_directly}.
	DialedDirectly interface{} `field:"required" json:"dialedDirectly" yaml:"dialedDirectly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_defaults#outbound_listener_port ConfigEntryServiceDefaults#outbound_listener_port}.
	OutboundListenerPort *float64 `field:"required" json:"outboundListenerPort" yaml:"outboundListenerPort"`
}

