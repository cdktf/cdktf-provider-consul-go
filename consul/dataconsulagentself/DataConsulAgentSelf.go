package dataconsulagentself

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-consul-go/consul/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-consul-go/consul/v5/dataconsulagentself/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/data-sources/agent_self consul_agent_self}.
type DataConsulAgentSelf interface {
	cdktf.TerraformDataSource
	AclDatacenter() *string
	AclDefaultPolicy() *string
	AclDisabledTtl() *string
	AclDownPolicy() *string
	AclEnforce08Semantics() cdktf.IResolvable
	AclTtl() *string
	Addresses() cdktf.StringMap
	AdvertiseAddr() *string
	AdvertiseAddrs() cdktf.StringMap
	AdvertiseAddrWan() *string
	AtlasJoin() cdktf.IResolvable
	BindAddr() *string
	BootstrapExpect() *float64
	BootstrapMode() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CheckDeregisterIntervalMin() *string
	CheckReapInterval() *string
	CheckUpdateInterval() *string
	ClientAddr() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Datacenter() *string
	DataDir() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DevMode() cdktf.IResolvable
	Dns() cdktf.StringMap
	DnsRecursors() *[]*string
	Domain() *string
	EnableAnonymousSignature() cdktf.IResolvable
	EnableCoordinates() cdktf.IResolvable
	EnableDebug() cdktf.IResolvable
	EnableRemoteExec() cdktf.IResolvable
	EnableSyslog() cdktf.IResolvable
	EnableUi() cdktf.IResolvable
	EnableUpdateCheck() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	LeaveOnInt() cdktf.IResolvable
	LeaveOnTerm() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogLevel() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Performance() cdktf.StringMap
	PidFile() *string
	Ports() cdktf.NumberMap
	ProtocolVersion() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReconnectTimeoutLan() *string
	ReconnectTimeoutWan() *string
	RejoinAfterLeave() cdktf.IResolvable
	RetryJoin() *[]*string
	RetryJoinEc2() cdktf.StringMap
	RetryJoinGce() cdktf.StringMap
	RetryJoinWan() *[]*string
	RetryMaxAttempts() *float64
	RetryMaxAttemptsWan() *float64
	SerfLanBindAddr() *string
	SerfWanBindAddr() *string
	ServerMode() cdktf.IResolvable
	ServerName() *string
	SessionTtlMin() *string
	StartJoin() *[]*string
	StartJoinWan() *[]*string
	SyslogFacility() *string
	TaggedAddresses() cdktf.StringMap
	Telemetry() cdktf.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsCaFile() *string
	TlsCertFile() *string
	TlsKeyFile() *string
	TlsMinVersion() *string
	TlsVerifyIncoming() cdktf.IResolvable
	TlsVerifyOutgoing() cdktf.IResolvable
	TlsVerifyServerHostname() cdktf.IResolvable
	TranslateWanAddrs() cdktf.IResolvable
	UiDir() *string
	UnixSockets() cdktf.StringMap
	Version() *string
	VersionPrerelease() *string
	VersionRevision() *string
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataConsulAgentSelf
type jsiiProxy_DataConsulAgentSelf struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataConsulAgentSelf) AclDatacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclDatacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) AclDefaultPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) AclDisabledTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclDisabledTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) AclDownPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclDownPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) AclEnforce08Semantics() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"aclEnforce08Semantics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) AclTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Addresses() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) AdvertiseAddr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseAddr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) AdvertiseAddrs() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"advertiseAddrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) AdvertiseAddrWan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseAddrWan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) AtlasJoin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"atlasJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) BindAddr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindAddr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) BootstrapExpect() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootstrapExpect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) BootstrapMode() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"bootstrapMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) CheckDeregisterIntervalMin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkDeregisterIntervalMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) CheckReapInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkReapInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) CheckUpdateInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkUpdateInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) ClientAddr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAddr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Datacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) DataDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) DevMode() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"devMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Dns() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) DnsRecursors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsRecursors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) EnableAnonymousSignature() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableAnonymousSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) EnableCoordinates() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableCoordinates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) EnableDebug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableDebug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) EnableRemoteExec() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableRemoteExec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) EnableSyslog() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableSyslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) EnableUi() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) EnableUpdateCheck() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableUpdateCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) LeaveOnInt() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"leaveOnInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) LeaveOnTerm() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"leaveOnTerm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Performance() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"performance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) PidFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Ports() cdktf.NumberMap {
	var returns cdktf.NumberMap
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) ProtocolVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) ReconnectTimeoutLan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reconnectTimeoutLan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) ReconnectTimeoutWan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reconnectTimeoutWan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) RejoinAfterLeave() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rejoinAfterLeave",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) RetryJoin() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) RetryJoinEc2() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"retryJoinEc2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) RetryJoinGce() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"retryJoinGce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) RetryJoinWan() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryJoinWan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) RetryMaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryMaxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) RetryMaxAttemptsWan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryMaxAttemptsWan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) SerfLanBindAddr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serfLanBindAddr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) SerfWanBindAddr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serfWanBindAddr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) ServerMode() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"serverMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) SessionTtlMin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionTtlMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) StartJoin() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) StartJoinWan() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startJoinWan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) SyslogFacility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syslogFacility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TaggedAddresses() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"taggedAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Telemetry() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"telemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TlsCaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCaFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TlsCertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TlsKeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsKeyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TlsVerifyIncoming() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"tlsVerifyIncoming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TlsVerifyOutgoing() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"tlsVerifyOutgoing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TlsVerifyServerHostname() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"tlsVerifyServerHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) TranslateWanAddrs() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"translateWanAddrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) UiDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uiDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) UnixSockets() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"unixSockets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) VersionPrerelease() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionPrerelease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataConsulAgentSelf) VersionRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionRevision",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/data-sources/agent_self consul_agent_self} Data Source.
func NewDataConsulAgentSelf(scope constructs.Construct, id *string, config *DataConsulAgentSelfConfig) DataConsulAgentSelf {
	_init_.Initialize()

	if err := validateNewDataConsulAgentSelfParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataConsulAgentSelf{}

	_jsii_.Create(
		"@cdktf/provider-consul.dataConsulAgentSelf.DataConsulAgentSelf",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/consul/2.17.0/docs/data-sources/agent_self consul_agent_self} Data Source.
func NewDataConsulAgentSelf_Override(d DataConsulAgentSelf, scope constructs.Construct, id *string, config *DataConsulAgentSelfConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-consul.dataConsulAgentSelf.DataConsulAgentSelf",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataConsulAgentSelf)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataConsulAgentSelf)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataConsulAgentSelf)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataConsulAgentSelf)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataConsulAgentSelf)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataConsulAgentSelf_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataConsulAgentSelf_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.dataConsulAgentSelf.DataConsulAgentSelf",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataConsulAgentSelf_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataConsulAgentSelf_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.dataConsulAgentSelf.DataConsulAgentSelf",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataConsulAgentSelf_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataConsulAgentSelf_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-consul.dataConsulAgentSelf.DataConsulAgentSelf",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataConsulAgentSelf_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-consul.dataConsulAgentSelf.DataConsulAgentSelf",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataConsulAgentSelf) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataConsulAgentSelf) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataConsulAgentSelf) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataConsulAgentSelf) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataConsulAgentSelf) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataConsulAgentSelf) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataConsulAgentSelf) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataConsulAgentSelf) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataConsulAgentSelf) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataConsulAgentSelf) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataConsulAgentSelf) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulAgentSelf) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataConsulAgentSelf) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataConsulAgentSelf) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulAgentSelf) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulAgentSelf) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataConsulAgentSelf) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

