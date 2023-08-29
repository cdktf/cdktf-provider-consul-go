// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-consul.provider.ConsulProvider",
		reflect.TypeOf((*ConsulProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "address", GoGetter: "Address"},
			_jsii_.MemberProperty{JsiiProperty: "addressInput", GoGetter: "AddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "authJwt", GoGetter: "AuthJwt"},
			_jsii_.MemberProperty{JsiiProperty: "authJwtInput", GoGetter: "AuthJwtInput"},
			_jsii_.MemberProperty{JsiiProperty: "caFile", GoGetter: "CaFile"},
			_jsii_.MemberProperty{JsiiProperty: "caFileInput", GoGetter: "CaFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "caPath", GoGetter: "CaPath"},
			_jsii_.MemberProperty{JsiiProperty: "caPathInput", GoGetter: "CaPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "caPem", GoGetter: "CaPem"},
			_jsii_.MemberProperty{JsiiProperty: "caPemInput", GoGetter: "CaPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certFile", GoGetter: "CertFile"},
			_jsii_.MemberProperty{JsiiProperty: "certFileInput", GoGetter: "CertFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "certPem", GoGetter: "CertPem"},
			_jsii_.MemberProperty{JsiiProperty: "certPemInput", GoGetter: "CertPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "datacenter", GoGetter: "Datacenter"},
			_jsii_.MemberProperty{JsiiProperty: "datacenterInput", GoGetter: "DatacenterInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "header", GoGetter: "Header"},
			_jsii_.MemberProperty{JsiiProperty: "headerInput", GoGetter: "HeaderInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpAuth", GoGetter: "HttpAuth"},
			_jsii_.MemberProperty{JsiiProperty: "httpAuthInput", GoGetter: "HttpAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecureHttps", GoGetter: "InsecureHttps"},
			_jsii_.MemberProperty{JsiiProperty: "insecureHttpsInput", GoGetter: "InsecureHttpsInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyFile", GoGetter: "KeyFile"},
			_jsii_.MemberProperty{JsiiProperty: "keyFileInput", GoGetter: "KeyFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyPem", GoGetter: "KeyPem"},
			_jsii_.MemberProperty{JsiiProperty: "keyPemInput", GoGetter: "KeyPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAddress", GoMethod: "ResetAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthJwt", GoMethod: "ResetAuthJwt"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaFile", GoMethod: "ResetCaFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaPath", GoMethod: "ResetCaPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaPem", GoMethod: "ResetCaPem"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertFile", GoMethod: "ResetCertFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertPem", GoMethod: "ResetCertPem"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatacenter", GoMethod: "ResetDatacenter"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeader", GoMethod: "ResetHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpAuth", GoMethod: "ResetHttpAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecureHttps", GoMethod: "ResetInsecureHttps"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyFile", GoMethod: "ResetKeyFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyPem", GoMethod: "ResetKeyPem"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheme", GoMethod: "ResetScheme"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberProperty{JsiiProperty: "scheme", GoGetter: "Scheme"},
			_jsii_.MemberProperty{JsiiProperty: "schemeInput", GoGetter: "SchemeInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ConsulProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-consul.provider.ConsulProviderAuthJwt",
		reflect.TypeOf((*ConsulProviderAuthJwt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-consul.provider.ConsulProviderConfig",
		reflect.TypeOf((*ConsulProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-consul.provider.ConsulProviderHeader",
		reflect.TypeOf((*ConsulProviderHeader)(nil)).Elem(),
	)
}
