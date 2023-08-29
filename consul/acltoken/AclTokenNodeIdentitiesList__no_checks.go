// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package acltoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AclTokenNodeIdentitiesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AclTokenNodeIdentitiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AclTokenNodeIdentitiesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AclTokenNodeIdentitiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AclTokenNodeIdentitiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AclTokenNodeIdentitiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAclTokenNodeIdentitiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

