// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configentryservicedefaults

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v8/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v8/configentryservicedefaults/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList
type jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList {
	_init_.Initialize()

	if err := validateNewConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList{}

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList_Override(c ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.configEntryServiceDefaults.ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList) Get(index *float64) ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigEntryServiceDefaultsUpstreamConfigOverridesPassiveHealthCheckList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

