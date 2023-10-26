// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceintentions


type ConfigEntryServiceIntentionsJwtProvidersVerifyClaims struct {
	// Specifies the path to the claim in the JSON web token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#path ConfigEntryServiceIntentions#path}
	Path *[]*string `field:"optional" json:"path" yaml:"path"`
	// Specifies the value to match on when verifying the the claim designated in path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_intentions#value ConfigEntryServiceIntentions#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

