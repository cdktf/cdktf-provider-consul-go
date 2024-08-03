// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicerouter


type ConfigEntryServiceRouterRoutesDestination struct {
	// Specifies the total amount of time permitted for the request stream to be idle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#idle_timeout ConfigEntryServiceRouter#idle_timeout}
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Specifies the Consul namespace to resolve the service from instead of the current namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#namespace ConfigEntryServiceRouter#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the number of times to retry the request when a retry condition occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#num_retries ConfigEntryServiceRouter#num_retries}
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// Specifies the Consul admin partition to resolve the service from instead of the current partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#partition ConfigEntryServiceRouter#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Specifies rewrites to the HTTP request path before proxying it to its final destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#prefix_rewrite ConfigEntryServiceRouter#prefix_rewrite}
	PrefixRewrite *string `field:"optional" json:"prefixRewrite" yaml:"prefixRewrite"`
	// request_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#request_headers ConfigEntryServiceRouter#request_headers}
	RequestHeaders *ConfigEntryServiceRouterRoutesDestinationRequestHeaders `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
	// Specifies the total amount of time permitted for the entire downstream request to be processed, including retry attempts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#request_timeout ConfigEntryServiceRouter#request_timeout}
	RequestTimeout *string `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// response_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#response_headers ConfigEntryServiceRouter#response_headers}
	ResponseHeaders *ConfigEntryServiceRouterRoutesDestinationResponseHeaders `field:"optional" json:"responseHeaders" yaml:"responseHeaders"`
	// Specifies a list of conditions for Consul to retry requests based on the response from an upstream service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#retry_on ConfigEntryServiceRouter#retry_on}
	RetryOn *[]*string `field:"optional" json:"retryOn" yaml:"retryOn"`
	// Specifies that connection failure errors that trigger a retry request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#retry_on_connect_failure ConfigEntryServiceRouter#retry_on_connect_failure}
	RetryOnConnectFailure interface{} `field:"optional" json:"retryOnConnectFailure" yaml:"retryOnConnectFailure"`
	// Specifies a list of integers for HTTP response status codes that trigger a retry request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#retry_on_status_codes ConfigEntryServiceRouter#retry_on_status_codes}
	RetryOnStatusCodes *[]*float64 `field:"optional" json:"retryOnStatusCodes" yaml:"retryOnStatusCodes"`
	// Specifies the name of the service to resolve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#service ConfigEntryServiceRouter#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// Specifies a named subset of the given service to resolve instead of the one defined as that service's `default_subset` in the service resolver configuration entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_router#service_subset ConfigEntryServiceRouter#service_subset}
	ServiceSubset *string `field:"optional" json:"serviceSubset" yaml:"serviceSubset"`
}

