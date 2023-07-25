package preparedquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-consul-go/consul/v6/preparedquery/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/prepared_query consul_prepared_query}.
type PreparedQuery interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Connect() interface{}
	SetConnect(val interface{})
	ConnectInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Datacenter() *string
	SetDatacenter(val *string)
	DatacenterInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Dns() PreparedQueryDnsOutputReference
	DnsInput() *PreparedQueryDns
	Failover() PreparedQueryFailoverOutputReference
	FailoverInput() *PreparedQueryFailover
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreCheckIds() *[]*string
	SetIgnoreCheckIds(val *[]*string)
	IgnoreCheckIdsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Near() *string
	SetNear(val *string)
	NearInput() *string
	// The tree node.
	Node() constructs.Node
	NodeMeta() *map[string]*string
	SetNodeMeta(val *map[string]*string)
	NodeMetaInput() *map[string]*string
	OnlyPassing() interface{}
	SetOnlyPassing(val interface{})
	OnlyPassingInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	ServiceMeta() *map[string]*string
	SetServiceMeta(val *map[string]*string)
	ServiceMetaInput() *map[string]*string
	Session() *string
	SetSession(val *string)
	SessionInput() *string
	StoredToken() *string
	SetStoredToken(val *string)
	StoredTokenInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	Template() PreparedQueryTemplateOutputReference
	TemplateInput() *PreparedQueryTemplate
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDns(value *PreparedQueryDns)
	PutFailover(value *PreparedQueryFailover)
	PutTemplate(value *PreparedQueryTemplate)
	ResetConnect()
	ResetDatacenter()
	ResetDns()
	ResetFailover()
	ResetId()
	ResetIgnoreCheckIds()
	ResetNear()
	ResetNodeMeta()
	ResetOnlyPassing()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceMeta()
	ResetSession()
	ResetStoredToken()
	ResetTags()
	ResetTemplate()
	ResetToken()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PreparedQuery
type jsiiProxy_PreparedQuery struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PreparedQuery) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Connect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) ConnectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Datacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) DatacenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Dns() PreparedQueryDnsOutputReference {
	var returns PreparedQueryDnsOutputReference
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) DnsInput() *PreparedQueryDns {
	var returns *PreparedQueryDns
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Failover() PreparedQueryFailoverOutputReference {
	var returns PreparedQueryFailoverOutputReference
	_jsii_.Get(
		j,
		"failover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) FailoverInput() *PreparedQueryFailover {
	var returns *PreparedQueryFailover
	_jsii_.Get(
		j,
		"failoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) IgnoreCheckIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreCheckIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) IgnoreCheckIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreCheckIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Near() *string {
	var returns *string
	_jsii_.Get(
		j,
		"near",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) NearInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) NodeMeta() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeMeta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) NodeMetaInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeMetaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) OnlyPassing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyPassing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) OnlyPassingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyPassingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) ServiceMeta() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"serviceMeta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) ServiceMetaInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"serviceMetaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Session() *string {
	var returns *string
	_jsii_.Get(
		j,
		"session",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) SessionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) StoredToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storedToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) StoredTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storedTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Template() PreparedQueryTemplateOutputReference {
	var returns PreparedQueryTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) TemplateInput() *PreparedQueryTemplate {
	var returns *PreparedQueryTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreparedQuery) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/prepared_query consul_prepared_query} Resource.
func NewPreparedQuery(scope constructs.Construct, id *string, config *PreparedQueryConfig) PreparedQuery {
	_init_.Initialize()

	if err := validateNewPreparedQueryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PreparedQuery{}

	_jsii_.Create(
		"@cdktf/provider-consul.preparedQuery.PreparedQuery",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/consul/2.18.0/docs/resources/prepared_query consul_prepared_query} Resource.
func NewPreparedQuery_Override(p PreparedQuery, scope constructs.Construct, id *string, config *PreparedQueryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.preparedQuery.PreparedQuery",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PreparedQuery)SetConnect(val interface{}) {
	if err := j.validateSetConnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connect",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetDatacenter(val *string) {
	if err := j.validateSetDatacenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenter",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetIgnoreCheckIds(val *[]*string) {
	if err := j.validateSetIgnoreCheckIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreCheckIds",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetNear(val *string) {
	if err := j.validateSetNearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"near",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetNodeMeta(val *map[string]*string) {
	if err := j.validateSetNodeMetaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeMeta",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetOnlyPassing(val interface{}) {
	if err := j.validateSetOnlyPassingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyPassing",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetServiceMeta(val *map[string]*string) {
	if err := j.validateSetServiceMetaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceMeta",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetSession(val *string) {
	if err := j.validateSetSessionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"session",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetStoredToken(val *string) {
	if err := j.validateSetStoredTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedToken",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_PreparedQuery)SetToken(val *string) {
	if err := j.validateSetTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func PreparedQuery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePreparedQuery_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.preparedQuery.PreparedQuery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PreparedQuery_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePreparedQuery_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.preparedQuery.PreparedQuery",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PreparedQuery_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePreparedQuery_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.preparedQuery.PreparedQuery",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PreparedQuery_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-consul.preparedQuery.PreparedQuery",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PreparedQuery) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PreparedQuery) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PreparedQuery) PutDns(value *PreparedQueryDns) {
	if err := p.validatePutDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDns",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PreparedQuery) PutFailover(value *PreparedQueryFailover) {
	if err := p.validatePutFailoverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFailover",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PreparedQuery) PutTemplate(value *PreparedQueryTemplate) {
	if err := p.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTemplate",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PreparedQuery) ResetConnect() {
	_jsii_.InvokeVoid(
		p,
		"resetConnect",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetDatacenter() {
	_jsii_.InvokeVoid(
		p,
		"resetDatacenter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetDns() {
	_jsii_.InvokeVoid(
		p,
		"resetDns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetFailover() {
	_jsii_.InvokeVoid(
		p,
		"resetFailover",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetIgnoreCheckIds() {
	_jsii_.InvokeVoid(
		p,
		"resetIgnoreCheckIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetNear() {
	_jsii_.InvokeVoid(
		p,
		"resetNear",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetNodeMeta() {
	_jsii_.InvokeVoid(
		p,
		"resetNodeMeta",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetOnlyPassing() {
	_jsii_.InvokeVoid(
		p,
		"resetOnlyPassing",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetServiceMeta() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceMeta",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetSession() {
	_jsii_.InvokeVoid(
		p,
		"resetSession",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetStoredToken() {
	_jsii_.InvokeVoid(
		p,
		"resetStoredToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetTemplate() {
	_jsii_.InvokeVoid(
		p,
		"resetTemplate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) ResetToken() {
	_jsii_.InvokeVoid(
		p,
		"resetToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreparedQuery) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreparedQuery) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

