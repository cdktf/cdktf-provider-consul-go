package dataconsulservice


type DataConsulServiceQueryOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/service#allow_stale DataConsulService#allow_stale}.
	AllowStale interface{} `field:"optional" json:"allowStale" yaml:"allowStale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/service#datacenter DataConsulService#datacenter}.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/service#namespace DataConsulService#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/service#near DataConsulService#near}.
	Near *string `field:"optional" json:"near" yaml:"near"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/service#node_meta DataConsulService#node_meta}.
	NodeMeta *map[string]*string `field:"optional" json:"nodeMeta" yaml:"nodeMeta"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/service#partition DataConsulService#partition}.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/service#require_consistent DataConsulService#require_consistent}.
	RequireConsistent interface{} `field:"optional" json:"requireConsistent" yaml:"requireConsistent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/service#token DataConsulService#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/service#wait_index DataConsulService#wait_index}.
	WaitIndex *float64 `field:"optional" json:"waitIndex" yaml:"waitIndex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/service#wait_time DataConsulService#wait_time}.
	WaitTime *string `field:"optional" json:"waitTime" yaml:"waitTime"`
}

