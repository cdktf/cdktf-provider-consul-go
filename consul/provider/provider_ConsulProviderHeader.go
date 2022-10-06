package provider


type ConsulProviderHeader struct {
	// The header name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul#name ConsulProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul#value ConsulProvider#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

