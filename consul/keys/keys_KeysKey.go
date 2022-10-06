package keys


type KeysKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/keys#path Keys#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/keys#default Keys#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/keys#delete Keys#delete}.
	Delete interface{} `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/keys#flags Keys#flags}.
	Flags *float64 `field:"optional" json:"flags" yaml:"flags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/keys#name Keys#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/keys#value Keys#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

