package provider


type ConsulProviderHeader struct {
	// The header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#name ConsulProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#value ConsulProvider#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

