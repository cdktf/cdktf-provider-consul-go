// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autopilotconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-consul-go/consul/v7/autopilotconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/autopilot_config consul_autopilot_config}.
type AutopilotConfig interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CleanupDeadServers() interface{}
	SetCleanupDeadServers(val interface{})
	CleanupDeadServersInput() interface{}
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
	Datacenter() *string
	SetDatacenter(val *string)
	DatacenterInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableUpgradeMigration() interface{}
	SetDisableUpgradeMigration(val interface{})
	DisableUpgradeMigrationInput() interface{}
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
	LastContactThreshold() *string
	SetLastContactThreshold(val *string)
	LastContactThresholdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxTrailingLogs() *float64
	SetMaxTrailingLogs(val *float64)
	MaxTrailingLogsInput() *float64
	// The tree node.
	Node() constructs.Node
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
	RedundancyZoneTag() *string
	SetRedundancyZoneTag(val *string)
	RedundancyZoneTagInput() *string
	ServerStabilizationTime() *string
	SetServerStabilizationTime(val *string)
	ServerStabilizationTimeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpgradeVersionTag() *string
	SetUpgradeVersionTag(val *string)
	UpgradeVersionTagInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetCleanupDeadServers()
	ResetDatacenter()
	ResetDisableUpgradeMigration()
	ResetId()
	ResetLastContactThreshold()
	ResetMaxTrailingLogs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRedundancyZoneTag()
	ResetServerStabilizationTime()
	ResetUpgradeVersionTag()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AutopilotConfig
type jsiiProxy_AutopilotConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AutopilotConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) CleanupDeadServers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupDeadServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) CleanupDeadServersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupDeadServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) Datacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) DatacenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) DisableUpgradeMigration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUpgradeMigration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) DisableUpgradeMigrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUpgradeMigrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) LastContactThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastContactThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) LastContactThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastContactThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) MaxTrailingLogs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTrailingLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) MaxTrailingLogsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTrailingLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) RedundancyZoneTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redundancyZoneTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) RedundancyZoneTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redundancyZoneTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) ServerStabilizationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverStabilizationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) ServerStabilizationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverStabilizationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) UpgradeVersionTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeVersionTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutopilotConfig) UpgradeVersionTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeVersionTagInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/autopilot_config consul_autopilot_config} Resource.
func NewAutopilotConfig(scope constructs.Construct, id *string, config *AutopilotConfigConfig) AutopilotConfig {
	_init_.Initialize()

	if err := validateNewAutopilotConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutopilotConfig{}

	_jsii_.Create(
		"@cdktf/provider-consul.autopilotConfig.AutopilotConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/autopilot_config consul_autopilot_config} Resource.
func NewAutopilotConfig_Override(a AutopilotConfig, scope constructs.Construct, id *string, config *AutopilotConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.autopilotConfig.AutopilotConfig",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetCleanupDeadServers(val interface{}) {
	if err := j.validateSetCleanupDeadServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanupDeadServers",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetDatacenter(val *string) {
	if err := j.validateSetDatacenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenter",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetDisableUpgradeMigration(val interface{}) {
	if err := j.validateSetDisableUpgradeMigrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUpgradeMigration",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetLastContactThreshold(val *string) {
	if err := j.validateSetLastContactThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastContactThreshold",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetMaxTrailingLogs(val *float64) {
	if err := j.validateSetMaxTrailingLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTrailingLogs",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetRedundancyZoneTag(val *string) {
	if err := j.validateSetRedundancyZoneTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redundancyZoneTag",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetServerStabilizationTime(val *string) {
	if err := j.validateSetServerStabilizationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverStabilizationTime",
		val,
	)
}

func (j *jsiiProxy_AutopilotConfig)SetUpgradeVersionTag(val *string) {
	if err := j.validateSetUpgradeVersionTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeVersionTag",
		val,
	)
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
func AutopilotConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutopilotConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.autopilotConfig.AutopilotConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutopilotConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutopilotConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.autopilotConfig.AutopilotConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutopilotConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutopilotConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.autopilotConfig.AutopilotConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AutopilotConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-consul.autopilotConfig.AutopilotConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutopilotConfig) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutopilotConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutopilotConfig) ResetCleanupDeadServers() {
	_jsii_.InvokeVoid(
		a,
		"resetCleanupDeadServers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutopilotConfig) ResetDatacenter() {
	_jsii_.InvokeVoid(
		a,
		"resetDatacenter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutopilotConfig) ResetDisableUpgradeMigration() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableUpgradeMigration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutopilotConfig) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutopilotConfig) ResetLastContactThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetLastContactThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutopilotConfig) ResetMaxTrailingLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTrailingLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutopilotConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutopilotConfig) ResetRedundancyZoneTag() {
	_jsii_.InvokeVoid(
		a,
		"resetRedundancyZoneTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutopilotConfig) ResetServerStabilizationTime() {
	_jsii_.InvokeVoid(
		a,
		"resetServerStabilizationTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutopilotConfig) ResetUpgradeVersionTag() {
	_jsii_.InvokeVoid(
		a,
		"resetUpgradeVersionTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutopilotConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutopilotConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

