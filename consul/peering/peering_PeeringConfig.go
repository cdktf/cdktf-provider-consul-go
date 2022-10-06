package peering

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PeeringConfig struct {
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
	// The peering token fetched from the peer cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/peering#peering_token Peering#peering_token}
	PeeringToken *string `field:"required" json:"peeringToken" yaml:"peeringToken"`
	// The name assigned to the peer cluster.
	//
	// The `peer_name` is used to reference the peer cluster in service discovery queries and configuration entries such as `service-intentions`. This field must be a valid DNS hostname label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/peering#peer_name Peering#peer_name}
	PeerName *string `field:"required" json:"peerName" yaml:"peerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/peering#id Peering#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies KV metadata to associate with the peering.
	//
	// This parameter is not required and does not directly impact the cluster peering process.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/peering#meta Peering#meta}
	Meta *map[string]*string `field:"optional" json:"meta" yaml:"meta"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/peering#partition Peering#partition}.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
}

