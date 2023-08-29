// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataconsulagentself

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-consul.dataConsulAgentSelf.DataConsulAgentSelf",
		reflect.TypeOf((*DataConsulAgentSelf)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aclDatacenter", GoGetter: "AclDatacenter"},
			_jsii_.MemberProperty{JsiiProperty: "aclDefaultPolicy", GoGetter: "AclDefaultPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "aclDisabledTtl", GoGetter: "AclDisabledTtl"},
			_jsii_.MemberProperty{JsiiProperty: "aclDownPolicy", GoGetter: "AclDownPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "aclEnforce08Semantics", GoGetter: "AclEnforce08Semantics"},
			_jsii_.MemberProperty{JsiiProperty: "aclTtl", GoGetter: "AclTtl"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "addresses", GoGetter: "Addresses"},
			_jsii_.MemberProperty{JsiiProperty: "advertiseAddr", GoGetter: "AdvertiseAddr"},
			_jsii_.MemberProperty{JsiiProperty: "advertiseAddrs", GoGetter: "AdvertiseAddrs"},
			_jsii_.MemberProperty{JsiiProperty: "advertiseAddrWan", GoGetter: "AdvertiseAddrWan"},
			_jsii_.MemberProperty{JsiiProperty: "atlasJoin", GoGetter: "AtlasJoin"},
			_jsii_.MemberProperty{JsiiProperty: "bindAddr", GoGetter: "BindAddr"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapExpect", GoGetter: "BootstrapExpect"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapMode", GoGetter: "BootstrapMode"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "checkDeregisterIntervalMin", GoGetter: "CheckDeregisterIntervalMin"},
			_jsii_.MemberProperty{JsiiProperty: "checkReapInterval", GoGetter: "CheckReapInterval"},
			_jsii_.MemberProperty{JsiiProperty: "checkUpdateInterval", GoGetter: "CheckUpdateInterval"},
			_jsii_.MemberProperty{JsiiProperty: "clientAddr", GoGetter: "ClientAddr"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "datacenter", GoGetter: "Datacenter"},
			_jsii_.MemberProperty{JsiiProperty: "dataDir", GoGetter: "DataDir"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "devMode", GoGetter: "DevMode"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dnsRecursors", GoGetter: "DnsRecursors"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "enableAnonymousSignature", GoGetter: "EnableAnonymousSignature"},
			_jsii_.MemberProperty{JsiiProperty: "enableCoordinates", GoGetter: "EnableCoordinates"},
			_jsii_.MemberProperty{JsiiProperty: "enableDebug", GoGetter: "EnableDebug"},
			_jsii_.MemberProperty{JsiiProperty: "enableRemoteExec", GoGetter: "EnableRemoteExec"},
			_jsii_.MemberProperty{JsiiProperty: "enableSyslog", GoGetter: "EnableSyslog"},
			_jsii_.MemberProperty{JsiiProperty: "enableUi", GoGetter: "EnableUi"},
			_jsii_.MemberProperty{JsiiProperty: "enableUpdateCheck", GoGetter: "EnableUpdateCheck"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "leaveOnInt", GoGetter: "LeaveOnInt"},
			_jsii_.MemberProperty{JsiiProperty: "leaveOnTerm", GoGetter: "LeaveOnTerm"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logLevel", GoGetter: "LogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "performance", GoGetter: "Performance"},
			_jsii_.MemberProperty{JsiiProperty: "pidFile", GoGetter: "PidFile"},
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
			_jsii_.MemberProperty{JsiiProperty: "protocolVersion", GoGetter: "ProtocolVersion"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "reconnectTimeoutLan", GoGetter: "ReconnectTimeoutLan"},
			_jsii_.MemberProperty{JsiiProperty: "reconnectTimeoutWan", GoGetter: "ReconnectTimeoutWan"},
			_jsii_.MemberProperty{JsiiProperty: "rejoinAfterLeave", GoGetter: "RejoinAfterLeave"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "retryJoin", GoGetter: "RetryJoin"},
			_jsii_.MemberProperty{JsiiProperty: "retryJoinEc2", GoGetter: "RetryJoinEc2"},
			_jsii_.MemberProperty{JsiiProperty: "retryJoinGce", GoGetter: "RetryJoinGce"},
			_jsii_.MemberProperty{JsiiProperty: "retryJoinWan", GoGetter: "RetryJoinWan"},
			_jsii_.MemberProperty{JsiiProperty: "retryMaxAttempts", GoGetter: "RetryMaxAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "retryMaxAttemptsWan", GoGetter: "RetryMaxAttemptsWan"},
			_jsii_.MemberProperty{JsiiProperty: "serfLanBindAddr", GoGetter: "SerfLanBindAddr"},
			_jsii_.MemberProperty{JsiiProperty: "serfWanBindAddr", GoGetter: "SerfWanBindAddr"},
			_jsii_.MemberProperty{JsiiProperty: "serverMode", GoGetter: "ServerMode"},
			_jsii_.MemberProperty{JsiiProperty: "serverName", GoGetter: "ServerName"},
			_jsii_.MemberProperty{JsiiProperty: "sessionTtlMin", GoGetter: "SessionTtlMin"},
			_jsii_.MemberProperty{JsiiProperty: "startJoin", GoGetter: "StartJoin"},
			_jsii_.MemberProperty{JsiiProperty: "startJoinWan", GoGetter: "StartJoinWan"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "syslogFacility", GoGetter: "SyslogFacility"},
			_jsii_.MemberProperty{JsiiProperty: "taggedAddresses", GoGetter: "TaggedAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "telemetry", GoGetter: "Telemetry"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsCaFile", GoGetter: "TlsCaFile"},
			_jsii_.MemberProperty{JsiiProperty: "tlsCertFile", GoGetter: "TlsCertFile"},
			_jsii_.MemberProperty{JsiiProperty: "tlsKeyFile", GoGetter: "TlsKeyFile"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMinVersion", GoGetter: "TlsMinVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tlsVerifyIncoming", GoGetter: "TlsVerifyIncoming"},
			_jsii_.MemberProperty{JsiiProperty: "tlsVerifyOutgoing", GoGetter: "TlsVerifyOutgoing"},
			_jsii_.MemberProperty{JsiiProperty: "tlsVerifyServerHostname", GoGetter: "TlsVerifyServerHostname"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "translateWanAddrs", GoGetter: "TranslateWanAddrs"},
			_jsii_.MemberProperty{JsiiProperty: "uiDir", GoGetter: "UiDir"},
			_jsii_.MemberProperty{JsiiProperty: "unixSockets", GoGetter: "UnixSockets"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionPrerelease", GoGetter: "VersionPrerelease"},
			_jsii_.MemberProperty{JsiiProperty: "versionRevision", GoGetter: "VersionRevision"},
		},
		func() interface{} {
			j := jsiiProxy_DataConsulAgentSelf{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-consul.dataConsulAgentSelf.DataConsulAgentSelfConfig",
		reflect.TypeOf((*DataConsulAgentSelfConfig)(nil)).Elem(),
	)
}
