// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceCheck struct {
	// An ID, *unique per agent*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#check_id Service#check_id}
	CheckId *string `field:"required" json:"checkId" yaml:"checkId"`
	// The interval to wait between each health-check invocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#interval Service#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The name of the health-check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#name Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies a timeout for outgoing connections in the case of a HTTP or TCP check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#timeout Service#timeout}
	Timeout *string `field:"required" json:"timeout" yaml:"timeout"`
	// The time after which the service is automatically deregistered when in the `critical` state. Defaults to `30s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#deregister_critical_service_after Service#deregister_critical_service_after}
	DeregisterCriticalServiceAfter *string `field:"optional" json:"deregisterCriticalServiceAfter" yaml:"deregisterCriticalServiceAfter"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#header Service#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// The HTTP endpoint to call for an HTTP check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#http Service#http}
	Http *string `field:"optional" json:"http" yaml:"http"`
	// The method to use for HTTP health-checks. Defaults to `GET`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#method Service#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// An opaque field meant to hold human readable text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#notes Service#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The initial health-check status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#status Service#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The TCP address and port to connect to for a TCP check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#tcp Service#tcp}
	Tcp *string `field:"optional" json:"tcp" yaml:"tcp"`
	// Whether to deactivate certificate verification for HTTP health-checks. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.19.0/docs/resources/service#tls_skip_verify Service#tls_skip_verify}
	TlsSkipVerify interface{} `field:"optional" json:"tlsSkipVerify" yaml:"tlsSkipVerify"`
}

