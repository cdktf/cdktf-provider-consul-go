package provider


type ConsulProviderHeader struct {
	// The name of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#name ConsulProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#value ConsulProvider#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

