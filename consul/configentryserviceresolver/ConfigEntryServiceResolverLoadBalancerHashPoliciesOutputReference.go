// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v10/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v10/configentryserviceresolver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference interface {
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
	CookieConfig() ConfigEntryServiceResolverLoadBalancerHashPoliciesCookieConfigList
	CookieConfigInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Field() *string
	SetField(val *string)
	FieldInput() *string
	FieldValue() *string
	SetFieldValue(val *string)
	FieldValueInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SourceIp() interface{}
	SetSourceIp(val interface{})
	SourceIpInput() interface{}
	Terminal() interface{}
	SetTerminal(val interface{})
	TerminalInput() interface{}
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
	PutCookieConfig(value interface{})
	ResetCookieConfig()
	ResetField()
	ResetFieldValue()
	ResetSourceIp()
	ResetTerminal()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference
type jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) CookieConfig() ConfigEntryServiceResolverLoadBalancerHashPoliciesCookieConfigList {
	var returns ConfigEntryServiceResolverLoadBalancerHashPoliciesCookieConfigList
	_jsii_.Get(
		j,
		"cookieConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) CookieConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookieConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) FieldValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) FieldValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) SourceIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) SourceIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) Terminal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) TerminalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceResolver.ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference_Override(c ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceResolver.ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference)SetFieldValue(val *string) {
	if err := j.validateSetFieldValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference)SetSourceIp(val interface{}) {
	if err := j.validateSetSourceIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIp",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference)SetTerminal(val interface{}) {
	if err := j.validateSetTerminalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminal",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) PutCookieConfig(value interface{}) {
	if err := c.validatePutCookieConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCookieConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) ResetCookieConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetCookieConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) ResetField() {
	_jsii_.InvokeVoid(
		c,
		"resetField",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) ResetFieldValue() {
	_jsii_.InvokeVoid(
		c,
		"resetFieldValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) ResetTerminal() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminal",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceResolverLoadBalancerHashPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

