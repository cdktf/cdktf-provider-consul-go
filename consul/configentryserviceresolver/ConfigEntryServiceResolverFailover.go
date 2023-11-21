// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver


type ConfigEntryServiceResolverFailover struct {
	// Name of subset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_resolver#subset_name ConfigEntryServiceResolver#subset_name}
	SubsetName *string `field:"required" json:"subsetName" yaml:"subsetName"`
	// Specifies an ordered list of datacenters at the failover location to attempt connections to during a failover scenario.
	//
	// When Consul cannot establish a connection with the first datacenter in the list, it proceeds sequentially until establishing a connection with another datacenter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_resolver#datacenters ConfigEntryServiceResolver#datacenters}
	Datacenters *[]*string `field:"optional" json:"datacenters" yaml:"datacenters"`
	// Specifies the namespace at the failover location where the failover services are deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_resolver#namespace ConfigEntryServiceResolver#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the sameness group at the failover location where the failover services are deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_resolver#sameness_group ConfigEntryServiceResolver#sameness_group}
	SamenessGroup *string `field:"optional" json:"samenessGroup" yaml:"samenessGroup"`
	// Specifies the name of the service to resolve at the failover location during a failover scenario.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_resolver#service ConfigEntryServiceResolver#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// Specifies the name of a subset of service instances to resolve at the failover location during a failover scenario.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_resolver#service_subset ConfigEntryServiceResolver#service_subset}
	ServiceSubset *string `field:"optional" json:"serviceSubset" yaml:"serviceSubset"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.20.0/docs/resources/config_entry_service_resolver#targets ConfigEntryServiceResolver#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

