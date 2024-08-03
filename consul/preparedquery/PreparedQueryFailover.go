// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package preparedquery


type PreparedQueryFailover struct {
	// Remote datacenters to return results from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/prepared_query#datacenters PreparedQuery#datacenters}
	Datacenters *[]*string `field:"optional" json:"datacenters" yaml:"datacenters"`
	// Return results from this many datacenters, sorted in ascending order of estimated RTT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/prepared_query#nearest_n PreparedQuery#nearest_n}
	NearestN *float64 `field:"optional" json:"nearestN" yaml:"nearestN"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/prepared_query#targets PreparedQuery#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

