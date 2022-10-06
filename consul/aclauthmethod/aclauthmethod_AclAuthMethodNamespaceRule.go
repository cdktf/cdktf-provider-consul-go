package aclauthmethod


type AclAuthMethodNamespaceRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_auth_method#bind_namespace AclAuthMethod#bind_namespace}.
	BindNamespace *string `field:"required" json:"bindNamespace" yaml:"bindNamespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_auth_method#selector AclAuthMethod#selector}.
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
}

