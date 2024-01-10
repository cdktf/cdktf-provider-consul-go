// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryserviceresolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v9/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v9/configentryserviceresolver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceResolverFailoverOutputReference interface {
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
	Datacenters() *[]*string
	SetDatacenters(val *[]*string)
	DatacentersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	SamenessGroup() *string
	SetSamenessGroup(val *string)
	SamenessGroupInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	ServiceSubset() *string
	SetServiceSubset(val *string)
	ServiceSubsetInput() *string
	SubsetName() *string
	SetSubsetName(val *string)
	SubsetNameInput() *string
	Targets() ConfigEntryServiceResolverFailoverTargetsList
	TargetsInput() interface{}
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
	PutTargets(value interface{})
	ResetDatacenters()
	ResetNamespace()
	ResetSamenessGroup()
	ResetService()
	ResetServiceSubset()
	ResetTargets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceResolverFailoverOutputReference
type jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) Datacenters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"datacenters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) DatacentersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"datacentersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) SamenessGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samenessGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) SamenessGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samenessGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ServiceSubset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSubset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ServiceSubsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSubsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) SubsetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subsetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) SubsetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subsetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) Targets() ConfigEntryServiceResolverFailoverTargetsList {
	var returns ConfigEntryServiceResolverFailoverTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceResolverFailoverOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ConfigEntryServiceResolverFailoverOutputReference {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceResolverFailoverOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceResolver.ConfigEntryServiceResolverFailoverOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConfigEntryServiceResolverFailoverOutputReference_Override(c ConfigEntryServiceResolverFailoverOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceResolver.ConfigEntryServiceResolverFailoverOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetDatacenters(val *[]*string) {
	if err := j.validateSetDatacentersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenters",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetSamenessGroup(val *string) {
	if err := j.validateSetSamenessGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samenessGroup",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetServiceSubset(val *string) {
	if err := j.validateSetServiceSubsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceSubset",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetSubsetName(val *string) {
	if err := j.validateSetSubsetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subsetName",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) PutTargets(value interface{}) {
	if err := c.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTargets",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ResetDatacenters() {
	_jsii_.InvokeVoid(
		c,
		"resetDatacenters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ResetSamenessGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetSamenessGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		c,
		"resetService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ResetServiceSubset() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceSubset",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ResetTargets() {
	_jsii_.InvokeVoid(
		c,
		"resetTargets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceResolverFailoverOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

