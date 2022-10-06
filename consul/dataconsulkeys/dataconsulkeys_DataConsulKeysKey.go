package dataconsulkeys


type DataConsulKeysKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/keys#name DataConsulKeys#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/keys#path DataConsulKeys#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/keys#default DataConsulKeys#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
}

