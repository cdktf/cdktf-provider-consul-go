// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataconsulcatalogservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v10/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v10/dataconsulcatalogservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataConsulCatalogServiceQueryOptionsOutputReference interface {
	cdktf.ComplexObject
	AllowStale() interface{}
	SetAllowStale(val interface{})
	AllowStaleInput() interface{}
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
	Datacenter() *string
	SetDatacenter(val *string)
	DatacenterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	Near() *string
	SetNear(val *string)
	NearInput() *string
	NodeMeta() *map[string]*string
	SetNodeMeta(val *map[string]*string)
	NodeMetaInput() *map[string]*string
	Partition() *string
	SetPartition(val *string)
	PartitionInput() *string
	RequireConsistent() interface{}
	SetRequireConsistent(val interface{})
	RequireConsistentInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	WaitIndex() *float64
	SetWaitIndex(val *float64)
	WaitIndexInput() *float64
	WaitTime() *string
	SetWaitTime(val *string)
	WaitTimeInput() *string
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
	ResetAllowStale()
	ResetDatacenter()
	ResetNamespace()
	ResetNear()
	ResetNodeMeta()
	ResetPartition()
	ResetRequireConsistent()
	ResetToken()
	ResetWaitIndex()
	ResetWaitTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataConsulCatalogServiceQueryOptionsOutputReference
type jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) AllowStale() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) AllowStaleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) Datacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) DatacenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) Near() *string {
	var returns *string
	_jsii_.Get(
		j,
		"near",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) NearInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) NodeMeta() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeMeta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) NodeMetaInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeMetaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) PartitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) RequireConsistent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireConsistent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) RequireConsistentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireConsistentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) WaitIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) WaitIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) WaitTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) WaitTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitTimeInput",
		&returns,
	)
	return returns
}


func NewDataConsulCatalogServiceQueryOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataConsulCatalogServiceQueryOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataConsulCatalogServiceQueryOptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.dataConsulCatalogService.DataConsulCatalogServiceQueryOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataConsulCatalogServiceQueryOptionsOutputReference_Override(d DataConsulCatalogServiceQueryOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.dataConsulCatalogService.DataConsulCatalogServiceQueryOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetAllowStale(val interface{}) {
	if err := j.validateSetAllowStaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowStale",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetDatacenter(val *string) {
	if err := j.validateSetDatacenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenter",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetNear(val *string) {
	if err := j.validateSetNearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"near",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetNodeMeta(val *map[string]*string) {
	if err := j.validateSetNodeMetaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeMeta",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetPartition(val *string) {
	if err := j.validateSetPartitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partition",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetRequireConsistent(val interface{}) {
	if err := j.validateSetRequireConsistentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireConsistent",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetToken(val *string) {
	if err := j.validateSetTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetWaitIndex(val *float64) {
	if err := j.validateSetWaitIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitIndex",
		val,
	)
}

func (j *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference)SetWaitTime(val *string) {
	if err := j.validateSetWaitTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitTime",
		val,
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ResetAllowStale() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowStale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ResetDatacenter() {
	_jsii_.InvokeVoid(
		d,
		"resetDatacenter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ResetNear() {
	_jsii_.InvokeVoid(
		d,
		"resetNear",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ResetNodeMeta() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeMeta",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ResetPartition() {
	_jsii_.InvokeVoid(
		d,
		"resetPartition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ResetRequireConsistent() {
	_jsii_.InvokeVoid(
		d,
		"resetRequireConsistent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ResetToken() {
	_jsii_.InvokeVoid(
		d,
		"resetToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ResetWaitIndex() {
	_jsii_.InvokeVoid(
		d,
		"resetWaitIndex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ResetWaitTime() {
	_jsii_.InvokeVoid(
		d,
		"resetWaitTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulCatalogServiceQueryOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

