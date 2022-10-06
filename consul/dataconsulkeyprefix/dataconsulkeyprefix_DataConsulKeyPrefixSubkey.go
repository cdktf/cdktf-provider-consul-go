package dataconsulkeyprefix


type DataConsulKeyPrefixSubkey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/key_prefix#name DataConsulKeyPrefix#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/key_prefix#path DataConsulKeyPrefix#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/key_prefix#default DataConsulKeyPrefix#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
}

