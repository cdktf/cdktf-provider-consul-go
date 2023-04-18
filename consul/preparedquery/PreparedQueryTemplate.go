package preparedquery


type PreparedQueryTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/prepared_query#regexp PreparedQuery#regexp}.
	Regexp *string `field:"required" json:"regexp" yaml:"regexp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/prepared_query#type PreparedQuery#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

