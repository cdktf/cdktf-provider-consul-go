package service

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#name Service#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#node Service#node}.
	NodeAttribute *string `field:"required" json:"nodeAttribute" yaml:"nodeAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#address Service#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// check block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#check Service#check}
	Check interface{} `field:"optional" json:"check" yaml:"check"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#datacenter Service#datacenter}.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#enable_tag_override Service#enable_tag_override}.
	EnableTagOverride interface{} `field:"optional" json:"enableTagOverride" yaml:"enableTagOverride"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#external Service#external}.
	External interface{} `field:"optional" json:"external" yaml:"external"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#id Service#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#meta Service#meta}.
	Meta *map[string]*string `field:"optional" json:"meta" yaml:"meta"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#namespace Service#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The partition the service is associated with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#partition Service#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#port Service#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#service_id Service#service_id}.
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/service#tags Service#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

