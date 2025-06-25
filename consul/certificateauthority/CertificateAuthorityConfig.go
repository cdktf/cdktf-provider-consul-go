// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificateauthority

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertificateAuthorityConfig struct {
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
	// Specifies the CA provider type to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/certificate_authority#connect_provider CertificateAuthority#connect_provider}
	ConnectProvider *string `field:"required" json:"connectProvider" yaml:"connectProvider"`
	// The raw configuration to use for the chosen provider.
	//
	// For more information on configuring the Connect CA providers, see [Provider Config](https://developer.hashicorp.com/consul/docs/connect/ca).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/certificate_authority#config CertificateAuthority#config}
	Config *map[string]*string `field:"optional" json:"config" yaml:"config"`
	// The raw configuration to use for the chosen provider.
	//
	// For more information on configuring the Connect CA providers, see [Provider Config](https://developer.hashicorp.com/consul/docs/connect/ca).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/certificate_authority#config_json CertificateAuthority#config_json}
	ConfigJson *string `field:"optional" json:"configJson" yaml:"configJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/consul/2.22.0/docs/resources/certificate_authority#id CertificateAuthority#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

