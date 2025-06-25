// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceintentions


type ConfigEntryServiceIntentionsSourcesPermissions struct {
	// Specifies the action to take when the source sends traffic to the destination service.
	//
	// The value is either allow or deny.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_intentions#action ConfigEntryServiceIntentions#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/config_entry_service_intentions#http ConfigEntryServiceIntentions#http}
	Http interface{} `field:"required" json:"http" yaml:"http"`
}

