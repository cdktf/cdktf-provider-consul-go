package aclrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AclRoleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The name of the ACL role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#name AclRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A free form human readable description of the role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#description AclRole#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#id AclRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#namespace AclRole#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// node_identities block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#node_identities AclRole#node_identities}
	NodeIdentities interface{} `field:"optional" json:"nodeIdentities" yaml:"nodeIdentities"`
	// The partition the ACL role is associated with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#partition AclRole#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// The list of policies that should be applied to the role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#policies AclRole#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// service_identities block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#service_identities AclRole#service_identities}
	ServiceIdentities interface{} `field:"optional" json:"serviceIdentities" yaml:"serviceIdentities"`
}

