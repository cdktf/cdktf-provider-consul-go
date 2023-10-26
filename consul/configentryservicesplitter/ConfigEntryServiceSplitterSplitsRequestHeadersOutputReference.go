// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicesplitter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v8/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v8/configentryservicesplitter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference interface {
	cdktf.ComplexObject
	Add() *map[string]*string
	SetAdd(val *map[string]*string)
	AddInput() *map[string]*string
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
	// Experimental.
	Fqn() *string
	InternalValue() *ConfigEntryServiceSplitterSplitsRequestHeaders
	SetInternalValue(val *ConfigEntryServiceSplitterSplitsRequestHeaders)
	Remove() *[]*string
	SetRemove(val *[]*string)
	RemoveInput() *[]*string
	Set() *map[string]*string
	SetSet(val *map[string]*string)
	SetInput() *map[string]*string
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
	ResetAdd()
	ResetRemove()
	ResetSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference
type jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) Add() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"add",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) AddInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"addInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) InternalValue() *ConfigEntryServiceSplitterSplitsRequestHeaders {
	var returns *ConfigEntryServiceSplitterSplitsRequestHeaders
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) Remove() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) RemoveInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"removeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) Set() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"set",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) SetInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"setInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceSplitterSplitsRequestHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceSplitterSplitsRequestHeadersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceSplitter.ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigEntryServiceSplitterSplitsRequestHeadersOutputReference_Override(c ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceSplitter.ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference)SetAdd(val *map[string]*string) {
	if err := j.validateSetAddParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"add",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference)SetInternalValue(val *ConfigEntryServiceSplitterSplitsRequestHeaders) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference)SetRemove(val *[]*string) {
	if err := j.validateSetRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remove",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference)SetSet(val *map[string]*string) {
	if err := j.validateSetSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"set",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) ResetAdd() {
	_jsii_.InvokeVoid(
		c,
		"resetAdd",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) ResetRemove() {
	_jsii_.InvokeVoid(
		c,
		"resetRemove",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) ResetSet() {
	_jsii_.InvokeVoid(
		c,
		"resetSet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceSplitterSplitsRequestHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

