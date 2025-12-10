// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicerouter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v10/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v10/configentryservicerouter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceRouterRoutesMatchHttpOutputReference interface {
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
	Header() ConfigEntryServiceRouterRoutesMatchHttpHeaderList
	HeaderInput() interface{}
	InternalValue() *ConfigEntryServiceRouterRoutesMatchHttp
	SetInternalValue(val *ConfigEntryServiceRouterRoutesMatchHttp)
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
	QueryParam() ConfigEntryServiceRouterRoutesMatchHttpQueryParamList
	QueryParamInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutHeader(value interface{})
	PutQueryParam(value interface{})
	ResetHeader()
	ResetMethods()
	ResetPathExact()
	ResetPathPrefix()
	ResetPathRegex()
	ResetQueryParam()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceRouterRoutesMatchHttpOutputReference
type jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) Header() ConfigEntryServiceRouterRoutesMatchHttpHeaderList {
	var returns ConfigEntryServiceRouterRoutesMatchHttpHeaderList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) InternalValue() *ConfigEntryServiceRouterRoutesMatchHttp {
	var returns *ConfigEntryServiceRouterRoutesMatchHttp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) Methods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) MethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) PathExact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathExact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) PathExactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathExactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) PathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) PathPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) PathRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) PathRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) QueryParam() ConfigEntryServiceRouterRoutesMatchHttpQueryParamList {
	var returns ConfigEntryServiceRouterRoutesMatchHttpQueryParamList
	_jsii_.Get(
		j,
		"queryParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) QueryParamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceRouterRoutesMatchHttpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConfigEntryServiceRouterRoutesMatchHttpOutputReference {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceRouterRoutesMatchHttpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceRouter.ConfigEntryServiceRouterRoutesMatchHttpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigEntryServiceRouterRoutesMatchHttpOutputReference_Override(c ConfigEntryServiceRouterRoutesMatchHttpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceRouter.ConfigEntryServiceRouterRoutesMatchHttpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference)SetInternalValue(val *ConfigEntryServiceRouterRoutesMatchHttp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference)SetMethods(val *[]*string) {
	if err := j.validateSetMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"methods",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference)SetPathExact(val *string) {
	if err := j.validateSetPathExactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathExact",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference)SetPathPrefix(val *string) {
	if err := j.validateSetPathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathPrefix",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference)SetPathRegex(val *string) {
	if err := j.validateSetPathRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathRegex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) PutHeader(value interface{}) {
	if err := c.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHeader",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) PutQueryParam(value interface{}) {
	if err := c.validatePutQueryParamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQueryParam",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) ResetMethods() {
	_jsii_.InvokeVoid(
		c,
		"resetMethods",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) ResetPathExact() {
	_jsii_.InvokeVoid(
		c,
		"resetPathExact",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) ResetPathPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetPathPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) ResetPathRegex() {
	_jsii_.InvokeVoid(
		c,
		"resetPathRegex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) ResetQueryParam() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryParam",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesMatchHttpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

