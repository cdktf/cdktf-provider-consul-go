// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults


type ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheck struct {
	// Specifies the minimum amount of time that an ejected host must remain outside the cluster before rejoining.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#base_ejection_time ConfigEntryServiceDefaults#base_ejection_time}
	BaseEjectionTime *string `field:"optional" json:"baseEjectionTime" yaml:"baseEjectionTime"`
	// Specifies a percentage that indicates how many times out of 100 that Consul ejects the host when it detects an outlier status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#enforcing_consecutive_5xx ConfigEntryServiceDefaults#enforcing_consecutive_5xx}
	EnforcingConsecutive5Xx *float64 `field:"optional" json:"enforcingConsecutive5Xx" yaml:"enforcingConsecutive5Xx"`
	// Specifies the time between checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#interval ConfigEntryServiceDefaults#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Specifies the maximum percentage of an upstream cluster that Consul ejects when the proxy reports an outlier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#max_ejection_percent ConfigEntryServiceDefaults#max_ejection_percent}
	MaxEjectionPercent *float64 `field:"optional" json:"maxEjectionPercent" yaml:"maxEjectionPercent"`
	// Specifies the number of consecutive failures allowed per check interval.
	//
	// If exceeded, Consul removes the host from the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults#max_failures ConfigEntryServiceDefaults#max_failures}
	MaxFailures *float64 `field:"optional" json:"maxFailures" yaml:"maxFailures"`
}

