package dataconsulconfigentry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataConsulConfigEntryConfig struct {
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
	// The kind of config entry to read.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/config_entry#kind DataConsulConfigEntry#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The name of the config entry to read.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/config_entry#name DataConsulConfigEntry#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/config_entry#id DataConsulConfigEntry#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The namespace the config entry is associated with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/config_entry#namespace DataConsulConfigEntry#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The partition the config entry is associated with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/d/config_entry#partition DataConsulConfigEntry#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
}

