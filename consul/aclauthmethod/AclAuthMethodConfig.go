package aclauthmethod

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclAuthMethodConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the ACL auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#name AclAuthMethod#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the ACL auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#type AclAuthMethod#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The raw configuration for this ACL auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#config AclAuthMethod#config}
	Config *map[string]*string `field:"optional" json:"config" yaml:"config"`
	// The raw configuration for this ACL auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#config_json AclAuthMethod#config_json}
	ConfigJson *string `field:"optional" json:"configJson" yaml:"configJson"`
	// A free form human readable description of the auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#description AclAuthMethod#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An optional name to use instead of the name attribute when displaying information about this auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#display_name AclAuthMethod#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#id AclAuthMethod#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The maximum life of any token created by this auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#max_token_ttl AclAuthMethod#max_token_ttl}
	MaxTokenTtl *string `field:"optional" json:"maxTokenTtl" yaml:"maxTokenTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#namespace AclAuthMethod#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// namespace_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#namespace_rule AclAuthMethod#namespace_rule}
	NamespaceRule interface{} `field:"optional" json:"namespaceRule" yaml:"namespaceRule"`
	// The partition the ACL auth method is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#partition AclAuthMethod#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// The kind of token that this auth method produces. This can be either 'local' or 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/resources/acl_auth_method#token_locality AclAuthMethod#token_locality}
	TokenLocality *string `field:"optional" json:"tokenLocality" yaml:"tokenLocality"`
}

