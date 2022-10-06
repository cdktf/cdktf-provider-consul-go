package aclrole


type AclRoleServiceIdentities struct {
	// The name of the service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#service_name AclRole#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The datacenters the effective policy is valid within.
	//
	// When no datacenters are provided the effective policy is valid in all datacenters including those which do not yet exist but may in the future.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/acl_role#datacenters AclRole#datacenters}
	Datacenters *[]*string `field:"optional" json:"datacenters" yaml:"datacenters"`
}

