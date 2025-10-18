// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceintentions


type ConfigEntryServiceIntentionsJwtProviders struct {
	// Specifies the name of a JWT provider defined in the Name field of the jwt-provider configuration entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_intentions#name ConfigEntryServiceIntentions#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// verify_claims block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/config_entry_service_intentions#verify_claims ConfigEntryServiceIntentions#verify_claims}
	VerifyClaims interface{} `field:"optional" json:"verifyClaims" yaml:"verifyClaims"`
}

