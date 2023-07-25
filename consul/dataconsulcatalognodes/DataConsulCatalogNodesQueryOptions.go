package dataconsulcatalognodes


type DataConsulCatalogNodesQueryOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/catalog_nodes#allow_stale DataConsulCatalogNodes#allow_stale}.
	AllowStale interface{} `field:"optional" json:"allowStale" yaml:"allowStale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/catalog_nodes#datacenter DataConsulCatalogNodes#datacenter}.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/catalog_nodes#near DataConsulCatalogNodes#near}.
	Near *string `field:"optional" json:"near" yaml:"near"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/catalog_nodes#node_meta DataConsulCatalogNodes#node_meta}.
	NodeMeta *map[string]*string `field:"optional" json:"nodeMeta" yaml:"nodeMeta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/catalog_nodes#partition DataConsulCatalogNodes#partition}.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/catalog_nodes#require_consistent DataConsulCatalogNodes#require_consistent}.
	RequireConsistent interface{} `field:"optional" json:"requireConsistent" yaml:"requireConsistent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/catalog_nodes#token DataConsulCatalogNodes#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/catalog_nodes#wait_index DataConsulCatalogNodes#wait_index}.
	WaitIndex *float64 `field:"optional" json:"waitIndex" yaml:"waitIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/catalog_nodes#wait_time DataConsulCatalogNodes#wait_time}.
	WaitTime *string `field:"optional" json:"waitTime" yaml:"waitTime"`
}

