// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver


type ConfigEntryServiceResolverFailoverTargets struct {
	// Specifies the WAN federated datacenter to use for the failover target. If empty, the current datacenter is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#datacenter ConfigEntryServiceResolver#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Specifies the namespace to use for the failover target. If empty, the default namespace is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#namespace ConfigEntryServiceResolver#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the admin partition within the same datacenter to use for the failover target.
	//
	// If empty, the default partition is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#partition ConfigEntryServiceResolver#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Specifies the destination cluster peer to resolve the target service name from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#peer ConfigEntryServiceResolver#peer}
	Peer *string `field:"optional" json:"peer" yaml:"peer"`
	// Specifies the service name to use for the failover target. If empty, the current service name is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#service ConfigEntryServiceResolver#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// Specifies the named subset to use for the failover target.
	//
	// If empty, the default subset for the requested service name is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#service_subset ConfigEntryServiceResolver#service_subset}
	ServiceSubset *string `field:"optional" json:"serviceSubset" yaml:"serviceSubset"`
}

