// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceintentions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v8/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v8/configentryserviceintentions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference interface {
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
	Headers() ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersList
	HeadersInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Methods() *[]*string
	SetMethods(val *[]*string)
	MethodsInput() *[]*string
	PathExact() *string
	SetPathExact(val *string)
	PathExactInput() *string
	PathPrefix() *string
	SetPathPrefix(val *string)
	PathPrefixInput() *string
	PathRegex() *string
	SetPathRegex(val *string)
	PathRegexInput() *string
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
	PutHeaders(value interface{})
	ResetHeaders()
	ResetMethods()
	ResetPathExact()
	ResetPathPrefix()
	ResetPathRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference
type jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) Headers() ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersList {
	var returns ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) Methods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) MethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) PathExact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathExact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) PathExactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathExactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) PathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) PathPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) PathRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) PathRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceIntentions.ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference_Override(c ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceIntentions.ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference)SetMethods(val *[]*string) {
	if err := j.validateSetMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"methods",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference)SetPathExact(val *string) {
	if err := j.validateSetPathExactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathExact",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference)SetPathPrefix(val *string) {
	if err := j.validateSetPathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathPrefix",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference)SetPathRegex(val *string) {
	if err := j.validateSetPathRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathRegex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) PutHeaders(value interface{}) {
	if err := c.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) ResetMethods() {
	_jsii_.InvokeVoid(
		c,
		"resetMethods",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) ResetPathExact() {
	_jsii_.InvokeVoid(
		c,
		"resetPathExact",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) ResetPathPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetPathPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) ResetPathRegex() {
	_jsii_.InvokeVoid(
		c,
		"resetPathRegex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

