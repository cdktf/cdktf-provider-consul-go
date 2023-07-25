package keyprefix


type KeyPrefixSubkey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/key_prefix#path KeyPrefix#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/key_prefix#value KeyPrefix#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/key_prefix#flags KeyPrefix#flags}.
	Flags *float64 `field:"optional" json:"flags" yaml:"flags"`
}

