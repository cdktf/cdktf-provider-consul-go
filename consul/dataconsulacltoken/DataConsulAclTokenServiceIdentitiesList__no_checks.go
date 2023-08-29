// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataconsulacltoken

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataConsulAclTokenServiceIdentitiesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataConsulAclTokenServiceIdentitiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataConsulAclTokenServiceIdentitiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataConsulAclTokenServiceIdentitiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataConsulAclTokenServiceIdentitiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataConsulAclTokenServiceIdentitiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

