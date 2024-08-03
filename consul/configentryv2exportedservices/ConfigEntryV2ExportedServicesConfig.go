// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryv2exportedservices

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryV2ExportedServicesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The kind of exported services config (ExportedServices, NamespaceExportedServices, PartitionExportedServices).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_v2_exported_services#kind ConfigEntryV2ExportedServices#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The name of the config entry to read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_v2_exported_services#name ConfigEntryV2ExportedServices#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The partition the config entry is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_v2_exported_services#partition ConfigEntryV2ExportedServices#partition}
	Partition *string `field:"required" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_v2_exported_services#id ConfigEntryV2ExportedServices#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The namespace the config entry is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_v2_exported_services#namespace ConfigEntryV2ExportedServices#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The exported service partition consumers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_v2_exported_services#partition_consumers ConfigEntryV2ExportedServices#partition_consumers}
	PartitionConsumers *[]*string `field:"optional" json:"partitionConsumers" yaml:"partitionConsumers"`
	// The exported service peer consumers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_v2_exported_services#peer_consumers ConfigEntryV2ExportedServices#peer_consumers}
	PeerConsumers *[]*string `field:"optional" json:"peerConsumers" yaml:"peerConsumers"`
	// The exported service sameness group consumers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_v2_exported_services#sameness_group_consumers ConfigEntryV2ExportedServices#sameness_group_consumers}
	SamenessGroupConsumers *[]*string `field:"optional" json:"samenessGroupConsumers" yaml:"samenessGroupConsumers"`
	// The exported services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_v2_exported_services#services ConfigEntryV2ExportedServices#services}
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
}

