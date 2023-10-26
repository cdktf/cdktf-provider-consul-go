// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autopilotconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutopilotConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/autopilot_config#cleanup_dead_servers AutopilotConfig#cleanup_dead_servers}.
	CleanupDeadServers interface{} `field:"optional" json:"cleanupDeadServers" yaml:"cleanupDeadServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/autopilot_config#datacenter AutopilotConfig#datacenter}.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/autopilot_config#disable_upgrade_migration AutopilotConfig#disable_upgrade_migration}.
	DisableUpgradeMigration interface{} `field:"optional" json:"disableUpgradeMigration" yaml:"disableUpgradeMigration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/autopilot_config#id AutopilotConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/autopilot_config#last_contact_threshold AutopilotConfig#last_contact_threshold}.
	LastContactThreshold *string `field:"optional" json:"lastContactThreshold" yaml:"lastContactThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/autopilot_config#max_trailing_logs AutopilotConfig#max_trailing_logs}.
	MaxTrailingLogs *float64 `field:"optional" json:"maxTrailingLogs" yaml:"maxTrailingLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/autopilot_config#redundancy_zone_tag AutopilotConfig#redundancy_zone_tag}.
	RedundancyZoneTag *string `field:"optional" json:"redundancyZoneTag" yaml:"redundancyZoneTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/autopilot_config#server_stabilization_time AutopilotConfig#server_stabilization_time}.
	ServerStabilizationTime *string `field:"optional" json:"serverStabilizationTime" yaml:"serverStabilizationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/autopilot_config#upgrade_version_tag AutopilotConfig#upgrade_version_tag}.
	UpgradeVersionTag *string `field:"optional" json:"upgradeVersionTag" yaml:"upgradeVersionTag"`
}

