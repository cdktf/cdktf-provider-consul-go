package service


type ServiceCheckHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/service#name Service#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/service#value Service#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

