package preparedquery


type PreparedQueryFailoverTargets struct {
	// Specifies a WAN federated datacenter to forward the query to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/prepared_query#datacenter PreparedQuery#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Specifies a cluster peer to use for failover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/prepared_query#peer PreparedQuery#peer}
	Peer *string `field:"optional" json:"peer" yaml:"peer"`
}

