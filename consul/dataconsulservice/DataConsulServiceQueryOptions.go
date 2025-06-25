// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataconsulservice


type DataConsulServiceQueryOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/data-sources/service#allow_stale DataConsulService#allow_stale}.
	AllowStale interface{} `field:"optional" json:"allowStale" yaml:"allowStale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/data-sources/service#datacenter DataConsulService#datacenter}.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/data-sources/service#namespace DataConsulService#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/data-sources/service#near DataConsulService#near}.
	Near *string `field:"optional" json:"near" yaml:"near"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/data-sources/service#node_meta DataConsulService#node_meta}.
	NodeMeta *map[string]*string `field:"optional" json:"nodeMeta" yaml:"nodeMeta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/data-sources/service#partition DataConsulService#partition}.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/data-sources/service#require_consistent DataConsulService#require_consistent}.
	RequireConsistent interface{} `field:"optional" json:"requireConsistent" yaml:"requireConsistent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/data-sources/service#token DataConsulService#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/data-sources/service#wait_index DataConsulService#wait_index}.
	WaitIndex *float64 `field:"optional" json:"waitIndex" yaml:"waitIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/data-sources/service#wait_time DataConsulService#wait_time}.
	WaitTime *string `field:"optional" json:"waitTime" yaml:"waitTime"`
}

