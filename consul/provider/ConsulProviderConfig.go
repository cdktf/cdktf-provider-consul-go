// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type ConsulProviderConfig struct {
	// The HTTP(S) API address of the agent to use. Defaults to "127.0.0.1:8500".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#address ConsulProvider#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#alias ConsulProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// auth_jwt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#auth_jwt ConsulProvider#auth_jwt}
	AuthJwt *ConsulProviderAuthJwt `field:"optional" json:"authJwt" yaml:"authJwt"`
	// A path to a PEM-encoded certificate authority used to verify the remote agent's certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#ca_file ConsulProvider#ca_file}
	CaFile *string `field:"optional" json:"caFile" yaml:"caFile"`
	// A path to a directory of PEM-encoded certificate authority files to use to check the authenticity of client and server connections.
	//
	// Can also be specified with the `CONSUL_CAPATH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#ca_path ConsulProvider#ca_path}
	CaPath *string `field:"optional" json:"caPath" yaml:"caPath"`
	// PEM-encoded certificate authority used to verify the remote agent's certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#ca_pem ConsulProvider#ca_pem}
	CaPem *string `field:"optional" json:"caPem" yaml:"caPem"`
	// A path to a PEM-encoded certificate provided to the remote agent; requires use of `key_file` or `key_pem`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#cert_file ConsulProvider#cert_file}
	CertFile *string `field:"optional" json:"certFile" yaml:"certFile"`
	// PEM-encoded certificate provided to the remote agent; requires use of `key_file` or `key_pem`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#cert_pem ConsulProvider#cert_pem}
	CertPem *string `field:"optional" json:"certPem" yaml:"certPem"`
	// The datacenter to use. Defaults to that of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#datacenter ConsulProvider#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#header ConsulProvider#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// HTTP Basic Authentication credentials to be used when communicating with Consul, in the format of either `user` or `user:pass`.
	//
	// This may also be specified using the `CONSUL_HTTP_AUTH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#http_auth ConsulProvider#http_auth}
	HttpAuth *string `field:"optional" json:"httpAuth" yaml:"httpAuth"`
	// Boolean value to disable SSL certificate verification;
	//
	// setting this value to true is not recommended for production use. Only use this with scheme set to "https".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#insecure_https ConsulProvider#insecure_https}
	InsecureHttps interface{} `field:"optional" json:"insecureHttps" yaml:"insecureHttps"`
	// A path to a PEM-encoded private key, required if `cert_file` or `cert_pem` is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#key_file ConsulProvider#key_file}
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
	// PEM-encoded private key, required if `cert_file` or `cert_pem` is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#key_pem ConsulProvider#key_pem}
	KeyPem *string `field:"optional" json:"keyPem" yaml:"keyPem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#namespace ConsulProvider#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The URL scheme of the agent to use ("http" or "https"). Defaults to "http".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#scheme ConsulProvider#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// The ACL token to use by default when making requests to the agent.
	//
	// Can also be specified with `CONSUL_HTTP_TOKEN` or `CONSUL_TOKEN` as an environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs#token ConsulProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

