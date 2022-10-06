package preparedquery


type PreparedQueryTemplate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#regexp PreparedQuery#regexp}.
	Regexp *string `field:"required" json:"regexp" yaml:"regexp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/prepared_query#type PreparedQuery#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

