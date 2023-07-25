package aclbindingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclBindingRuleConfig struct {
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
	// The name of the ACL auth method this rule apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/acl_binding_rule#auth_method AclBindingRule#auth_method}
	AuthMethod *string `field:"required" json:"authMethod" yaml:"authMethod"`
	// The name to bind to a token at login-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/acl_binding_rule#bind_name AclBindingRule#bind_name}
	BindName *string `field:"required" json:"bindName" yaml:"bindName"`
	// Specifies the way the binding rule affects a token created at login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/acl_binding_rule#bind_type AclBindingRule#bind_type}
	BindType *string `field:"required" json:"bindType" yaml:"bindType"`
	// A free form human readable description of the binding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/acl_binding_rule#description AclBindingRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/acl_binding_rule#id AclBindingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/acl_binding_rule#namespace AclBindingRule#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The partition the ACL binding rule is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/acl_binding_rule#partition AclBindingRule#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// The expression used to math this rule against valid identities returned from an auth method validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/acl_binding_rule#selector AclBindingRule#selector}
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
}

