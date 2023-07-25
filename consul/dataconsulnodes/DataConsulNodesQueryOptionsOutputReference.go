package dataconsulnodes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v6/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v6/dataconsulnodes/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataConsulNodesQueryOptionsOutputReference interface {
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

// The jsii proxy struct for DataConsulNodesQueryOptionsOutputReference
type jsiiProxy_DataConsulNodesQueryOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) AllowStale() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) AllowStaleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) Datacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) DatacenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) Near() *string {
	var returns *string
	_jsii_.Get(
		j,
		"near",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) NearInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) NodeMeta() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeMeta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) NodeMetaInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeMetaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) PartitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) RequireConsistent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireConsistent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) RequireConsistentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireConsistentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) WaitIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) WaitIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) WaitTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) WaitTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitTimeInput",
		&returns,
	)
	return returns
}


func NewDataConsulNodesQueryOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataConsulNodesQueryOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataConsulNodesQueryOptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataConsulNodesQueryOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.dataConsulNodes.DataConsulNodesQueryOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataConsulNodesQueryOptionsOutputReference_Override(d DataConsulNodesQueryOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.dataConsulNodes.DataConsulNodesQueryOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetAllowStale(val interface{}) {
	if err := j.validateSetAllowStaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowStale",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetDatacenter(val *string) {
	if err := j.validateSetDatacenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenter",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetNear(val *string) {
	if err := j.validateSetNearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"near",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetNodeMeta(val *map[string]*string) {
	if err := j.validateSetNodeMetaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeMeta",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetPartition(val *string) {
	if err := j.validateSetPartitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partition",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetRequireConsistent(val interface{}) {
	if err := j.validateSetRequireConsistentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireConsistent",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetToken(val *string) {
	if err := j.validateSetTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetWaitIndex(val *float64) {
	if err := j.validateSetWaitIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitIndex",
		val,
	)
}

func (j *jsiiProxy_DataConsulNodesQueryOptionsOutputReference)SetWaitTime(val *string) {
	if err := j.validateSetWaitTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitTime",
		val,
	)
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ResetAllowStale() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowStale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ResetDatacenter() {
	_jsii_.InvokeVoid(
		d,
		"resetDatacenter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ResetNear() {
	_jsii_.InvokeVoid(
		d,
		"resetNear",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ResetNodeMeta() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeMeta",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ResetPartition() {
	_jsii_.InvokeVoid(
		d,
		"resetPartition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ResetRequireConsistent() {
	_jsii_.InvokeVoid(
		d,
		"resetRequireConsistent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ResetToken() {
	_jsii_.InvokeVoid(
		d,
		"resetToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ResetWaitIndex() {
	_jsii_.InvokeVoid(
		d,
		"resetWaitIndex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ResetWaitTime() {
	_jsii_.InvokeVoid(
		d,
		"resetWaitTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataConsulNodesQueryOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

