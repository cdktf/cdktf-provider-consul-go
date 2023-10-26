// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceintentions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v8/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v8/configentryserviceintentions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference interface {
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
	Exact() *string
	SetExact(val *string)
	ExactInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Invert() interface{}
	SetInvert(val interface{})
	InvertInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	Present() interface{}
	SetPresent(val interface{})
	PresentInput() interface{}
	Regex() *string
	SetRegex(val *string)
	RegexInput() *string
	Suffix() *string
	SetSuffix(val *string)
	SuffixInput() *string
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
	ResetExact()
	ResetInvert()
	ResetPrefix()
	ResetPresent()
	ResetRegex()
	ResetSuffix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference
type jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) Exact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ExactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) Invert() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) InvertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) Present() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"present",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) PresentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"presentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) RegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) Suffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) SuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceIntentions.ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference_Override(c ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceIntentions.ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetExact(val *string) {
	if err := j.validateSetExactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exact",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetInvert(val interface{}) {
	if err := j.validateSetInvertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invert",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetPresent(val interface{}) {
	if err := j.validateSetPresentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"present",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetRegex(val *string) {
	if err := j.validateSetRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetSuffix(val *string) {
	if err := j.validateSetSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suffix",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ResetExact() {
	_jsii_.InvokeVoid(
		c,
		"resetExact",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ResetInvert() {
	_jsii_.InvokeVoid(
		c,
		"resetInvert",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ResetPresent() {
	_jsii_.InvokeVoid(
		c,
		"resetPresent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		c,
		"resetRegex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ResetSuffix() {
	_jsii_.InvokeVoid(
		c,
		"resetSuffix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceIntentionsSourcesPermissionsHttpHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

