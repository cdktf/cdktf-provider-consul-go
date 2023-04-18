package provider


type ConsulProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#address ConsulProvider#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#alias ConsulProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#ca_file ConsulProvider#ca_file}.
	CaFile *string `field:"optional" json:"caFile" yaml:"caFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#ca_path ConsulProvider#ca_path}.
	CaPath *string `field:"optional" json:"caPath" yaml:"caPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#ca_pem ConsulProvider#ca_pem}.
	CaPem *string `field:"optional" json:"caPem" yaml:"caPem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#cert_file ConsulProvider#cert_file}.
	CertFile *string `field:"optional" json:"certFile" yaml:"certFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#cert_pem ConsulProvider#cert_pem}.
	CertPem *string `field:"optional" json:"certPem" yaml:"certPem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#datacenter ConsulProvider#datacenter}.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#header ConsulProvider#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#http_auth ConsulProvider#http_auth}.
	HttpAuth *string `field:"optional" json:"httpAuth" yaml:"httpAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#insecure_https ConsulProvider#insecure_https}.
	InsecureHttps interface{} `field:"optional" json:"insecureHttps" yaml:"insecureHttps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#key_file ConsulProvider#key_file}.
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#key_pem ConsulProvider#key_pem}.
	KeyPem *string `field:"optional" json:"keyPem" yaml:"keyPem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#namespace ConsulProvider#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#scheme ConsulProvider#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs#token ConsulProvider#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

