// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autopilotconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-consul.autopilotConfig.AutopilotConfig",
		reflect.TypeOf((*AutopilotConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cleanupDeadServers", GoGetter: "CleanupDeadServers"},
			_jsii_.MemberProperty{JsiiProperty: "cleanupDeadServersInput", GoGetter: "CleanupDeadServersInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "datacenter", GoGetter: "Datacenter"},
			_jsii_.MemberProperty{JsiiProperty: "datacenterInput", GoGetter: "DatacenterInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disableUpgradeMigration", GoGetter: "DisableUpgradeMigration"},
			_jsii_.MemberProperty{JsiiProperty: "disableUpgradeMigrationInput", GoGetter: "DisableUpgradeMigrationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastContactThreshold", GoGetter: "LastContactThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "lastContactThresholdInput", GoGetter: "LastContactThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxTrailingLogs", GoGetter: "MaxTrailingLogs"},
			_jsii_.MemberProperty{JsiiProperty: "maxTrailingLogsInput", GoGetter: "MaxTrailingLogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "redundancyZoneTag", GoGetter: "RedundancyZoneTag"},
			_jsii_.MemberProperty{JsiiProperty: "redundancyZoneTagInput", GoGetter: "RedundancyZoneTagInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCleanupDeadServers", GoMethod: "ResetCleanupDeadServers"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatacenter", GoMethod: "ResetDatacenter"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableUpgradeMigration", GoMethod: "ResetDisableUpgradeMigration"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLastContactThreshold", GoMethod: "ResetLastContactThreshold"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTrailingLogs", GoMethod: "ResetMaxTrailingLogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedundancyZoneTag", GoMethod: "ResetRedundancyZoneTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetServerStabilizationTime", GoMethod: "ResetServerStabilizationTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpgradeVersionTag", GoMethod: "ResetUpgradeVersionTag"},
			_jsii_.MemberProperty{JsiiProperty: "serverStabilizationTime", GoGetter: "ServerStabilizationTime"},
			_jsii_.MemberProperty{JsiiProperty: "serverStabilizationTimeInput", GoGetter: "ServerStabilizationTimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeVersionTag", GoGetter: "UpgradeVersionTag"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeVersionTagInput", GoGetter: "UpgradeVersionTagInput"},
		},
		func() interface{} {
			j := jsiiProxy_AutopilotConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-consul.autopilotConfig.AutopilotConfigConfig",
		reflect.TypeOf((*AutopilotConfigConfig)(nil)).Elem(),
	)
}
