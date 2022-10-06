package dataconsulservices


type DataConsulServicesQueryOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/services#allow_stale DataConsulServices#allow_stale}.
	AllowStale interface{} `field:"optional" json:"allowStale" yaml:"allowStale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/services#datacenter DataConsulServices#datacenter}.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/services#namespace DataConsulServices#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/services#near DataConsulServices#near}.
	Near *string `field:"optional" json:"near" yaml:"near"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/services#node_meta DataConsulServices#node_meta}.
	NodeMeta *map[string]*string `field:"optional" json:"nodeMeta" yaml:"nodeMeta"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/services#partition DataConsulServices#partition}.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/services#require_consistent DataConsulServices#require_consistent}.
	RequireConsistent interface{} `field:"optional" json:"requireConsistent" yaml:"requireConsistent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/services#token DataConsulServices#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/services#wait_index DataConsulServices#wait_index}.
	WaitIndex *float64 `field:"optional" json:"waitIndex" yaml:"waitIndex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/services#wait_time DataConsulServices#wait_time}.
	WaitTime *string `field:"optional" json:"waitTime" yaml:"waitTime"`
}

