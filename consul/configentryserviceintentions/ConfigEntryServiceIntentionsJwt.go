// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceintentions


type ConfigEntryServiceIntentionsJwt struct {
	// providers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_intentions#providers ConfigEntryServiceIntentions#providers}
	Providers interface{} `field:"optional" json:"providers" yaml:"providers"`
}

