package preparedquery


type PreparedQueryDns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/prepared_query#ttl PreparedQuery#ttl}.
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

