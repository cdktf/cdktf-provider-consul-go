// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-consul-go/consul/v10/configentryservicedefaults/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults consul_config_entry_service_defaults}.
type ConfigEntryServiceDefaults interface {
	cdktf.TerraformResource
	BalanceInboundConnections() *string
	SetBalanceInboundConnections(val *string)
	BalanceInboundConnectionsInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Destination() ConfigEntryServiceDefaultsDestinationList
	DestinationInput() interface{}
	EnvoyExtensions() ConfigEntryServiceDefaultsEnvoyExtensionsList
	EnvoyExtensionsInput() interface{}
	Expose() ConfigEntryServiceDefaultsExposeList
	ExposeInput() interface{}
	ExternalSni() *string
	SetExternalSni(val *string)
	ExternalSniInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalConnectTimeoutMs() *float64
	SetLocalConnectTimeoutMs(val *float64)
	LocalConnectTimeoutMsInput() *float64
	LocalRequestTimeoutMs() *float64
	SetLocalRequestTimeoutMs(val *float64)
	LocalRequestTimeoutMsInput() *float64
	MaxInboundConnections() *float64
	SetMaxInboundConnections(val *float64)
	MaxInboundConnectionsInput() *float64
	MeshGateway() ConfigEntryServiceDefaultsMeshGatewayList
	MeshGatewayInput() interface{}
	Meta() *map[string]*string
	SetMeta(val *map[string]*string)
	MetaInput() *map[string]*string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	MutualTlsMode() *string
	SetMutualTlsMode(val *string)
	MutualTlsModeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Partition() *string
	SetPartition(val *string)
	PartitionInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TransparentProxy() ConfigEntryServiceDefaultsTransparentProxyList
	TransparentProxyInput() interface{}
	UpstreamConfig() ConfigEntryServiceDefaultsUpstreamConfigList
	UpstreamConfigInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDestination(value interface{})
	PutEnvoyExtensions(value interface{})
	PutExpose(value interface{})
	PutMeshGateway(value interface{})
	PutTransparentProxy(value interface{})
	PutUpstreamConfig(value interface{})
	ResetBalanceInboundConnections()
	ResetDestination()
	ResetEnvoyExtensions()
	ResetExternalSni()
	ResetId()
	ResetLocalConnectTimeoutMs()
	ResetLocalRequestTimeoutMs()
	ResetMaxInboundConnections()
	ResetMeshGateway()
	ResetMeta()
	ResetMode()
	ResetMutualTlsMode()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartition()
	ResetTransparentProxy()
	ResetUpstreamConfig()
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

// The jsii proxy struct for ConfigEntryServiceDefaults
type jsiiProxy_ConfigEntryServiceDefaults struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) BalanceInboundConnections() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balanceInboundConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) BalanceInboundConnectionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balanceInboundConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Destination() ConfigEntryServiceDefaultsDestinationList {
	var returns ConfigEntryServiceDefaultsDestinationList
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) DestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) EnvoyExtensions() ConfigEntryServiceDefaultsEnvoyExtensionsList {
	var returns ConfigEntryServiceDefaultsEnvoyExtensionsList
	_jsii_.Get(
		j,
		"envoyExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) EnvoyExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envoyExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Expose() ConfigEntryServiceDefaultsExposeList {
	var returns ConfigEntryServiceDefaultsExposeList
	_jsii_.Get(
		j,
		"expose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) ExposeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exposeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) ExternalSni() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalSni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) ExternalSniInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalSniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) LocalConnectTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localConnectTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) LocalConnectTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localConnectTimeoutMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) LocalRequestTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localRequestTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) LocalRequestTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localRequestTimeoutMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) MaxInboundConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInboundConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) MaxInboundConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInboundConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) MeshGateway() ConfigEntryServiceDefaultsMeshGatewayList {
	var returns ConfigEntryServiceDefaultsMeshGatewayList
	_jsii_.Get(
		j,
		"meshGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) MeshGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"meshGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Meta() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"meta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) MetaInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) MutualTlsMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mutualTlsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) MutualTlsModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mutualTlsModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) PartitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) TransparentProxy() ConfigEntryServiceDefaultsTransparentProxyList {
	var returns ConfigEntryServiceDefaultsTransparentProxyList
	_jsii_.Get(
		j,
		"transparentProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) TransparentProxyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transparentProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) UpstreamConfig() ConfigEntryServiceDefaultsUpstreamConfigList {
	var returns ConfigEntryServiceDefaultsUpstreamConfigList
	_jsii_.Get(
		j,
		"upstreamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaults) UpstreamConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upstreamConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults consul_config_entry_service_defaults} Resource.
func NewConfigEntryServiceDefaults(scope constructs.Construct, id *string, config *ConfigEntryServiceDefaultsConfig) ConfigEntryServiceDefaults {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceDefaultsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceDefaults{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaults",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/consul/2.21.0/docs/resources/config_entry_service_defaults consul_config_entry_service_defaults} Resource.
func NewConfigEntryServiceDefaults_Override(c ConfigEntryServiceDefaults, scope constructs.Construct, id *string, config *ConfigEntryServiceDefaultsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaults",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetBalanceInboundConnections(val *string) {
	if err := j.validateSetBalanceInboundConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"balanceInboundConnections",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetExternalSni(val *string) {
	if err := j.validateSetExternalSniParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalSni",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetLocalConnectTimeoutMs(val *float64) {
	if err := j.validateSetLocalConnectTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localConnectTimeoutMs",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetLocalRequestTimeoutMs(val *float64) {
	if err := j.validateSetLocalRequestTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localRequestTimeoutMs",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetMaxInboundConnections(val *float64) {
	if err := j.validateSetMaxInboundConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInboundConnections",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetMeta(val *map[string]*string) {
	if err := j.validateSetMetaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meta",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetMutualTlsMode(val *string) {
	if err := j.validateSetMutualTlsModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mutualTlsMode",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetPartition(val *string) {
	if err := j.validateSetPartitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partition",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaults)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ConfigEntryServiceDefaults resource upon running "cdktf plan <stack-name>".
func ConfigEntryServiceDefaults_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateConfigEntryServiceDefaults_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaults",
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
func ConfigEntryServiceDefaults_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigEntryServiceDefaults_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaults",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConfigEntryServiceDefaults_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigEntryServiceDefaults_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaults",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConfigEntryServiceDefaults_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigEntryServiceDefaults_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaults",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConfigEntryServiceDefaults_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaults",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) PutDestination(value interface{}) {
	if err := c.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDestination",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) PutEnvoyExtensions(value interface{}) {
	if err := c.validatePutEnvoyExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnvoyExtensions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) PutExpose(value interface{}) {
	if err := c.validatePutExposeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExpose",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) PutMeshGateway(value interface{}) {
	if err := c.validatePutMeshGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMeshGateway",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) PutTransparentProxy(value interface{}) {
	if err := c.validatePutTransparentProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTransparentProxy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) PutUpstreamConfig(value interface{}) {
	if err := c.validatePutUpstreamConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUpstreamConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetBalanceInboundConnections() {
	_jsii_.InvokeVoid(
		c,
		"resetBalanceInboundConnections",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetDestination() {
	_jsii_.InvokeVoid(
		c,
		"resetDestination",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetEnvoyExtensions() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvoyExtensions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetExternalSni() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalSni",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetLocalConnectTimeoutMs() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalConnectTimeoutMs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetLocalRequestTimeoutMs() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalRequestTimeoutMs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetMaxInboundConnections() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxInboundConnections",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetMeshGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetMeshGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetMeta() {
	_jsii_.InvokeVoid(
		c,
		"resetMeta",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetMode() {
	_jsii_.InvokeVoid(
		c,
		"resetMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetMutualTlsMode() {
	_jsii_.InvokeVoid(
		c,
		"resetMutualTlsMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetPartition() {
	_jsii_.InvokeVoid(
		c,
		"resetPartition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetTransparentProxy() {
	_jsii_.InvokeVoid(
		c,
		"resetTransparentProxy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ResetUpstreamConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetUpstreamConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaults) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

