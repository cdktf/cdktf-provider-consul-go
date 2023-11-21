// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataconsulcatalogservice


type DataConsulCatalogServiceQueryOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/catalog_service#allow_stale DataConsulCatalogService#allow_stale}.
	AllowStale interface{} `field:"optional" json:"allowStale" yaml:"allowStale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/catalog_service#datacenter DataConsulCatalogService#datacenter}.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/catalog_service#namespace DataConsulCatalogService#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/catalog_service#near DataConsulCatalogService#near}.
	Near *string `field:"optional" json:"near" yaml:"near"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/catalog_service#node_meta DataConsulCatalogService#node_meta}.
	NodeMeta *map[string]*string `field:"optional" json:"nodeMeta" yaml:"nodeMeta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/catalog_service#partition DataConsulCatalogService#partition}.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/catalog_service#require_consistent DataConsulCatalogService#require_consistent}.
	RequireConsistent interface{} `field:"optional" json:"requireConsistent" yaml:"requireConsistent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/catalog_service#token DataConsulCatalogService#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/catalog_service#wait_index DataConsulCatalogService#wait_index}.
	WaitIndex *float64 `field:"optional" json:"waitIndex" yaml:"waitIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/data-sources/catalog_service#wait_time DataConsulCatalogService#wait_time}.
	WaitTime *string `field:"optional" json:"waitTime" yaml:"waitTime"`
}

