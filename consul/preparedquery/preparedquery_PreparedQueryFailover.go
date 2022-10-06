package preparedquery


type PreparedQueryFailover struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#datacenters PreparedQuery#datacenters}.
	Datacenters *[]*string `field:"optional" json:"datacenters" yaml:"datacenters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#nearest_n PreparedQuery#nearest_n}.
	NearestN *float64 `field:"optional" json:"nearestN" yaml:"nearestN"`
}

