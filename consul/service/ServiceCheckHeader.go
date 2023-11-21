// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceCheckHeader struct {
	// The name of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#name Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header's list of values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/service#value Service#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

