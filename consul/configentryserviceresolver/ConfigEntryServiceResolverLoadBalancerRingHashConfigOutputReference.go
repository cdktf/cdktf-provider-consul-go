// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v10/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v10/configentryserviceresolver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaximumRingSize() *float64
	SetMaximumRingSize(val *float64)
	MaximumRingSizeInput() *float64
	MinimumRingSize() *float64
	SetMinimumRingSize(val *float64)
	MinimumRingSizeInput() *float64
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
	ResetMaximumRingSize()
	ResetMinimumRingSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference
type jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) MaximumRingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) MaximumRingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) MinimumRingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumRingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) MinimumRingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumRingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceResolver.ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference_Override(c ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceResolver.ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference)SetMaximumRingSize(val *float64) {
	if err := j.validateSetMaximumRingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRingSize",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference)SetMinimumRingSize(val *float64) {
	if err := j.validateSetMinimumRingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumRingSize",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) ResetMaximumRingSize() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumRingSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) ResetMinimumRingSize() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumRingSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerRingHashConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

