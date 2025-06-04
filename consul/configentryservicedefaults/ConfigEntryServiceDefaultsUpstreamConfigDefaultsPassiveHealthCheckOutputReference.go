// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v10/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v10/configentryservicedefaults/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference interface {
	cdktf.ComplexObject
	BaseEjectionTime() *string
	SetBaseEjectionTime(val *string)
	BaseEjectionTimeInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnforcingConsecutive5Xx() *float64
	SetEnforcingConsecutive5Xx(val *float64)
	EnforcingConsecutive5XxInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	MaxEjectionPercent() *float64
	SetMaxEjectionPercent(val *float64)
	MaxEjectionPercentInput() *float64
	MaxFailures() *float64
	SetMaxFailures(val *float64)
	MaxFailuresInput() *float64
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
	ResetBaseEjectionTime()
	ResetEnforcingConsecutive5Xx()
	ResetInterval()
	ResetMaxEjectionPercent()
	ResetMaxFailures()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference
type jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) BaseEjectionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseEjectionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) BaseEjectionTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseEjectionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) EnforcingConsecutive5Xx() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutive5Xx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) EnforcingConsecutive5XxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutive5XxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) MaxEjectionPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) MaxEjectionPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) MaxFailures() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) MaxFailuresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference_Override(c ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference)SetBaseEjectionTime(val *string) {
	if err := j.validateSetBaseEjectionTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseEjectionTime",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference)SetEnforcingConsecutive5Xx(val *float64) {
	if err := j.validateSetEnforcingConsecutive5XxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcingConsecutive5Xx",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference)SetInterval(val *string) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference)SetMaxEjectionPercent(val *float64) {
	if err := j.validateSetMaxEjectionPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxEjectionPercent",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference)SetMaxFailures(val *float64) {
	if err := j.validateSetMaxFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFailures",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) ResetBaseEjectionTime() {
	_jsii_.InvokeVoid(
		c,
		"resetBaseEjectionTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) ResetEnforcingConsecutive5Xx() {
	_jsii_.InvokeVoid(
		c,
		"resetEnforcingConsecutive5Xx",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) ResetMaxEjectionPercent() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxEjectionPercent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) ResetMaxFailures() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxFailures",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigDefaultsPassiveHealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

