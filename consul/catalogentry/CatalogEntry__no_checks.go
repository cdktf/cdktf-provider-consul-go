// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package catalogentry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogEntry) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntry) validatePutServiceParameters(value interface{}) error {
	return nil
}

func validateCatalogEntry_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCatalogEntry_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCatalogEntry_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCatalogEntry_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogEntry) validateSetAddressParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogEntry) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogEntry) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogEntry) validateSetDatacenterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogEntry) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogEntry) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_CatalogEntry) validateSetNodeAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogEntry) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogEntry) validateSetTokenParameters(val *string) error {
	return nil
}

func validateNewCatalogEntryParameters(scope constructs.Construct, id *string, config *CatalogEntryConfig) error {
	return nil
}

