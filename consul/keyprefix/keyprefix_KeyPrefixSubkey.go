package keyprefix


type KeyPrefixSubkey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/key_prefix#path KeyPrefix#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/key_prefix#value KeyPrefix#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/key_prefix#flags KeyPrefix#flags}.
	Flags *float64 `field:"optional" json:"flags" yaml:"flags"`
}

