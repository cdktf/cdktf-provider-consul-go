package service

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v5/jsii"

	"github.com/cdktf/cdktf-provider-consul-go/consul/v5/service/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceCheckOutputReference interface {
	cdktf.ComplexObject
	CheckId() *string
	SetCheckId(val *string)
	CheckIdInput() *string
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
	DeregisterCriticalServiceAfter() *string
	SetDeregisterCriticalServiceAfter(val *string)
	DeregisterCriticalServiceAfterInput() *string
	// Experimental.
	Fqn() *string
	Header() ServiceCheckHeaderList
	HeaderInput() interface{}
	Http() *string
	SetHttp(val *string)
	HttpInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	Method() *string
	SetMethod(val *string)
	MethodInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Notes() *string
	SetNotes(val *string)
	NotesInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tcp() *string
	SetTcp(val *string)
	TcpInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	TlsSkipVerify() interface{}
	SetTlsSkipVerify(val interface{})
	TlsSkipVerifyInput() interface{}
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
	PutHeader(value interface{})
	ResetDeregisterCriticalServiceAfter()
	ResetHeader()
	ResetHttp()
	ResetMethod()
	ResetNotes()
	ResetStatus()
	ResetTcp()
	ResetTlsSkipVerify()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceCheckOutputReference
type jsiiProxy_ServiceCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceCheckOutputReference) CheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) CheckIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) DeregisterCriticalServiceAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deregisterCriticalServiceAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) DeregisterCriticalServiceAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deregisterCriticalServiceAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) Header() ServiceCheckHeaderList {
	var returns ServiceCheckHeaderList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) Http() *string {
	var returns *string
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) HttpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) Notes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) NotesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) Tcp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) TcpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) TlsSkipVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsSkipVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCheckOutputReference) TlsSkipVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsSkipVerifyInput",
		&returns,
	)
	return returns
}


func NewServiceCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceCheckOutputReference {
	_init_.Initialize()

	if err := validateNewServiceCheckOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-consul.service.ServiceCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceCheckOutputReference_Override(s ServiceCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.service.ServiceCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetCheckId(val *string) {
	if err := j.validateSetCheckIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkId",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetDeregisterCriticalServiceAfter(val *string) {
	if err := j.validateSetDeregisterCriticalServiceAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deregisterCriticalServiceAfter",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetHttp(val *string) {
	if err := j.validateSetHttpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetInterval(val *string) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetNotes(val *string) {
	if err := j.validateSetNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetTcp(val *string) {
	if err := j.validateSetTcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcp",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_ServiceCheckOutputReference)SetTlsSkipVerify(val interface{}) {
	if err := j.validateSetTlsSkipVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsSkipVerify",
		val,
	)
}

func (s *jsiiProxy_ServiceCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) PutHeader(value interface{}) {
	if err := s.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHeader",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceCheckOutputReference) ResetDeregisterCriticalServiceAfter() {
	_jsii_.InvokeVoid(
		s,
		"resetDeregisterCriticalServiceAfter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceCheckOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		s,
		"resetHeader",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceCheckOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		s,
		"resetHttp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceCheckOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceCheckOutputReference) ResetNotes() {
	_jsii_.InvokeVoid(
		s,
		"resetNotes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceCheckOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceCheckOutputReference) ResetTcp() {
	_jsii_.InvokeVoid(
		s,
		"resetTcp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceCheckOutputReference) ResetTlsSkipVerify() {
	_jsii_.InvokeVoid(
		s,
		"resetTlsSkipVerify",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

