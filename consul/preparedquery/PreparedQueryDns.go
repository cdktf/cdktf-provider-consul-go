package preparedquery


type PreparedQueryDns struct {
	// The TTL to send when returning DNS results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/prepared_query#ttl PreparedQuery#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

