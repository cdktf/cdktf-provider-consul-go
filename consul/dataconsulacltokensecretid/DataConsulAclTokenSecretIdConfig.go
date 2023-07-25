package dataconsulacltokensecretid

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataConsulAclTokenSecretIdConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/acl_token_secret_id#accessor_id DataConsulAclTokenSecretId#accessor_id}.
	AccessorId *string `field:"required" json:"accessorId" yaml:"accessorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/acl_token_secret_id#id DataConsulAclTokenSecretId#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/acl_token_secret_id#namespace DataConsulAclTokenSecretId#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/acl_token_secret_id#partition DataConsulAclTokenSecretId#partition}.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/data-sources/acl_token_secret_id#pgp_key DataConsulAclTokenSecretId#pgp_key}.
	PgpKey *string `field:"optional" json:"pgpKey" yaml:"pgpKey"`
}

