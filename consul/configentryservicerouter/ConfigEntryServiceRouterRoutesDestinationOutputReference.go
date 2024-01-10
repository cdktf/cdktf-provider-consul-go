// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicerouter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v9/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v9/configentryservicerouter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceRouterRoutesDestinationOutputReference interface {
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
	IdleTimeout() *string
	SetIdleTimeout(val *string)
	IdleTimeoutInput() *string
	InternalValue() *ConfigEntryServiceRouterRoutesDestination
	SetInternalValue(val *ConfigEntryServiceRouterRoutesDestination)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	NumRetries() *float64
	SetNumRetries(val *float64)
	NumRetriesInput() *float64
	Partition() *string
	SetPartition(val *string)
	PartitionInput() *string
	PrefixRewrite() *string
	SetPrefixRewrite(val *string)
	PrefixRewriteInput() *string
	RequestHeaders() ConfigEntryServiceRouterRoutesDestinationRequestHeadersOutputReference
	RequestHeadersInput() *ConfigEntryServiceRouterRoutesDestinationRequestHeaders
	RequestTimeout() *string
	SetRequestTimeout(val *string)
	RequestTimeoutInput() *string
	ResponseHeaders() ConfigEntryServiceRouterRoutesDestinationResponseHeadersOutputReference
	ResponseHeadersInput() *ConfigEntryServiceRouterRoutesDestinationResponseHeaders
	RetryOn() *[]*string
	SetRetryOn(val *[]*string)
	RetryOnConnectFailure() interface{}
	SetRetryOnConnectFailure(val interface{})
	RetryOnConnectFailureInput() interface{}
	RetryOnInput() *[]*string
	RetryOnStatusCodes() *[]*float64
	SetRetryOnStatusCodes(val *[]*float64)
	RetryOnStatusCodesInput() *[]*float64
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	ServiceSubset() *string
	SetServiceSubset(val *string)
	ServiceSubsetInput() *string
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
	PutRequestHeaders(value *ConfigEntryServiceRouterRoutesDestinationRequestHeaders)
	PutResponseHeaders(value *ConfigEntryServiceRouterRoutesDestinationResponseHeaders)
	ResetIdleTimeout()
	ResetNamespace()
	ResetNumRetries()
	ResetPartition()
	ResetPrefixRewrite()
	ResetRequestHeaders()
	ResetRequestTimeout()
	ResetResponseHeaders()
	ResetRetryOn()
	ResetRetryOnConnectFailure()
	ResetRetryOnStatusCodes()
	ResetService()
	ResetServiceSubset()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceRouterRoutesDestinationOutputReference
type jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) IdleTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) IdleTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) InternalValue() *ConfigEntryServiceRouterRoutesDestination {
	var returns *ConfigEntryServiceRouterRoutesDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) NumRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) NumRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) PartitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) PrefixRewrite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) PrefixRewriteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) RequestHeaders() ConfigEntryServiceRouterRoutesDestinationRequestHeadersOutputReference {
	var returns ConfigEntryServiceRouterRoutesDestinationRequestHeadersOutputReference
	_jsii_.Get(
		j,
		"requestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) RequestHeadersInput() *ConfigEntryServiceRouterRoutesDestinationRequestHeaders {
	var returns *ConfigEntryServiceRouterRoutesDestinationRequestHeaders
	_jsii_.Get(
		j,
		"requestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) RequestTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) RequestTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResponseHeaders() ConfigEntryServiceRouterRoutesDestinationResponseHeadersOutputReference {
	var returns ConfigEntryServiceRouterRoutesDestinationResponseHeadersOutputReference
	_jsii_.Get(
		j,
		"responseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResponseHeadersInput() *ConfigEntryServiceRouterRoutesDestinationResponseHeaders {
	var returns *ConfigEntryServiceRouterRoutesDestinationResponseHeaders
	_jsii_.Get(
		j,
		"responseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) RetryOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) RetryOnConnectFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnConnectFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) RetryOnConnectFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOnConnectFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) RetryOnInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) RetryOnStatusCodes() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"retryOnStatusCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) RetryOnStatusCodesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"retryOnStatusCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ServiceSubset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSubset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ServiceSubsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSubsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceRouterRoutesDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConfigEntryServiceRouterRoutesDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceRouterRoutesDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceRouter.ConfigEntryServiceRouterRoutesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigEntryServiceRouterRoutesDestinationOutputReference_Override(c ConfigEntryServiceRouterRoutesDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceRouter.ConfigEntryServiceRouterRoutesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetIdleTimeout(val *string) {
	if err := j.validateSetIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetInternalValue(val *ConfigEntryServiceRouterRoutesDestination) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetNumRetries(val *float64) {
	if err := j.validateSetNumRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numRetries",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetPartition(val *string) {
	if err := j.validateSetPartitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partition",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetPrefixRewrite(val *string) {
	if err := j.validateSetPrefixRewriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixRewrite",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetRequestTimeout(val *string) {
	if err := j.validateSetRequestTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetRetryOn(val *[]*string) {
	if err := j.validateSetRetryOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryOn",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetRetryOnConnectFailure(val interface{}) {
	if err := j.validateSetRetryOnConnectFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryOnConnectFailure",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetRetryOnStatusCodes(val *[]*float64) {
	if err := j.validateSetRetryOnStatusCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryOnStatusCodes",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetServiceSubset(val *string) {
	if err := j.validateSetServiceSubsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceSubset",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) PutRequestHeaders(value *ConfigEntryServiceRouterRoutesDestinationRequestHeaders) {
	if err := c.validatePutRequestHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) PutResponseHeaders(value *ConfigEntryServiceRouterRoutesDestinationResponseHeaders) {
	if err := c.validatePutResponseHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResponseHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetNumRetries() {
	_jsii_.InvokeVoid(
		c,
		"resetNumRetries",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetPartition() {
	_jsii_.InvokeVoid(
		c,
		"resetPartition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetPrefixRewrite() {
	_jsii_.InvokeVoid(
		c,
		"resetPrefixRewrite",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetRequestHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetResponseHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetResponseHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetRetryOn() {
	_jsii_.InvokeVoid(
		c,
		"resetRetryOn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetRetryOnConnectFailure() {
	_jsii_.InvokeVoid(
		c,
		"resetRetryOnConnectFailure",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetRetryOnStatusCodes() {
	_jsii_.InvokeVoid(
		c,
		"resetRetryOnStatusCodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		c,
		"resetService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ResetServiceSubset() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceSubset",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceRouterRoutesDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

