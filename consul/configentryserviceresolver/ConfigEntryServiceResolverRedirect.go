// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver


type ConfigEntryServiceResolverRedirect struct {
	// Specifies the datacenter at the redirect’s destination that resolves local upstream requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#datacenter ConfigEntryServiceResolver#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Specifies the namespace at the redirect’s destination that resolves local upstream requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#namespace ConfigEntryServiceResolver#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the admin partition at the redirect’s destination that resolves local upstream requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#partition ConfigEntryServiceResolver#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Specifies the cluster with an active cluster peering connection at the redirect’s destination that resolves local upstream requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#peer ConfigEntryServiceResolver#peer}
	Peer *string `field:"optional" json:"peer" yaml:"peer"`
	// Specifies the sameness group at the redirect’s destination that resolves local upstream requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#sameness_group ConfigEntryServiceResolver#sameness_group}
	SamenessGroup *string `field:"optional" json:"samenessGroup" yaml:"samenessGroup"`
	// Specifies the name of a service at the redirect’s destination that resolves local upstream requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#service ConfigEntryServiceResolver#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// Specifies the name of a subset of services at the redirect’s destination that resolves local upstream requests.
	//
	// If empty, the default subset is used. If specified, you must also specify at least one of the following in the same Redirect map: Service, Namespace, andDatacenter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/config_entry_service_resolver#service_subset ConfigEntryServiceResolver#service_subset}
	ServiceSubset *string `field:"optional" json:"serviceSubset" yaml:"serviceSubset"`
}

