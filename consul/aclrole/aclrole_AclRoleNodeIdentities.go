package aclrole


type AclRoleNodeIdentities struct {
	// Specifies the node's datacenter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#datacenter AclRole#datacenter}
	Datacenter *string `field:"required" json:"datacenter" yaml:"datacenter"`
	// The name of the node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#node_name AclRole#node_name}
	NodeName *string `field:"required" json:"nodeName" yaml:"nodeName"`
}

