// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v10/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v10/configentryservicedefaults/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference interface {
	cdktf.ComplexObject
	BalanceOutboundConnections() *string
	SetBalanceOutboundConnections(val *string)
	BalanceOutboundConnectionsInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectTimeoutMs() *float64
	SetConnectTimeoutMs(val *float64)
	ConnectTimeoutMsInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Limits() ConfigEntryServiceDefaultsUpstreamConfigDefaultsLimitsList
	LimitsInput() interface{}
	MeshGateway() ConfigEntryServiceDefaultsUpstreamConfigDefaultsMeshGatewayList
	MeshGatewayInput() interface{}
	PassiveHealthCheck() ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckList
	PassiveHealthCheckInput() interface{}
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutLimits(value interface{})
	PutMeshGateway(value interface{})
	PutPassiveHealthCheck(value interface{})
	ResetBalanceOutboundConnections()
	ResetConnectTimeoutMs()
	ResetLimits()
	ResetMeshGateway()
	ResetPassiveHealthCheck()
	ResetProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference
type jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) BalanceOutboundConnections() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balanceOutboundConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) BalanceOutboundConnectionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balanceOutboundConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ConnectTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ConnectTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) Limits() ConfigEntryServiceDefaultsUpstreamConfigDefaultsLimitsList {
	var returns ConfigEntryServiceDefaultsUpstreamConfigDefaultsLimitsList
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) LimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) MeshGateway() ConfigEntryServiceDefaultsUpstreamConfigDefaultsMeshGatewayList {
	var returns ConfigEntryServiceDefaultsUpstreamConfigDefaultsMeshGatewayList
	_jsii_.Get(
		j,
		"meshGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) MeshGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"meshGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) PassiveHealthCheck() ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckList {
	var returns ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckList
	_jsii_.Get(
		j,
		"passiveHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) PassiveHealthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passiveHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference_Override(c ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference)SetBalanceOutboundConnections(val *string) {
	if err := j.validateSetBalanceOutboundConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"balanceOutboundConnections",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference)SetConnectTimeoutMs(val *float64) {
	if err := j.validateSetConnectTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeoutMs",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) PutLimits(value interface{}) {
	if err := c.validatePutLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLimits",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) PutMeshGateway(value interface{}) {
	if err := c.validatePutMeshGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMeshGateway",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) PutPassiveHealthCheck(value interface{}) {
	if err := c.validatePutPassiveHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPassiveHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ResetBalanceOutboundConnections() {
	_jsii_.InvokeVoid(
		c,
		"resetBalanceOutboundConnections",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ResetConnectTimeoutMs() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectTimeoutMs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ResetLimits() {
	_jsii_.InvokeVoid(
		c,
		"resetLimits",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ResetMeshGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetMeshGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ResetPassiveHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetPassiveHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

