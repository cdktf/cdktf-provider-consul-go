package service


type ServiceCheck struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#check_id Service#check_id}.
	CheckId *string `field:"required" json:"checkId" yaml:"checkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#interval Service#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#name Service#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#timeout Service#timeout}.
	Timeout *string `field:"required" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#deregister_critical_service_after Service#deregister_critical_service_after}.
	DeregisterCriticalServiceAfter *string `field:"optional" json:"deregisterCriticalServiceAfter" yaml:"deregisterCriticalServiceAfter"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#header Service#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#http Service#http}.
	Http *string `field:"optional" json:"http" yaml:"http"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#method Service#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#notes Service#notes}.
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#status Service#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#tcp Service#tcp}.
	Tcp *string `field:"optional" json:"tcp" yaml:"tcp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#tls_skip_verify Service#tls_skip_verify}.
	TlsSkipVerify interface{} `field:"optional" json:"tlsSkipVerify" yaml:"tlsSkipVerify"`
}

