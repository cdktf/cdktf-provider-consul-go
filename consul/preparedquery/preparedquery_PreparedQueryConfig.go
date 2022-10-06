package preparedquery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PreparedQueryConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#name PreparedQuery#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#service PreparedQuery#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#connect PreparedQuery#connect}.
	Connect interface{} `field:"optional" json:"connect" yaml:"connect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#datacenter PreparedQuery#datacenter}.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#dns PreparedQuery#dns}
	Dns *PreparedQueryDns `field:"optional" json:"dns" yaml:"dns"`
	// failover block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#failover PreparedQuery#failover}
	Failover *PreparedQueryFailover `field:"optional" json:"failover" yaml:"failover"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#id PreparedQuery#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#ignore_check_ids PreparedQuery#ignore_check_ids}.
	IgnoreCheckIds *[]*string `field:"optional" json:"ignoreCheckIds" yaml:"ignoreCheckIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#near PreparedQuery#near}.
	Near *string `field:"optional" json:"near" yaml:"near"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#node_meta PreparedQuery#node_meta}.
	NodeMeta *map[string]*string `field:"optional" json:"nodeMeta" yaml:"nodeMeta"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#only_passing PreparedQuery#only_passing}.
	OnlyPassing interface{} `field:"optional" json:"onlyPassing" yaml:"onlyPassing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#service_meta PreparedQuery#service_meta}.
	ServiceMeta *map[string]*string `field:"optional" json:"serviceMeta" yaml:"serviceMeta"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#session PreparedQuery#session}.
	Session *string `field:"optional" json:"session" yaml:"session"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#stored_token PreparedQuery#stored_token}.
	StoredToken *string `field:"optional" json:"storedToken" yaml:"storedToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#tags PreparedQuery#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#template PreparedQuery#template}
	Template *PreparedQueryTemplate `field:"optional" json:"template" yaml:"template"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#token PreparedQuery#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

