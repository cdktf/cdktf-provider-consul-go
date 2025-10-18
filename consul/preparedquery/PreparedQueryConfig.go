// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package preparedquery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PreparedQueryConfig struct {
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
	// The name of the prepared query.
	//
	// Used to identify the prepared query during requests. Can be specified as an empty string to configure the query as a catch-all.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#name PreparedQuery#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the service to query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#service PreparedQuery#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// When `true` the prepared query will return connect proxy services for a queried service.
	//
	// Conditions such as `tags` in the prepared query will be matched against the proxy service. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#connect PreparedQuery#connect}
	Connect interface{} `field:"optional" json:"connect" yaml:"connect"`
	// The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#datacenter PreparedQuery#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#dns PreparedQuery#dns}
	Dns *PreparedQueryDns `field:"optional" json:"dns" yaml:"dns"`
	// failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#failover PreparedQuery#failover}
	Failover *PreparedQueryFailover `field:"optional" json:"failover" yaml:"failover"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#id PreparedQuery#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies a list of check IDs that should be ignored when filtering unhealthy instances.
	//
	// This is mostly useful in an emergency or as a temporary measure when a health check is found to be unreliable. Being able to ignore it in centrally-defined queries can be simpler than de-registering the check as an interim solution until the check can be fixed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#ignore_check_ids PreparedQuery#ignore_check_ids}
	IgnoreCheckIds *[]*string `field:"optional" json:"ignoreCheckIds" yaml:"ignoreCheckIds"`
	// Allows specifying the name of a node to sort results near using Consul's distance sorting and network coordinates.
	//
	// The magic `_agent` value can be used to always sort nearest the node servicing the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#near PreparedQuery#near}
	Near *string `field:"optional" json:"near" yaml:"near"`
	// Specifies a list of user-defined key/value pairs that will be used for filtering the query results to nodes with the given metadata values present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#node_meta PreparedQuery#node_meta}
	NodeMeta *map[string]*string `field:"optional" json:"nodeMeta" yaml:"nodeMeta"`
	// When `true`, the prepared query will only return nodes with passing health checks in the result.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#only_passing PreparedQuery#only_passing}
	OnlyPassing interface{} `field:"optional" json:"onlyPassing" yaml:"onlyPassing"`
	// Specifies a list of user-defined key/value pairs that will be used for filtering the query results to services with the given metadata values present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#service_meta PreparedQuery#service_meta}
	ServiceMeta *map[string]*string `field:"optional" json:"serviceMeta" yaml:"serviceMeta"`
	// The name of the Consul session to tie this query's lifetime to.
	//
	// This is an advanced parameter that should not be used without a complete understanding of Consul sessions and the implications of their use (it is recommended to leave this blank in nearly all cases).  If this parameter is omitted the query will not expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#session PreparedQuery#session}
	Session *string `field:"optional" json:"session" yaml:"session"`
	// The ACL token to store with the prepared query.
	//
	// This token will be used by default whenever the query is executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#stored_token PreparedQuery#stored_token}
	StoredToken *string `field:"optional" json:"storedToken" yaml:"storedToken"`
	// The list of required and/or disallowed tags.
	//
	// If a tag is in this list it must be present.  If the tag is preceded with a "!" then it is disallowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#tags PreparedQuery#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#template PreparedQuery#template}
	Template *PreparedQueryTemplate `field:"optional" json:"template" yaml:"template"`
	// The ACL token to use when saving the prepared query.
	//
	// This overrides the token that the agent provides by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.1/docs/resources/prepared_query#token PreparedQuery#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

