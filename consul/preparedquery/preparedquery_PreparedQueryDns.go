package preparedquery


type PreparedQueryDns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#ttl PreparedQuery#ttl}.
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

