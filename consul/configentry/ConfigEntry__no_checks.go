// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package configentry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigEntry) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ConfigEntry) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateConfigEntry_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateConfigEntry_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfigEntry_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateConfigEntry_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigEntry) validateSetConfigJsonParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigEntry) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigEntry) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigEntry) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigEntry) validateSetKindParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigEntry) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ConfigEntry) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigEntry) validateSetNamespaceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigEntry) validateSetPartitionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigEntry) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewConfigEntryParameters(scope constructs.Construct, id *string, config *ConfigEntryConfig) error {
	return nil
}

