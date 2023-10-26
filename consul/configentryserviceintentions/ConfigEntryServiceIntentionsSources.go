// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceintentions


type ConfigEntryServiceIntentionsSources struct {
	// Specifies the action to take when the source sends traffic to the destination service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#action ConfigEntryServiceIntentions#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Specifies a description of the intention.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#description ConfigEntryServiceIntentions#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the name of the source that the intention allows or denies traffic from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#name ConfigEntryServiceIntentions#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the traffic source namespace that the intention allows or denies traffic from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#namespace ConfigEntryServiceIntentions#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the name of an admin partition that the intention allows or denies traffic from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#partition ConfigEntryServiceIntentions#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Specifies the name of a peered Consul cluster that the intention allows or denies traffic from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#peer ConfigEntryServiceIntentions#peer}
	Peer *string `field:"optional" json:"peer" yaml:"peer"`
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#permissions ConfigEntryServiceIntentions#permissions}
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// The Precedence field contains a read-only integer.
	//
	// Consul generates the value based on name configurations for the source and destination services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#precedence ConfigEntryServiceIntentions#precedence}
	Precedence *float64 `field:"optional" json:"precedence" yaml:"precedence"`
	// Specifies the name of a sameness group that the intention allows or denies traffic from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#sameness_group ConfigEntryServiceIntentions#sameness_group}
	SamenessGroup *string `field:"optional" json:"samenessGroup" yaml:"samenessGroup"`
	// Specifies the type of destination service that the configuration entry applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#type ConfigEntryServiceIntentions#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

