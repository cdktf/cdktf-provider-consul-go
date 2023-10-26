// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package preparedquery


type PreparedQueryTemplate struct {
	// The regular expression to match with. When using `name_prefix_match`, this regex is applied against the query name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/prepared_query#regexp PreparedQuery#regexp}
	Regexp *string `field:"required" json:"regexp" yaml:"regexp"`
	// The type of template matching to perform. Currently only `name_prefix_match` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/prepared_query#type PreparedQuery#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// If set to true, will cause the tags list inside the service structure to be stripped of any empty strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/prepared_query#remove_empty_tags PreparedQuery#remove_empty_tags}
	RemoveEmptyTags interface{} `field:"optional" json:"removeEmptyTags" yaml:"removeEmptyTags"`
}

