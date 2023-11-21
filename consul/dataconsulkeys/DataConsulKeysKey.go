// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataconsulkeys


type DataConsulKeysKey struct {
	// This is the name of the key.
	//
	// This value of the key is exposed as `var.<name>`. This is not the path of the key in Consul.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/keys#name DataConsulKeys#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// This is the path in Consul that should be read or written to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/keys#path DataConsulKeys#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// This is the default value to set for `var.<name>` if the key does not exist in Consul. Defaults to an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/keys#default DataConsulKeys#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
}

