// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicesplitter


type ConfigEntryServiceSplitterSplitsRequestHeaders struct {
	// Map of one or more key-value pairs.
	//
	// Defines a set of key-value pairs to add to the header. Use header names as the keys. Header names are not case-sensitive. If header values with the same name already exist, the value is appended and Consul applies both headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_splitter#add ConfigEntryServiceSplitter#add}
	Add *map[string]*string `field:"optional" json:"add" yaml:"add"`
	// Defines an list of headers to remove. Consul removes only headers containing exact matches. Header names are not case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_splitter#remove ConfigEntryServiceSplitter#remove}
	Remove *[]*string `field:"optional" json:"remove" yaml:"remove"`
	// Map of one or more key-value pairs.
	//
	// Defines a set of key-value pairs to add to the request header or to replace existing header values with. Use header names as the keys. Header names are not case-sensitive. If header values with the same names already exist, Consul replaces the header values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_splitter#set ConfigEntryServiceSplitter#set}
	Set *map[string]*string `field:"optional" json:"set" yaml:"set"`
}

