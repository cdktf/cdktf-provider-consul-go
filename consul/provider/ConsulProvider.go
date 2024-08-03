// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-consul-go/consul/v9/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs consul}.
type ConsulProvider interface {
	cdktf.TerraformProvider
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AuthJwt() *ConsulProviderAuthJwt
	SetAuthJwt(val *ConsulProviderAuthJwt)
	AuthJwtInput() *ConsulProviderAuthJwt
	CaFile() *string
	SetCaFile(val *string)
	CaFileInput() *string
	CaPath() *string
	SetCaPath(val *string)
	CaPathInput() *string
	CaPem() *string
	SetCaPem(val *string)
	CaPemInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertFile() *string
	SetCertFile(val *string)
	CertFileInput() *string
	CertPem() *string
	SetCertPem(val *string)
	CertPemInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Datacenter() *string
	SetDatacenter(val *string)
	DatacenterInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Header() interface{}
	SetHeader(val interface{})
	HeaderInput() interface{}
	HttpAuth() *string
	SetHttpAuth(val *string)
	HttpAuthInput() *string
	InsecureHttps() interface{}
	SetInsecureHttps(val interface{})
	InsecureHttpsInput() interface{}
	KeyFile() *string
	SetKeyFile(val *string)
	KeyFileInput() *string
	KeyPem() *string
	SetKeyPem(val *string)
	KeyPemInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	Scheme() *string
	SetScheme(val *string)
	SchemeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAddress()
	ResetAlias()
	ResetAuthJwt()
	ResetCaFile()
	ResetCaPath()
	ResetCaPem()
	ResetCertFile()
	ResetCertPem()
	ResetDatacenter()
	ResetHeader()
	ResetHttpAuth()
	ResetInsecureHttps()
	ResetKeyFile()
	ResetKeyPem()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScheme()
	ResetToken()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ConsulProvider
type jsiiProxy_ConsulProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_ConsulProvider) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) AuthJwt() *ConsulProviderAuthJwt {
	var returns *ConsulProviderAuthJwt
	_jsii_.Get(
		j,
		"authJwt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) AuthJwtInput() *ConsulProviderAuthJwt {
	var returns *ConsulProviderAuthJwt
	_jsii_.Get(
		j,
		"authJwtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CaFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CaPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CaPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CaPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CaPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CertPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) CertPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) Datacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) DatacenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) Header() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) HttpAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) HttpAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) InsecureHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) InsecureHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) KeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) KeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) KeyPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) KeyPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) SchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs consul} Resource.
func NewConsulProvider(scope constructs.Construct, id *string, config *ConsulProviderConfig) ConsulProvider {
	_init_.Initialize()

	if err := validateNewConsulProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConsulProvider{}

	_jsii_.Create(
		"@cdktf/provider-consul.provider.ConsulProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs consul} Resource.
func NewConsulProvider_Override(c ConsulProvider, scope constructs.Construct, id *string, config *ConsulProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.provider.ConsulProvider",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConsulProvider)SetAddress(val *string) {
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetAuthJwt(val *ConsulProviderAuthJwt) {
	if err := j.validateSetAuthJwtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authJwt",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetCaFile(val *string) {
	_jsii_.Set(
		j,
		"caFile",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetCaPath(val *string) {
	_jsii_.Set(
		j,
		"caPath",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetCaPem(val *string) {
	_jsii_.Set(
		j,
		"caPem",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetCertFile(val *string) {
	_jsii_.Set(
		j,
		"certFile",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetCertPem(val *string) {
	_jsii_.Set(
		j,
		"certPem",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetDatacenter(val *string) {
	_jsii_.Set(
		j,
		"datacenter",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetHeader(val interface{}) {
	if err := j.validateSetHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"header",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetHttpAuth(val *string) {
	_jsii_.Set(
		j,
		"httpAuth",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetInsecureHttps(val interface{}) {
	if err := j.validateSetInsecureHttpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureHttps",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetKeyFile(val *string) {
	_jsii_.Set(
		j,
		"keyFile",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetKeyPem(val *string) {
	_jsii_.Set(
		j,
		"keyPem",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetScheme(val *string) {
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_ConsulProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

// Generates CDKTF code for importing a ConsulProvider resource upon running "cdktf plan <stack-name>".
func ConsulProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateConsulProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.provider.ConsulProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ConsulProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsulProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.provider.ConsulProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConsulProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsulProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.provider.ConsulProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConsulProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsulProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.provider.ConsulProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConsulProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-consul.provider.ConsulProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConsulProvider) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConsulProvider) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConsulProvider) ResetAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		c,
		"resetAlias",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetAuthJwt() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthJwt",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetCaFile() {
	_jsii_.InvokeVoid(
		c,
		"resetCaFile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetCaPath() {
	_jsii_.InvokeVoid(
		c,
		"resetCaPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetCaPem() {
	_jsii_.InvokeVoid(
		c,
		"resetCaPem",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetCertFile() {
	_jsii_.InvokeVoid(
		c,
		"resetCertFile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetCertPem() {
	_jsii_.InvokeVoid(
		c,
		"resetCertPem",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetDatacenter() {
	_jsii_.InvokeVoid(
		c,
		"resetDatacenter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetHttpAuth() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpAuth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetInsecureHttps() {
	_jsii_.InvokeVoid(
		c,
		"resetInsecureHttps",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetKeyFile() {
	_jsii_.InvokeVoid(
		c,
		"resetKeyFile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetKeyPem() {
	_jsii_.InvokeVoid(
		c,
		"resetKeyPem",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetScheme() {
	_jsii_.InvokeVoid(
		c,
		"resetScheme",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) ResetToken() {
	_jsii_.InvokeVoid(
		c,
		"resetToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

