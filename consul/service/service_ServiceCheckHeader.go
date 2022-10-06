package service


type ServiceCheckHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#name Service#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#value Service#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

